package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alextonkonogov/gb_go_postgres/homework4/pkg/config"
	"github.com/alextonkonogov/gb_go_postgres/homework4/pkg/database"
	"github.com/alextonkonogov/gb_go_postgres/homework4/pkg/models"
)

func main() {
	err := start()
	if err != nil {
		log.Fatal(err)
	}
}

func start() (err error) {
	cnfg, err := config.NewAppConfig()
	if err != nil {
		return
	}

	ctx := context.Background()
	dbpool, err := db.InitDBConn(ctx, cnfg)
	if err != nil {
		return
	}
	defer dbpool.Close()

	if cnfg.InitDB {
		err = db.InitTables(ctx, dbpool)
		if err != nil {
			return
		}
	}

	request := *config.Author

	authors, err := models.GetAuthorBySurname(ctx, dbpool, request)
	if err != nil {
		return
	}

	if len(authors) == 0 {
		fmt.Printf("По вашему запросу ничего не найдено\n")
		return
	}

	for _, a := range authors {
		books, err := models.GetBooksByAuthorId(ctx, dbpool, a.Id)
		if err != nil {
			return err
		}

		fmt.Printf("По запросу \"%s\" найдены книги:\n", request)
		for _, v := range books {
			fmt.Printf("- %s\n", v.Title)
		}
	}
	return
}
