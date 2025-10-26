package db

import (
	"database/sql"
	"time"
)

type UrlEntity struct {
	Id          string
	ShortUrl    string
	OriginalUrl string
	CreatedAt   time.Time
	DeletedAt   sql.NullTime
}
