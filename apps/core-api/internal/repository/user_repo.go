package repository

import (
	"context"
	"farming/internal/model"

	"github.com/jackc/pgx/v5"
)

// ====================================================================
// Menentukan kontrak yang harus dimiliki repository user
type UserRepository interface {
	// Create membuat user baru di database
	// Menerima context (untuk timeout/cancel) dan objek user
	// Mengembalikan user yang sudah dibuat beserta error jika ada
	Create(ctx context.Context, user model.User) (model.User, error)
	GetbyID(ctx context.Context, id int) (model.User, error)
	Update(ctx context.Context, user model.User, id int) (model.User, error)
	Delete(ctx context.Context, id int) (model.User, error)
}

// ====================================================================
// Implementasi UserRepository menggunakan pgx.Conn
type userRepo struct {
	db *pgx.Conn // koneksi database PostgreSQL
}

// ====================================================================
// Constructor NewUserRepo
// Mengembalikan instance userRepo sebagai UserRepository
func NewUserRepo(db *pgx.Conn) UserRepository {
	return &userRepo{db: db}
}

// ====================================================================
// Implementasi Create
func (r *userRepo) Create(ctx context.Context, user model.User) (model.User, error) {
	query := `INSERT INTO users(name, phone, password) VALUES($1, $2, $3) RETURNING id, name, phone, created_at`

	// Scan digunakan untuk menyalin hasil RETURNING ke struct user
	err := r.db.QueryRow(ctx, query, user.Name, user.Phone, user.Password).Scan(
		&user.ID, &user.Name, &user.Phone, &user.CreatedAt,
	)
	return user, err
}

// ====================================================================
// Implementasi GetbyID
func (r *userRepo) GetbyID(ctx context.Context, id int) (model.User, error) {
	var user model.User
	query := `SELECT id, name, phone, created_at FROM users WHERE id=$1`

	// Scan menyalin hasil query ke struct user
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.Name, &user.Phone, &user.CreatedAt,
	)

	return user, err
}

// ====================================================================
// Implementasi Update User
func (r *userRepo) Update(ctx context.Context, user model.User, id int) (model.User, error) {

	// membuat query kosong
	query := `UPDATE users SET name = $1, phone = $2, password = $3 WHERE id = $4 RETURNING id, name, phone, updated_at`
	err := r.db.QueryRow(ctx, query, user.Name, user.Phone, user.Password, id).Scan(
		&user.ID, &user.Name, &user.Phone, &user.UpdatedAt,
	)

	return user, err
}

// ====================================================================
func (r *userRepo) Delete(ctx context.Context, id int) (model.User, error) {

	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)
	return model.User{}, err
}
