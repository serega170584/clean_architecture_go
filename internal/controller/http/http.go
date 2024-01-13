package controller

import (
	"clean_architecture_go/internal/config"
	"clean_architecture_go/internal/infrastructure/repository"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net"
	"net/http"
	"time"
)

type UserUseCase interface {
	Do(login string, password string) (*repository.Token, error)
	AddOrders(ordersChunkJSON []byte) ([]byte, error)
	AddTransfers(transfersChunkJSON []byte) ([]byte, error)
	TransferJobListen()
}

type Controller struct {
	uc   UserUseCase
	conf *config.AppConfig
}

type BankOrder struct {
	Sum           int64     `json:"sum"`
	OperationDate time.Time `json:"date"`
}

type OrdersChunk struct {
	Token  string      `json:"token"`
	Id     uuid.UUID   `json:"uuid"`
	Orders []BankOrder `json:"orders"`
}

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type TransfersChunk struct {
	Token     string     `json:"token"`
	Id        uuid.UUID  `json:"uuid"`
	Transfers []Transfer `json:"transfers"`
}

type Transfer struct {
	Sum           int64     `json:"sum"`
	OperationDate time.Time `json:"date"`
}

func New(uc UserUseCase, conf *config.AppConfig) *Controller { return &Controller{uc, conf} }

func (c *Controller) Serve() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, fmt.Sprintf("Method %s not allowed", r.Method), http.StatusMethodNotAllowed)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var user User
		err := decoder.Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := c.uc.Do(user.Login, user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responseToken := &Token{Token: token.Token()}
		_ = json.NewEncoder(w).Encode(responseToken)
	})

	http.HandleFunc("/transfers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, fmt.Sprintf("Method %s not allowed", r.Method), http.StatusMethodNotAllowed)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var transfersChunk TransfersChunk
		err := decoder.Decode(&transfersChunk)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		transfersChunkJSON, _ := json.Marshal(transfersChunk)
		_, _ = c.uc.AddTransfers(transfersChunkJSON)
	})

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, fmt.Sprintf("Method %s not allowed", r.Method), http.StatusMethodNotAllowed)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var ordersChunk OrdersChunk
		err := decoder.Decode(&ordersChunk)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		ordersChunkJSON, _ := json.Marshal(ordersChunk)
		_, _ = c.uc.AddOrders(ordersChunkJSON)
	})

	c.uc.TransferJobListen()

	log.Fatal(http.ListenAndServe(net.JoinHostPort(c.conf.Host, c.conf.Port), nil))
}
