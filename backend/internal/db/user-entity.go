package db

import (
	"database/sql"
	"time"
)

type UserEntity struct {
	Id           string
	Email        string
	Username     string
	PasswordHash string
	CreatedAt    time.Time
	DeletedAt    sql.NullTime
}
