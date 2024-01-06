package connection

import (
	"clean_architecture_go/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"net"
	"os"
)

func NewConnection(config *config.DBConfig) *pgx.Conn {
	url := fmt.Sprintf("postgres://%s:%s@%s/%s", config.User, config.Password, net.JoinHostPort(config.Host, config.Port), config.Database)
	c, err := pgx.Connect(context.Background(), url)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return c
}
