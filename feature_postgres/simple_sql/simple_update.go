package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Update(ctx context.Context, conn *pgx.Conn, book BookModel) error {
	sqlQuery := `
	UPDATE books
	SET
		title=$1, 
		author=$2, 
		rewiev=$3, 
		release_date=$4, 
		full_readed=$5, 
		receipt_time=$6, 
		fullreaded_time=$7
	WHERE ID=$8;
	`
	_, err := conn.Exec(
		ctx,
		sqlQuery,

		book.Title,
		book.Author,
		book.Rewiev,
		book.ReleaseDate,
		book.FullReaded,
		book.ReceiptTime,
		book.FullReadedTime,
		book.ID,
	)

	return err
}
