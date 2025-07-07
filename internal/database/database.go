package database

import (
	"context"

	"github.com/uptrace/bun"
)

type DB interface {
	AddFeed(ctx context.Context, url string) error
	GetFeeds(ctx context.Context) ([]string, error)
	IsArticleProcessed(ctx context.Context, guid string) (bool, error)
	MarkArticleAsProcessed(ctx context.Context, guid string) error
	Close() error
}

type bunDB struct {
	db *bun.DB
}

func (b *bunDB) AddFeed(ctx context.Context, url string) error {
	_, err := b.db.NewInsert().Model(&Feed{URL: url}).Exec(ctx)
	return err
}

func (b *bunDB) GetFeeds(ctx context.Context) ([]string, error) {
	var feeds []Feed
	if err := b.db.NewSelect().Model(&feeds).Scan(ctx); err != nil {
		return nil, err
	}

	urls := make([]string, len(feeds))
	for i, f := range feeds {
		urls[i] = f.URL
	}

	return urls, nil
}

func (b *bunDB) IsArticleProcessed(ctx context.Context, guid string) (bool, error) {
	exists, err := b.db.NewSelect().Model((*ProcessedArticle)(nil)).Where("article_guid = ?", guid).Exists(ctx)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (b *bunDB) MarkArticleAsProcessed(ctx context.Context, guid string) error {
	_, err := b.db.NewInsert().Model(&ProcessedArticle{ArticleGUID: guid}).Exec(ctx)
	return err
}

func (b *bunDB) Close() error {
	return b.db.Close()
}

