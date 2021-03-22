package infrastructure

import (
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/lib/pq"

	"github.com/juicypy/todo_list_service/src/config"
)

const DIALECT = "postgres"

func ConnectDB(cfg config.StorageConfig) (*sql.DB, *goqu.Database, error) {
	db, err := sql.Open(DIALECT, cfg.DSN())
	if err != nil {
		return nil, nil, err
	}
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(5 * time.Minute)
	return db, goqu.New(DIALECT, db), nil
}
