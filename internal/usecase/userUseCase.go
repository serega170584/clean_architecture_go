package usecase

import (
	"clean_architecture_go/internal/infrastructure/repository"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

type Repository interface {
	Get(login string, password string) (*repository.Token, error)
}

type UserUseCase struct {
	repo Repository
}

type Transfer struct {
	Sum           int64     `json:"sum"`
	OperationDate time.Time `json:"date"`
}

type TransfersChunk struct {
	Token     string     `json:"token"`
	Id        uuid.UUID  `json:"uuid"`
	Transfers []Transfer `json:"transfer"`
}

func New(r Repository) *UserUseCase { return &UserUseCase{r} }

func (uc *UserUseCase) Do(login string, password string) (*repository.Token, error) {
	return uc.repo.Get(login, password)
}

func (uc *UserUseCase) AddTransfers(transfersChunkJSON []byte) ([]byte, error) {
	var transferChunk TransfersChunk
	err := json.Unmarshal(transfersChunkJSON, &transferChunk)
	if err != nil {
		return nil, errors.Errorf("Transfers chunk decode error: %s", err.Error())
	}
	return nil, nil
}
