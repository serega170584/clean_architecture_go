package app

import (
	controller "clean_architecture_go/internal/controller/http"
	"clean_architecture_go/internal/infrastructure/repository"
	"clean_architecture_go/internal/usecase"
)

func Run() {
	repo := repository.New()

	uc := usecase.New(repo)

	c := controller.New(uc)
	c.Serve()
}
