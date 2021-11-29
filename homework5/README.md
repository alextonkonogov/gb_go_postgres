# gb_go_postgres Homework 5

## Как запустить интеграционный тест

```bat
go test -run TestMain ./pkg/storage/ -v -count 1
```

Пример:

![alt text](integration.png)

## Как запустить миграцию

Запускаем контейнер с пустой БД

```bat
docker run \
    --rm -it \
    -p 5432:5432 \
    --name postgres \
    -e POSTGRES_PASSWORD=password \
    -e PGDATA=/var/lib/postgresql/data \
    postgres:13.1
```

Запускаем миграцию БД

```console
migrate -database "postgresql://postgres:password@localhost:5432/postgres?sslmode=disable" -path migrations up
```

Проверяем работу приложения

```console
go run cmd/main.go --author=Кинг
```

Отказываем миграцию при необходимости

```console
migrate -database "postgresql://postgres:password@localhost:5432/postgres?sslmode=disable" -path migrations down
```

Пример:

![alt text](migration.png)