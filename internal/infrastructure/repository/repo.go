package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type Repo struct {
	conn *pgx.Conn
}

func (r Repo) Get(login string, password string) (*User, error) {
	user := &User{}
	err := r.conn.QueryRow(context.Background(), `
		SELECT login, first_name, second_name
		FROM product_user
		WHERE login = %1 AND password = %2
`, login, password).Scan(user.login, user.firstName, user.secondName)
	return user, err
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
