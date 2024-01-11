package pool

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Listener struct {
	ch     chan Job
	size   int
	dbConn *pgxpool.Pool
}

func NewListener(size int, dbConn *pgxpool.Pool) *Listener {
	ch := make(chan Job, size)
	return &Listener{ch: ch, size: size, dbConn: dbConn}
}

func (listener *Listener) Listen() {
	for i := 0; i < listener.size; i++ {
		go func(listener *Listener) {
			job := <-listener.ch
			conn := listener.dbConn
			_, _ = conn.Exec(context.Background(), `
INSERT INTO transfer(sum, operation_date, chunk_uuid) 
VALUES($1, $2, $3)`, job.Sum, job.OperationDate, job.Id)
		}(listener)
	}
}

func (listener *Listener) Handle(transferJobJSON []byte) {
	pool := NewJobsPool(transferJobJSON, listener.ch)
	pool.send()
}
