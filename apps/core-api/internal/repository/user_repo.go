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
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	GetUserById(ctx context.Context, id int) (model.User, error)
	UpdateUser(ctx context.Context, user model.User, id int) (model.User, error)
	DeleteUser(ctx context.Context, id int) (model.User, error)
	GetAllUsers(ctx context.Context, limit, offset int) ([]model.User, int, error)
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
func (r *userRepo) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	query := `INSERT INTO users(name, phone, password, role) VALUES($1, $2, $3, $4) RETURNING id, name, phone, role, created_at`

	// Scan digunakan untuk menyalin hasil RETURNING ke struct user
	err := r.db.QueryRow(ctx, query, user.Name, user.Phone, user.Password, user.Role).Scan(
		&user.ID, &user.Name, &user.Phone, &user.Role, &user.CreatedAt,
	)
	return user, err
}

// ====================================================================
// Implementasi GetbyID
func (r *userRepo) GetUserById(ctx context.Context, id int) (model.User, error) {
	var user model.User
	query := `SELECT id, name, phone, role, created_at FROM users WHERE id=$1`

	// Scan menyalin hasil query ke struct user
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.Name, &user.Phone, &user.Role, &user.CreatedAt,
	)

	return user, err
}

// ====================================================================
// Implementasi Update User
func (r *userRepo) UpdateUser(ctx context.Context, user model.User, id int) (model.User, error) {

	// membuat query kosong
	query := `UPDATE users SET name = $1, phone = $2, role = $3 WHERE id = $4 RETURNING id, name, phone, role, updated_at`
	err := r.db.QueryRow(ctx, query, user.Name, user.Phone, user.Role, id).Scan(
		&user.ID, &user.Name, &user.Phone, &user.Role, &user.UpdatedAt,
	)

	return user, err
}

// ====================================================================
func (r *userRepo) DeleteUser(ctx context.Context, id int) (model.User, error) {

	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)
	return model.User{}, err
}

func (r *userRepo) GetAllUsers(ctx context.Context, limit, offset int) ([]model.User, int, error) {

	// ambil total keseluruhan data
	var total int
	countQuery := `SELECT COUNT(*) FROM users`

	err := r.db.QueryRow(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Query ambil semua user
	query := `SELECT id, name, phone, role, created_at, updated_at FROM users`

	// Eksekusi query (karena banyak data pakai Query, bukan QueryRow)
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close() // penting untuk mencegah memory leak

	var users []model.User

	// Loop setiap baris hasil query
	for rows.Next() {
		var user model.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Phone,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}

		users = append(users, user)
	}

	return users, total, nil
}
