package simple_connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() {
	ctx := context.Background()

	conn, _ := pgx.Connect(ctx, "postgres://postgres:admin@localhost:5432/postgres")

	err := conn.Ping(ctx)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("succeed!")
	}
}

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "postgres://postgres:admin@localhost:5432/postgres")
}
