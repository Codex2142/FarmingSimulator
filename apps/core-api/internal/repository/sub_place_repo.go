package repository

import (
	"context"
	"farming/internal/model"

	"github.com/jackc/pgx/v5"
)

// SubPlaceRepository mendefinisikan kontrak untuk operasi sub_place di database
type SubPlaceRepository interface {
	CreateSubPlace(ctx context.Context, subPlace model.SubPlace) (model.SubPlace, error)
	GetSubPlaceById(ctx context.Context, id int) (model.SubPlace, error)
	UpdateSubPlace(ctx context.Context, subPlace model.SubPlace, id int) (model.SubPlace, error)
	DeleteSubPlace(ctx context.Context, id int) (model.SubPlace, error)
	GetAllSubPlaces(ctx context.Context, limit, offset int) ([]model.SubPlace, int, error)
}

// Implementasi SubPlaceRepository menggunakan pgx.Conn
type subPlaceRepo struct {
	db *pgx.Conn
}

// Constructor NewSubPlaceRepo
func NewSubPlaceRepo(db *pgx.Conn) SubPlaceRepository {
	return &subPlaceRepo{db: db}
}

// CreateSubPlace membuat sub_place baru
func (r *subPlaceRepo) CreateSubPlace(ctx context.Context, subPlace model.SubPlace) (model.SubPlace, error) {
	query := `
		INSERT INTO sub_places(farm_id, name, type, size, status) 
		VALUES($1, $2, $3, $4, $5) 
		RETURNING id, farm_id, name, type, size, status, created_at, updated_at
	`

	err := r.db.QueryRow(ctx, query, subPlace.FarmID, subPlace.Name, subPlace.Type, subPlace.Size, subPlace.Status).Scan(
		&subPlace.ID, &subPlace.FarmID, &subPlace.Name, &subPlace.Type, &subPlace.Size, &subPlace.Status, &subPlace.CreatedAt, &subPlace.UpdatedAt,
	)
	return subPlace, err
}

// GetSubPlaceById mengambil sub_place berdasarkan ID
func (r *subPlaceRepo) GetSubPlaceById(ctx context.Context, id int) (model.SubPlace, error) {
	var subPlace model.SubPlace
	query := `
		SELECT id, farm_id, name, type, size, status, created_at, updated_at 
		FROM sub_places 
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&subPlace.ID, &subPlace.FarmID, &subPlace.Name, &subPlace.Type, &subPlace.Size, &subPlace.Status, &subPlace.CreatedAt, &subPlace.UpdatedAt,
	)

	return subPlace, err
}

// UpdateSubPlace mengupdate sub_place
func (r *subPlaceRepo) UpdateSubPlace(ctx context.Context, subPlace model.SubPlace, id int) (model.SubPlace, error) {
	query := `
		UPDATE sub_places 
		SET name = $1, type = $2, size = $3, status = $4 
		WHERE id = $5 
		RETURNING id, farm_id, name, type, size, status, created_at, updated_at
	`

	err := r.db.QueryRow(ctx, query, subPlace.Name, subPlace.Type, subPlace.Size, subPlace.Status, id).Scan(
		&subPlace.ID, &subPlace.FarmID, &subPlace.Name, &subPlace.Type, &subPlace.Size, &subPlace.Status, &subPlace.CreatedAt, &subPlace.UpdatedAt,
	)

	return subPlace, err
}

// DeleteSubPlace menghapus sub_place
func (r *subPlaceRepo) DeleteSubPlace(ctx context.Context, id int) (model.SubPlace, error) {
	query := `DELETE FROM sub_places WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)
	return model.SubPlace{}, err
}

// GetAllSubPlaces mengambil semua sub_place dengan pagination
func (r *subPlaceRepo) GetAllSubPlaces(ctx context.Context, limit, offset int) ([]model.SubPlace, int, error) {
	// Hitung total
	var total int
	countQuery := `SELECT COUNT(*) FROM sub_places`

	err := r.db.QueryRow(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Query dengan pagination
	query := `
		SELECT id, farm_id, name, type, size, status, created_at, updated_at 
		FROM sub_places 
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var subPlaces []model.SubPlace

	for rows.Next() {
		var subPlace model.SubPlace

		err := rows.Scan(
			&subPlace.ID, &subPlace.FarmID, &subPlace.Name, &subPlace.Type, &subPlace.Size, &subPlace.Status, &subPlace.CreatedAt, &subPlace.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}

		subPlaces = append(subPlaces, subPlace)
	}

	return subPlaces, total, nil
}
