package controller

import (
	"clean_architecture_go/internal/config"
	"clean_architecture_go/internal/infrastructure/repository"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
)

type UserUseCase interface {
	Do(login string, password string) (*repository.Token, error)
}

type Controller struct {
	uc   UserUseCase
	conf *config.AppConfig
}

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
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
		_ = json.NewEncoder(w).Encode("{}")
		c.uc.Do(user.Login, user.Password)
	})
	log.Fatal(http.ListenAndServe(net.JoinHostPort(c.conf.Host, c.conf.Port), nil))
}
