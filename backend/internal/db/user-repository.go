package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Daci1/url-shortener-atad/internal/errs"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) RegisterUser(entity *UserEntity) (*UserEntity, errs.CustomError) {
	_, err := r.db.Query(
		"INSERT INTO users (id, email, username, password_hash, created_at) VALUES ($1, $2, $3, $4, $5)",
		entity.Id,
		entity.Email,
		entity.Username,
		entity.PasswordHash,
		entity.CreatedAt,
	)
	if err != nil {
		// Try to unwrap pgx error
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				// Unique constraint violation
				return nil, errs.Conflict("user already exists")
			}
		}

		return nil, errs.Internal(fmt.Sprintf("failed to create user: %s", err.Error()))
	}

	return entity, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*UserEntity, errs.CustomError) {
	var entity UserEntity

	err := r.db.QueryRow(
		"SELECT id, email, username, password_hash, created_at FROM users WHERE email = $1 AND deleted_at IS NULL",
		email,
	).Scan(
		&entity.Id,
		&entity.Email,
		&entity.Username,
		&entity.PasswordHash,
		&entity.CreatedAt,
	)

	// TODO: test for not found user
	if err != nil {
		return nil, errs.Internal(fmt.Sprintf("Error when retrieving user %s: %s", email, err.Error()))

	}

	return &entity, nil
}
