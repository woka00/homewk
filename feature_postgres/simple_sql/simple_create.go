package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS books(
		ID SERIAL PRIMARY KEY,
		title VARCHAR(200) NOT NULL,
		author VARCHAR(50) NOT NULL,
		rewiev VARCHAR(50),
		release_date TIMESTAMP NOT NULL,
		full_readed BOOLEAN,
		receipt_time TIMESTAMP NOT NULL,
		fullreaded_time TIMESTAMP
	);
	`
	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
