package db

import (
	"database/sql"
)

type UrlRepository struct {
	db *sql.DB
}

func NewUrlRepository(db *sql.DB) *UrlRepository {
	return &UrlRepository{
		db: db,
	}
}

func (u *UrlRepository) GetCounterAndIncrement() (int64, error) {
	var counter int64
	err := u.db.QueryRow("SELECT nextval('url_counter')").Scan(&counter)
	if err != nil {
		return 0, err
	}

	return counter, nil
}

func (u *UrlRepository) GetByShortUrl(url string) (e *UrlEntity, err error) {
	var entity UrlEntity
	err = u.db.QueryRow(
		"SELECT id, short_url, original_url, created_at, deleted_at FROM urls WHERE short_url = $1",
		url,
	).Scan(
		&entity.Id,
		&entity.ShortUrl,
		&entity.OriginalUrl,
		&entity.CreatedAt,
		&entity.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func CreateUrl(entity UrlEntity) error {
	_, err := dbConnection.Exec(
		"INSERT INTO urls (id, short_url, original_url, created_at) VALUES ($1, $2, $3, $4)",
		entity.Id,
		entity.ShortUrl,
		entity.OriginalUrl,
		entity.CreatedAt,
	)
	return err

}
