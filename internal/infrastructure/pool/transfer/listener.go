package pool

import (
	"github.com/jackc/pgx/v5"
)

type Listener struct {
	ch     chan Job
	size   int
	dbConn *pgx.Conn
}

func NewListener(size int, dbConn *pgx.Conn) *Listener {
	ch := make(chan Job, size)
	return &Listener{ch: ch, size: size, dbConn: dbConn}
}

func (listener *Listener) Listen() {
	for i := 0; i < listener.size; i++ {
		go func(ch chan Job) {
			<-ch
		}(listener.ch)
	}
}

func (listener *Listener) Handle(transferJobJSON []byte) {
	pool := NewJobsPool(transferJobJSON, listener.ch)
	pool.send()
}
