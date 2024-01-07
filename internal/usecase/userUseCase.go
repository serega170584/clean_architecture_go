package usecase

import (
	"clean_architecture_go/internal/infrastructure/repository"
)

type Repository interface {
	Get(login string, password string) (*repository.Token, error)
}

type UserUseCase struct {
	repo Repository
}

func New(r Repository) *UserUseCase { return &UserUseCase{r} }

func (uc *UserUseCase) Do(login string, password string) (*repository.Token, error) {
	return uc.repo.Get(login, password)
}
