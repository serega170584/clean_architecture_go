package controller

import (
	"clean_architecture_go/internal/config"
	"fmt"
	"log"
	"net"
	"net/http"
)

type UseCase interface {
	Do()
}

type Controller struct {
	uc   UseCase
	conf *config.AppConfig
}

func New(uc UseCase, conf *config.AppConfig) *Controller { return &Controller{uc, conf} }

func (c *Controller) Serve() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello, %q", r.URL.Path)
		c.uc.Do()
	})
	log.Fatal(http.ListenAndServe(net.JoinHostPort(c.conf.Host, c.conf.Port), nil))
}
