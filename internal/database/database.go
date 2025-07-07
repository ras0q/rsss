package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Create feeds table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS feeds (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		url TEXT NOT NULL UNIQUE
	)`)
	if err != nil {
		return nil, err
	}

	// Create processed_articles table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS processed_articles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		article_guid TEXT NOT NULL UNIQUE
	)`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AddFeed(db *sql.DB, url string) error {
	_, err := db.Exec("INSERT INTO feeds (url) VALUES (?)", url)
	return err
}

func GetFeeds(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT url FROM feeds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []string
	for rows.Next() {
		var url string
		if err := rows.Scan(&url); err != nil {
			return nil, err
		}

		urls = append(urls, url)
	}

	return urls, nil
}

func IsArticleProcessed(db *sql.DB, guid string) (bool, error) {
	var exists bool

	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM processed_articles WHERE article_guid = ?)", guid).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	return exists, nil
}

func MarkArticleAsProcessed(db *sql.DB, guid string) error {
	_, err := db.Exec("INSERT INTO processed_articles (article_guid) VALUES (?)", guid)
	return err
}
