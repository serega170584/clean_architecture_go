package app

import (
	"clean_architecture/app/internal/controller/http"
	repository2 "clean_architecture/app/internal/infrastructure/repository"
	"clean_architecture/app/internal/usecase"
	"fmt"
	"go.uber.org/zap"
)

func Run() {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			fmt.Println("logger err")
		}
	}(logger)

	repo := repository2.New(logger)
	uc := usecase.New(repo)

	c := http.New(uc)
	c.Serve()
}
