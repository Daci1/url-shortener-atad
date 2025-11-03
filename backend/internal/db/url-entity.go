package db

import (
	"database/sql"
	"time"
)

type UrlEntity struct {
	Id          string
	ShortUrl    string
	OriginalUrl string
	UserId      string
	CreatedAt   time.Time
	DeletedAt   sql.NullTime
}
