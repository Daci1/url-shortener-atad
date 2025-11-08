package db

import (
	"database/sql"
	"errors"
	"fmt"
	"net"

	"github.com/Daci1/url-shortener-atad/internal/errs"
	"github.com/Daci1/url-shortener-atad/internal/helper"
	"github.com/Daci1/url-shortener-atad/internal/shortener"
)

type UrlRepository struct {
	db *sql.DB
}

func NewUrlRepository(db *sql.DB) *UrlRepository {
	return &UrlRepository{
		db: db,
	}
}

func (r *UrlRepository) getCounterAndIncrement() (int64, errs.CustomError) {
	var counter int64
	err := r.db.QueryRow("SELECT nextval('url_counter')").Scan(&counter)
	if err != nil {
		return 0, errs.Internal(fmt.Sprintf("Error getting next counter value: %s", err))
	}

	return counter, nil
}

func (r *UrlRepository) GetNextAvailableShortUrl() (string, errs.CustomError) {
	for {
		counter, err := r.getCounterAndIncrement()
		if err != nil {
			return "", err
		}

		shortUrl := shortener.ToBase62(counter)

		// Check if the short URL already exists (user might have taken it as a custom one)
		exists, err := r.ShortUrlExists(shortUrl)
		if err != nil {
			return "", err
		}
		if !exists {
			return shortUrl, nil
		}
	}
}

func (r *UrlRepository) ShortUrlExists(shortUrl string) (bool, errs.CustomError) {
	var exists bool
	err := r.db.QueryRow(`
		SELECT EXISTS(SELECT 1 FROM urls WHERE short_url = $1)
	`, shortUrl).Scan(&exists)

	if err != nil {
		return false, errs.Internal(fmt.Sprintf("Error querying if short url exists: %s", err))
	}
	return exists, nil
}

func (r *UrlRepository) GetByShortUrlAndIncrementAnalytics(url, requestIp string) (*UrlEntity, errs.CustomError) {
	var entity UrlEntity
	err := r.db.QueryRow(
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound("Short url not found")
		}
		return nil, errs.Internal(fmt.Sprintf("Error querying urls: %s", err))
	}

	_, err = r.db.Exec(`
		INSERT INTO analytics (url_id, visited_count)
		VALUES ($1, 1)
		ON CONFLICT (url_id)
		DO UPDATE SET visited_count = analytics.visited_count + 1
	`, entity.Id)

	if err != nil {
		return &entity, errs.Internal(fmt.Sprintf("Error incrementing analytics: %s", err))
	}

	if requestIp != "" {
		err = r.upsertUniqueVisitor(&entity, requestIp)
		if err != nil {
			return &entity, errs.Internal(fmt.Sprintf("Error registering unique visitor: %s", err))
		}
		if !helper.IsPrivateIP(net.ParseIP(requestIp)) {
			err = r.upsertGeolocation(&entity, requestIp)
			if err != nil {
				return &entity, errs.Internal(fmt.Sprintf("Error registering geolocation: %s", err))
			}
		}
	}

	return &entity, nil
}

func (r *UrlRepository) upsertUniqueVisitor(urlEntity *UrlEntity, requestIp string) errs.CustomError {
	_, err := r.db.Exec(`
		INSERT INTO unique_visitors (url_id, visitor_ip, visited_count)
		VALUES ($1, $2, 1)
		ON CONFLICT (url_id)
		DO UPDATE SET visited_count = unique_visitors.visited_count + 1,
		              visitor_ip = EXCLUDED.visitor_ip
		`,
		urlEntity.Id,
		requestIp,
	)

	if err != nil {
		return errs.Internal("Failed to persist visitor data")
	}
	return nil
}

func (r *UrlRepository) upsertGeolocation(urlEntity *UrlEntity, requestIp string) errs.CustomError {
	geoData, err := helper.GetGeoData(requestIp)
	if geoData == nil {
		return errs.Internal("Failed to retrieve geolocation data")
	}
	_, err = r.db.Exec(`
	INSERT INTO ip_locations (
		city, region, country, latitude, longitude, url_id, visited_counter
	)
	VALUES ($1, $2, $3, $4, $5, $6, 1)
	ON CONFLICT (city, region, country, latitude, longitude, url_id)
	DO UPDATE SET
		visited_counter = ip_locations.visited_counter + 1
	`,
		geoData.City,
		geoData.Region,
		geoData.Country,
		geoData.Latitude,
		geoData.Longitude,
		urlEntity.Id,
	)

	if err != nil {
		return errs.Internal("Failed to persist geolocation data")
	}
	return nil
}

func (r *UrlRepository) CreateUrl(entity UrlEntity) errs.CustomError {
	_, err := r.db.Exec(
		"INSERT INTO urls (id, short_url, original_url, created_at) VALUES ($1, $2, $3, $4)",
		entity.Id,
		entity.ShortUrl,
		entity.OriginalUrl,
		entity.CreatedAt,
	)
	if err != nil {
		return errs.Internal(fmt.Sprintf("Error creating url: %s", err))
	}
	return nil
}

func (r *UrlRepository) CreateUrlWithUser(entity UrlEntity) errs.CustomError {
	_, err := r.db.Exec(
		"INSERT INTO urls (id, short_url, original_url, user_id, created_at) VALUES ($1, $2, $3, $4, $5)",
		entity.Id,
		entity.ShortUrl,
		entity.OriginalUrl,
		entity.UserId,
		entity.CreatedAt,
	)
	if err != nil {
		return errs.Internal(fmt.Sprintf("Error creating url with user: %s", err))
	}

	return nil
}
