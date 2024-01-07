package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type Repo struct {
	conn *pgx.Conn
}

type Token struct {
	token string
}

func (r Repo) Get(login string, password string) (*Token, error) {
	token := &Token{}
	err := r.conn.QueryRow(context.Background(), `
		SELECT token
		FROM product_user
		WHERE login = %1 AND password = %2
`, login, password).Scan(token.token)
	return token, err
}

func New(conn *pgx.Conn) *Repo {
	return &Repo{conn}
}

type User struct {
	login      string
	firstName  string
	secondName string
}

func (u *User) GetLogin() string {
	return u.login
}

func (u *User) GetFirstName() string {
	return u.firstName
}

func (u *User) GetSecondName() string {
	return u.secondName
}
