package db

import (
	"database/sql"
	"reqship-api/helpers/env"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func Init() (db *bun.DB) {
	dsn := env.GetEnvVarByName("POSTGRES_CONNECTION_STRING")
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db = bun.NewDB(sqldb, pgdialect.New())
	return
}
