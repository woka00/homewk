package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

/*
1. Опишите на Golang функцию, которая принимает в аргументах
структуру-модель книги и при помощи SQL запроса добавляет эту
книгу в таблицу books. Вызовите эту функцию в вашем коде,
запустите программу, посмотрите через pgAdmin что действительно
в заданная книга создалась в базе данных. Поменяйте значения полей
передаваемой функции модели, после перезапустите вашу программу.
Через pgAdmin проверьте, что в таблицу books появилась вторая
запись, представляющая вторую сохранённую на книжной полке книгу.
Проделайте эти операции несколько раз на ваше усмотрение,
чтобы в таблицу books добавить 3–5 различных книг.
*/

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
