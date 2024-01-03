package repository

import "fmt"

type Repo struct{}

func (r Repo) Get() {
	fmt.Println("Repo called")
}

func New() *Repo {
	return &Repo{}
}
