package models_test

import (
	"context"
	"fmt"
	"github.com/alextonkonogov/gb_go_postgres/pkg/config"
	db "github.com/alextonkonogov/gb_go_postgres/pkg/database"
	"github.com/alextonkonogov/gb_go_postgres/pkg/models"
	"testing"
)

func TestNewAuthor(t *testing.T) {
	ctx := context.Background()
	dbpool, err := db.InitDBConn(ctx, &config.AppConfig{1, 1, true})
	if err != nil {
		return
	}
	defer dbpool.Close()

	author := models.NewAuthor("Джон", "Толкиен")

	id, err := author.Add(ctx, dbpool)
	if err != nil {
		t.Fail()
	}
	fmt.Printf("Добавлен новый автор %s %s с идентификатором %d\n", author.Name, author.Surname, id)

	err = author.Delete(ctx, dbpool)
	if err != nil {
		t.Fail()
	}
	fmt.Printf("Удален автор %s %s с идентификатором %d\n", author.Name, author.Surname, id)

}
