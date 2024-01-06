package repository

import (
	"clean_architecture_go/internal/config"
	"fmt"
)

type Repo struct {
	config *config.DBConfig
}

func (r Repo) Get() {
	fmt.Println("Repo called")
}

func New(config *config.DBConfig) *Repo {
	return &Repo{config}
}
