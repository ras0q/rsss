package database

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func NewSQLiteDB(dataSourceName string) (DB, error) {
	sqldb, err := sql.Open(sqliteshim.ShimName, dataSourceName)
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())
	if err := createSchema(db); err != nil {
		return nil, err
	}

	return &bunDB{db: db}, nil
}
