package repository

import (
	"database/sql"
	"go.uber.org/zap"
)

type Repository struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Repository {
	return &Repository{logger: logger}
}

func (repo *Repository) Get() {
	db, err := sql.Open("sdasd", "database=test")
	if err != nil {
		repo.logger.Error("Db connection error")
	}
	_, _ = db.Query("SELECT * FROM user")
}
