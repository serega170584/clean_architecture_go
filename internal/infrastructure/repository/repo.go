package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	conn *pgxpool.Pool
}

type Token struct {
	token string
}

func (t *Token) Token() string {
	return t.token
}

type TransferChunk struct {
	chunk uuid.UUID
}

func (tc *TransferChunk) Uuid() uuid.UUID {
	return tc.chunk
}

func (r Repo) Get(login string, password string) (*Token, error) {
	token := &Token{}
	var t string
	err := r.conn.QueryRow(context.Background(), `
		SELECT token
		FROM product_user
		WHERE login = $1 AND password = $2
`, login, password).Scan(&t)

	if err != nil {
		return token, err
	}

	token.token = t
	return token, err
}

func New(conn *pgxpool.Pool) *Repo {
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
