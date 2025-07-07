package database

import "context"

type contextKey string

const dbKey = contextKey("db")

func ToCtx(ctx context.Context, db DB) context.Context {
	return context.WithValue(ctx, dbKey, db)
}

func FromCtx(ctx context.Context) (DB, bool) {
	db, ok := ctx.Value(dbKey).(DB)
	return db, ok
}
