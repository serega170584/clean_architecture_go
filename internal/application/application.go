package application

import (
	"clean_architecture_go/internal/config"
	controller "clean_architecture_go/internal/controller/http"
	"clean_architecture_go/internal/infrastructure/repository"
	"clean_architecture_go/internal/usecase"
)

type Application struct {
	config *config.Config
}

func New(config *config.Config) *Application {
	return &Application{config}
}

func (app *Application) Run() {
	repo := repository.New(app.config.DB)

	uc := usecase.New(repo)

	c := controller.New(uc, app.config.App)
	c.Serve()
}
