package simple_sql

import "time"

type BookModel struct {
	ID             int
	Title          string
	Author         string
	Rewiev         string
	ReleaseDate    *time.Time
	FullReaded     bool
	ReceiptTime    time.Time
	FullReadedTime *time.Time
}
