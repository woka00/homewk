package simple_sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func ListPages(
	n int,
	ctx context.Context,
	conn *pgx.Conn,

) error {
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
	LIMIT $1
	OFFSET $2
	`

	offset := 0
	for {
		count := 0
		rows, err := conn.Query(ctx, sqlQuery, n, offset)
		if err != nil {
			return err
		}

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
				panic(err)
				return err
			}

			count++
			fmt.Println(book)
		}

		rows.Close()
		fmt.Println("---------------------------------Конец страницы--------------")
		offset += n

		if count < 5 {
			break
		}

	}
	return nil
}
