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

func (r *UrlRepository) GetCounterAndIncrement() (int64, error) {
	var counter int64
	err := r.db.QueryRow("SELECT nextval('url_counter')").Scan(&counter)
	if err != nil {
		return 0, err
	}

	return counter, nil
}

// TODO: adjust errors to new type
func (r *UrlRepository) GetByShortUrl(url string) (e *UrlEntity, err error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var entity UrlEntity
	err = r.db.QueryRow(
		"SELECT id, short_url, original_url, created_at, deleted_at FROM urls WHERE short_url = $1",
		url,
	).Scan(
		&entity.Id,
		&entity.ShortUrl,
		&entity.OriginalUrl,
		&entity.CreatedAt,
		&entity.DeletedAt,
	)
	_, err = tx.Exec(`
		INSERT INTO analytics (url_id, visited_count)
		VALUES ($1, 1)
		ON CONFLICT (url_id)
		DO UPDATE SET visited_count = analytics.visited_count + 1
	`, entity.Id)

	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *UrlRepository) CreateUrl(entity UrlEntity) error {
	_, err := r.db.Exec(
		"INSERT INTO urls (id, short_url, original_url, created_at) VALUES ($1, $2, $3, $4)",
		entity.Id,
		entity.ShortUrl,
		entity.OriginalUrl,
		entity.CreatedAt,
	)
	return err

}
