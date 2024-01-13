package queue

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type BankOrder struct {
	Sum           int64     `json:"sum"`
	OperationDate time.Time `json:"date"`
}

type OrdersChunk struct {
	Token      string      `json:"token"`
	Id         uuid.UUID   `json:"uuid"`
	BankOrders []BankOrder `json:"orders"`
}

type JobsPool struct {
	jobs []Job
}

type Job struct {
	Id            uuid.UUID `json:"uuid"`
	Sum           int64     `json:"sum"`
	OperationDate time.Time `json:"date"`
}

func NewJobsPool(ordersChunkJSON []byte) *JobsPool {
	var ordersChunk OrdersChunk
	_ = json.Unmarshal(ordersChunkJSON, &ordersChunk)
	jobs := make([]Job, len(ordersChunk.BankOrders))
	for ind, bankOrder := range ordersChunk.BankOrders {
		jobs[ind] = Job{
			Id:            ordersChunk.Id,
			Sum:           bankOrder.Sum,
			OperationDate: bankOrder.OperationDate,
		}
	}
	return &JobsPool{jobs: jobs}
}
