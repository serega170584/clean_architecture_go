package application

import (
	"clean_architecture_go/internal/config"
	controller "clean_architecture_go/internal/controller/http"
	"clean_architecture_go/internal/infrastructure/connection"
	"clean_architecture_go/internal/infrastructure/repository"
	"clean_architecture_go/internal/usecase"
	"context"
	"github.com/jackc/pgx/v5"
)

type Application struct {
	config *config.Config
}

func New(config *config.Config) *Application {
	return &Application{config}
}

func (app *Application) Run() {
	conn := connection.NewConnection(app.config.DB)
	defer func(conn *pgx.Conn, ctx context.Context) {
		_ = conn.Close(ctx)
	}(conn, context.Background())

	repo := repository.New(conn)

	uc := usecase.New(repo)

	c := controller.New(uc, app.config.App)
	c.Serve()
}
