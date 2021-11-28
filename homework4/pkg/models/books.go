package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type book struct {
	Id       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	AuthorId int    `json:"author_id" db:"author_id"`
}

func GetBooksByAuthorId(ctx context.Context, dbpool *pgxpool.Pool, authorId int) (books []book, err error) {
	rows, err := dbpool.Query(ctx, `select id, title from books where author_id = $1`, authorId)
	if err != nil {
		err = fmt.Errorf("failed to query data: %w", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var b book
		err = rows.Scan(&b.Id, &b.Title)
		if err != nil {
			err = fmt.Errorf("failed to scan row: %w", err)
			return
		}
		books = append(books, b)
	}

	// Проверка, что во время выборки данных не происходило ошибок
	if rows.Err() != nil {
		err = fmt.Errorf("failed to read response: %w", rows.Err())
		return
	}

	return
}
