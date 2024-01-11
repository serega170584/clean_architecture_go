package pool

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Transfer struct {
	Sum           int64     `json:"sum"`
	OperationDate time.Time `json:"date"`
}

type TransfersChunk struct {
	Token     string     `json:"token"`
	Id        uuid.UUID  `json:"uuid"`
	Transfers []Transfer `json:"transfers"`
}

type JobsPool struct {
	jobs []Job
	ch   chan Job
}

type Job struct {
	Id            uuid.UUID `json:"uuid"`
	Sum           int64     `json:"sum"`
	OperationDate time.Time `json:"date"`
}

func NewJobsPool(transfersChunkJSON []byte, ch chan Job) *JobsPool {
	var transfersChunk TransfersChunk
	_ = json.Unmarshal(transfersChunkJSON, &transfersChunk)
	jobs := make([]Job, len(transfersChunk.Transfers))
	for ind, transfer := range transfersChunk.Transfers {
		jobs[ind] = Job{
			Id:            transfersChunk.Id,
			Sum:           transfer.Sum,
			OperationDate: transfer.OperationDate,
		}
	}
	return &JobsPool{jobs: jobs, ch: ch}
}

func (jobsPool *JobsPool) send() {
	for _, job := range jobsPool.jobs {
		go func(job Job, ch chan Job) {
			ch <- job
		}(job, jobsPool.ch)
	}
}
