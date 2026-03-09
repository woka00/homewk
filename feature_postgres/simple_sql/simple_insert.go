package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Insert(
	ctx context.Context,
	conn *pgx.Conn,
	book BookModel,
) error {
	sqlQuery := `
	INSERT INTO books(
		title, 
		author, 
		rewiev, 
		release_date, 
		full_readed,
		receipt_time,
		fullreaded_time
		)
	VALUES ($1,$2,$3,$4,$5,$6,$7)
	`
	_, err := conn.Exec(ctx,
		sqlQuery,
		book.Title,
		book.Author,
		book.Rewiev,
		book.ReleaseDate,
		book.FullReaded,
		book.ReceiptTime,
		book.FullReadedTime,
	)

	return err
}
