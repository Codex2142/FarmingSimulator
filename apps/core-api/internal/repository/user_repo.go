package repository

import (
	"context"
	"farming/internal/model"

	"github.com/jackc/pgx/v5"
)

// Repository ini pure DB access, tidak ada logic bisnis.

type UserRepository interface {
	Create(ctx context.Context, user model.User) (model.User, error)
	GetbyID(ctx context.Context, id int) (model.User, error)
}

type userRepo struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user model.User) (model.User, error) {
	query := `INSERT INTO users(name, phone) VALUES($1, $2) RETURNING id, name, phone, created_at`
	err := r.db.QueryRow(ctx, query, user.Name, user.Phone).Scan(
		&user.ID, &user.Name, &user.Phone, &user.CreatedAt,
	)
	return user, err
}

func (r *userRepo) GetbyID(ctx context.Context, id int) (model.User, error) {
	var user model.User
	query := `SELECT id, name, phone, created_at FROM users WHERE id=$1`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.Name, &user.Phone, &user.CreatedAt,
	)

	return user, err
}
