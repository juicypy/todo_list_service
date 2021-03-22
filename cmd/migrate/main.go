package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/juicypy/todo_list_service/src/config"
)

func main() {
	forceVersion := flag.Int("version", 0, "force version")
	direction := flag.String("direction", "up", "up / down")
	flag.Parse()

	log.Println("Migrating...")

	cfg, err := config.StorageConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://./migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	if *forceVersion > 0 {
		err = m.Migrate(uint(*forceVersion))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if *direction == "up" {
		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	} else if *direction == "down" {
		err = m.Steps(-1)
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

	log.Println("Done")
}
