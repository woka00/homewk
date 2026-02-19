package main

import (
	"context"
	"postgres/feature_postgres/simple_connection"
	"postgres/feature_postgres/simple_sql"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	// now := time.Now()
	// t, err := time.Parse("2006.01.02", "2026.02.17")
	// if err != nil {
	// 	panic(err)
	// }
	// book := simple_sql.BookModel{
	// 	Title:          "tujh11",
	// 	Author:         "Чудинов Е.К.",
	// 	Rewiev:         "",
	// 	ReleaseDate:    &t,
	// 	FullReaded:     true,
	// 	ReceiptTime:    now,
	// 	FullReadedTime: nil,
	// }
	// if err := simple_sql.Insert(ctx, conn, book); err != nil {
	// 	panic(err)
	// }

	// if err := simple_sql.Update(ctx, conn, book); err != nil {
	// 	fmt.Println("Ошибка в функции Update")
	// 	panic(err)
	// }

	// sliceIDs := []int{1, 3}
	// if err := simple_sql.Delete(ctx, conn, sliceIDs); err != nil {
	// 	panic(err)
	// }

	// books, err := simple_sql.Select(ctx, conn)
	// if err != nil {
	// 	panic(err)
	// }

	// pp.Println(books)

	simple_sql.ListPages(5, ctx, conn)

}
