package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Select(ctx context.Context, conn *pgx.Conn) ([]BookModel, error) {
	sqlQuery := `
	SELECT 
		ID, 
		title, 
		author, 
		rewiev, 
		release_date, 
		full_readed, 
		receipt_time, 
		fullreaded_time
	FROM books
	ORDER BY ID;
	`
	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	books := make([]BookModel, 0)

	for rows.Next() {
		var book BookModel
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Rewiev,
			&book.ReleaseDate,
			&book.FullReaded,
			&book.ReceiptTime,
			&book.FullReadedTime,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, err
}
