package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/alextonkonogov/gb_go_postgres/homework4/pkg/config"
	db "github.com/alextonkonogov/gb_go_postgres/homework4/pkg/database"
	"github.com/alextonkonogov/gb_go_postgres/homework4/pkg/models"
)

type AttackResults struct {
	Duration         time.Duration
	Threads          int
	QueriesPerformed uint64
}

func attack(ctx context.Context, duration time.Duration, threads int, dbpool *pgxpool.Pool) AttackResults {
	var queries uint64
	request := *config.Author

	attacker := func(stopAt time.Time) {
		for {

			authors, err := models.GetAuthorBySurname(ctx, dbpool, request)
			if err != nil {
				log.Fatal(err)
			}
			for _, a := range authors {
				_, err := models.GetBooksByAuthorId(ctx, dbpool, a.Id)
				if err != nil {
					log.Fatal(err)
				}
			}

			atomic.AddUint64(&queries, 1)

			if time.Now().After(stopAt) {
				return
			}
		}
	}

	var wg sync.WaitGroup
	wg.Add(threads)

	startAt := time.Now()
	stopAt := startAt.Add(duration)

	for i := 0; i < threads; i++ {
		go func() {
			attacker(stopAt)
			wg.Done()
		}()
	}

	wg.Wait()

	return AttackResults{
		Duration:         time.Now().Sub(startAt),
		Threads:          threads,
		QueriesPerformed: queries,
	}
}

func main() {
	cnfg, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	dbpool, err := db.InitDBConn(ctx, cnfg)
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	if cnfg.InitDB {
		err = db.InitTables(ctx, dbpool)
		if err != nil {
			log.Fatal(err)
		}
	}

	duration := time.Duration(10 * time.Second)
	threads := 1000
	fmt.Println("start attack")
	res := attack(ctx, duration, threads, dbpool)

	fmt.Println("duration:", res.Duration)
	fmt.Println("threads:", res.Threads)
	fmt.Println("queries:", res.QueriesPerformed)
	qps := res.QueriesPerformed / uint64(res.Duration.Seconds())
	fmt.Println("QPS:", qps)
}
