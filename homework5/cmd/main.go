package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alextonkonogov/gb_go_postgres/homework5/pkg/config"
	db "github.com/alextonkonogov/gb_go_postgres/homework5/pkg/storage"
)

//1. Реализовать приложение, которое реализует основной use-case вашей системы, т.е. поддерживает выполнение типовых запросов  (из файла queries.sql урока 3,
//   достаточно покрыть один-два запроса).
//   Необходимо реализовать только Storage Layer вашего приложения, т.е. только часть взаимодействия с базой данных.
//2. Реализовать интеграционное тестирование функциональности по выборке данных из  базы.
//3. Реализовать автоматизацию миграции структуры базы данных (файл schema.sql из предыдущих уроков).
//   В файле README.md в корне проекта описать, как запускать миграцию структуры базы данных.

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

	authors, err := db.GetAuthorBySurname(ctx, dbpool, request)
	if err != nil {
		return
	}

	if len(authors) == 0 {
		fmt.Printf("По вашему запросу ничего не найдено\n")
		return
	}

	for _, a := range authors {
		books, err := db.GetBooksByAuthorId(ctx, dbpool, a.Id)
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
