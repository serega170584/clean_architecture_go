package application

import (
	"clean_architecture_go/internal/config"
	controller "clean_architecture_go/internal/controller/http"
	"clean_architecture_go/internal/infrastructure/connection"
	"clean_architecture_go/internal/infrastructure/deferred"
	pool "clean_architecture_go/internal/infrastructure/pool/transfer"
	"clean_architecture_go/internal/infrastructure/queue"
	"clean_architecture_go/internal/infrastructure/repository"
	"clean_architecture_go/internal/usecase"
	"github.com/jackc/pgx/v5/pgxpool"
)

const TransferListenerSize = 5
const OrderQueueSize = 3

type Application struct {
	config *config.Config
}

func New(config *config.Config) *Application {
	return &Application{config}
}

func (app *Application) Run() {
	conn := connection.NewConnection(app.config.DB)
	defer func(conn *pgxpool.Pool) {
		conn.Close()
	}(conn)

	repo := repository.New(conn)

	transferJobListener := pool.NewListener(TransferListenerSize, conn)

	q := queue.New(OrderQueueSize, app.config.Queue.IsAsync, conn)

	d := deferred.New(3, 5, 3, conn, 3)
	d.Handle()

	uc := usecase.New(repo, transferJobListener, q)

	c := controller.New(uc, app.config.App)
	c.Serve()
}
