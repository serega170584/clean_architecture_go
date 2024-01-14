package deferred

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Deferred struct {
	ch         chan Job
	limit      int
	cnt        int
	dbConn     *pgxpool.Pool
	tickPeriod int
}

func New(size int, limit int, cnt int, dbConn *pgxpool.Pool, tickPeriod int) *Deferred {
	ch := make(chan Job, size)
	return &Deferred{ch: ch, limit: limit, cnt: cnt, dbConn: dbConn, tickPeriod: tickPeriod}
}

func (d *Deferred) Handle() {
	ticker := time.NewTicker(time.Duration(d.tickPeriod) * time.Second)
	go func(ticker *time.Ticker) {
		for {
			select {
			case <-ticker.C:
				d.generateJobs()
			}
		}
	}(ticker)

	go func(ch chan Job) {
		for {
			job := <-ch
			fmt.Println(job.id)
		}
	}(d.ch)
}

func (d *Deferred) generateJobs() {
	for i := 0; i < d.cnt; i++ {
		offset := i * d.limit
		go func(d *Deferred) {
			var id int
			_ = d.dbConn.QueryRow(context.Background(), `
		SELECT id
		FROM bank_order
		WHERE status = TRUE
		ORDER BY id
		OFFSET $1
		LIMIT $2
`, offset, d.limit).Scan(&id)
			if id != 0 {
				_, _ = d.dbConn.Exec(context.Background(), `
UPDATE bank_order
SET status = FALSE
WHERE id = $1`, id)
				d.ch <- Job{id: id}
			}
		}(d)
	}
}
