package database

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Feed struct {
	bun.BaseModel `bun:"table:feeds"`

	ID        int64     `bun:"id,pk,autoincrement"`
	URL       string    `bun:"url,unique,notnull"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

type ProcessedArticle struct {
	bun.BaseModel `bun:"table:processed_articles"`

	ID          int64     `bun:"id,pk,autoincrement"`
	ArticleGUID string    `bun:"article_guid,unique,notnull"`
	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

func createSchema(db *bun.DB) error {
	models := []any{
		(*Feed)(nil),
		(*ProcessedArticle)(nil),
	}

	for _, model := range models {
		if _, err := db.NewCreateTable().Model(model).IfNotExists().Exec(context.Background()); err != nil {
			return err
		}
	}

	return nil
}
