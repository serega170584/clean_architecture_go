package queue

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Queue struct {
	ch      chan Job
	isAsynq bool
	dbConn  *pgxpool.Pool
}

func New(size int, isAsynq bool) *Queue {
	ch := make(chan Job, size)
	return &Queue{ch: ch, isAsynq: isAsynq}
}

func (q *Queue) processJobs(jobs []Job) {
	for _, job := range jobs {
		q.process(job)
	}
}

func (q *Queue) Handle(orderJobJSON []byte) {
	pool := NewJobsPool(orderJobJSON)
	processJobs := q.processJobs
	if q.isAsynq {
		go processJobs(pool.jobs)
	} else {
		processJobs(pool.jobs)
	}
}

func (q *Queue) process(job Job) {
	q.ch <- job

	go func(q *Queue) {
		job := <-q.ch
		conn := q.dbConn
		_, _ = conn.Exec(context.Background(), `
INSERT INTO bank_order(sum, operation_date, chunk_uuid) 
VALUES($1, $2, $3)`, job.Sum, job.OperationDate, job.Id)
	}(q)
}
