package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type author struct {
	Id      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Surname string `json:"surname" db:"surname"`
}

func NewAuthor(name, surname string) *author {
	return &author{Name: name, Surname: surname}
}

func (a *author) Add(ctx context.Context, dbpool *pgxpool.Pool) (id int, err error) {
	err = dbpool.QueryRow(ctx, `insert into authors (name, surname) values ($1, $2) returning id`,
		a.Name,
		a.Surname,
	).Scan(&id)
	if err != nil {
		err = fmt.Errorf("failed to insert author: %w", err)
	}
	a.Id = id
	return
}

func (a *author) Delete(ctx context.Context, dbpool *pgxpool.Pool) (err error) {
	_, err = dbpool.Exec(ctx, `delete from authors where id = $1`, a.Id)
	if err != nil {
		err = fmt.Errorf("failed to delete author: %w", err)
	}
	return
}

func GetAuthorBySurname(ctx context.Context, dbpool *pgxpool.Pool, surname string) (authors []author, err error) {
	rows, err := dbpool.Query(ctx, `select id, name, surname from authors where surname = $1`, surname)
	if err != nil {
		err = fmt.Errorf("failed to query data: %w", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var a author
		err = rows.Scan(&a.Id, &a.Name, &a.Surname)
		if err != nil {
			err = fmt.Errorf("failed to scan row: %w", err)
			return
		}
		authors = append(authors, a)
	}

	if rows.Err() != nil {
		err = fmt.Errorf("failed to read response: %w", rows.Err())
		return
	}

	return
}
