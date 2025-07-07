package database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func NewMySQLDB(dsn string) (DB, error) {
	config, err := mysql.ParseDSN(dsn)
	if err != nil {
		return nil, err
	}

	config.ParseTime = true

	sqldb, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, mysqldialect.New())
	if err := createSchema(db); err != nil {
		return nil, err
	}

	return &bunDB{db: db}, nil
}
