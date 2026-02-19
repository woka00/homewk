package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Delete(ctx context.Context, conn *pgx.Conn, slice []int) error {
	sqlQuery := `
	DELETE FROM books
	WHERE ID = ANY($1)
	`
	_, err := conn.Exec(ctx, sqlQuery, slice)

	return err
}
