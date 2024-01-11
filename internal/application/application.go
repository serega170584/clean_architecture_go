package application

import (
	"clean_architecture_go/internal/config"
	controller "clean_architecture_go/internal/controller/http"
	"clean_architecture_go/internal/infrastructure/connection"
	pool "clean_architecture_go/internal/infrastructure/pool/transfer"
	"clean_architecture_go/internal/infrastructure/repository"
	"clean_architecture_go/internal/usecase"
	"github.com/jackc/pgx/v5/pgxpool"
)

const TransferListenerSize = 5

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

	uc := usecase.New(repo, transferJobListener)

	c := controller.New(uc, app.config.App)
	c.Serve()
}
