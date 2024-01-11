package connection

import (
	"clean_architecture_go/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"net"
	"os"
)

func NewConnection(config *config.DBConfig) *pgxpool.Pool {
	url := fmt.Sprintf("postgres://%s:%s@%s/%s", config.User, config.Password, net.JoinHostPort(config.Host, config.Port), config.Database)
	poolConfig, _ := pgxpool.ParseConfig(url)
	c, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return c
}
