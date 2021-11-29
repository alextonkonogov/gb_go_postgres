package storage_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/alextonkonogov/gb_go_postgres/homework5/pkg/config"
	"github.com/alextonkonogov/gb_go_postgres/homework5/pkg/storage"
)

func TestNewAuthor(t *testing.T) {
	ctx := context.Background()
	dbpool, err := storage.InitDBConn(ctx, &config.AppConfig{1, 1, true})
	if err != nil {
		return
	}
	defer dbpool.Close()

	author := storage.NewAuthor("Джон", "Толкиен")

	id, err := author.Add(ctx, dbpool)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	fmt.Printf("Добавлен новый автор %s %s с идентификатором %d\n", author.Name, author.Surname, id)

	_, err = storage.GetAuthorById(ctx, dbpool, id)
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	fmt.Printf("Найден автор %s %s с идентификатором %d\n", author.Name, author.Surname, id)

	err = author.Delete(ctx, dbpool)
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	fmt.Printf("Удален автор %s %s с идентификатором %d\n", author.Name, author.Surname, id)
}

func TestGetAuthorById(t *testing.T) {
	ctx := context.Background()
	dbpool, err := storage.InitDBConn(ctx, &config.AppConfig{1, 1, true})
	if err != nil {
		return
	}
	defer dbpool.Close()

	ids := []int{
		1, 2, 3, 4,
	}

	for _, id := range ids {
		author, err := storage.GetAuthorById(ctx, dbpool, id)
		if err != nil {
			log.Println(err)
			t.Fail()
		}

		fmt.Printf("Идентификатор: %d - найден автор %s %s\n", id, author.Name, author.Surname)
	}
}
