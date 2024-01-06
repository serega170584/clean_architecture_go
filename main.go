package main

import (
	"clean_architecture_go/internal/application"
	"clean_architecture_go/internal/config"
)

func main() {
	conf := config.New()

	app := application.New(conf)

	app.Run()
}
