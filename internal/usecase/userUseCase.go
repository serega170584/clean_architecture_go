package usecase

import (
	"clean_architecture_go/internal/infrastructure/repository"
	"github.com/google/uuid"
	"time"
)

type Queue interface {
	Handle(orderJobJSON []byte)
}

type Repository interface {
	Get(login string, password string) (*repository.Token, error)
}

type TransferJobListener interface {
	Listen()
	Handle(transfersJSON []byte)
}

type UserUseCase struct {
	repo                Repository
	transferJobListener TransferJobListener
	orderJobQueue       Queue
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

func New(r Repository, transferJobListener TransferJobListener, orderJobQueue Queue) *UserUseCase {
	return &UserUseCase{repo: r, transferJobListener: transferJobListener, orderJobQueue: orderJobQueue}
}

func (uc *UserUseCase) Do(login string, password string) (*repository.Token, error) {
	return uc.repo.Get(login, password)
}

func (uc *UserUseCase) AddOrders(ordersChunkJSON []byte) ([]byte, error) {
	uc.orderJobQueue.Handle(ordersChunkJSON)
	return nil, nil
}

func (uc *UserUseCase) AddTransfers(transfersChunkJSON []byte) ([]byte, error) {
	uc.transferJobListener.Handle(transfersChunkJSON)
	return nil, nil
}

func (uc *UserUseCase) TransferJobListen() {
	uc.transferJobListener.Listen()
}
