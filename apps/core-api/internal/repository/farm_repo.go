package repository

import (
	"context"
	"farming/internal/model"

	"github.com/jackc/pgx/v5"
)

type FarmRepository interface {
	CreateFarm(ctx context.Context, farm model.Farm) (model.Farm, error)
	GetFarmById(ctx context.Context, id int) (model.Farm, error)
	UpdateFarm(ctx context.Context, farm model.Farm, id int) (model.Farm, error)
	DeleteFarm(ctx context.Context, id int) (model.Farm, error)
	GetAllFarms(ctx context.Context, limit, offset int) ([]model.Farm, int, error)
}

type farmRepo struct {
	db *pgx.Conn
}

func NewFarmRepo(db *pgx.Conn) FarmRepository {
	return &farmRepo{db: db}
}

func (r *farmRepo) CreateFarm(ctx context.Context, farm model.Farm) (model.Farm, error) {
	query := `INSERT INTO farms (name, location, leader_id) VALUES($1, $2, $3) RETURNING id, name, location, leader_id`

	// Scan digunakan untuk menyalin hasil RETURNING ke struct farm
	err := r.db.QueryRow(ctx, query, farm.Name, farm.Location, farm.LeaderID).Scan(
		&farm.ID, &farm.Name, &farm.Location, &farm.LeaderID,
	)

	return farm, err
}

func (r *farmRepo) GetFarmById(ctx context.Context, id int) (model.Farm, error) {
	var farm model.Farm
	query := `
		SELECT 
			f.id, 
			f.name, 
			f.location, 
			f.leader_id,
			u.id,
			u.name,
			u.phone
		FROM farms f
		LEFT JOIN users u ON f.leader_id = u.id
		WHERE f.id = $1
		`

	var user model.User
	// Gunakan pointer untuk nullable fields dari LEFT JOIN
	var userID *int
	var userName *string
	var userPhone *string

	err := r.db.QueryRow(ctx, query, id).Scan(
		&farm.ID,
		&farm.Name,
		&farm.Location,
		&farm.LeaderID,
		&userID,
		&userName,
		&userPhone,
	)

	// Hanya assign leader jika data user valid (tidak NULL dari LEFT JOIN)
	if farm.LeaderID != nil && userID != nil {
		user.ID = *userID
		if userName != nil {
			user.Name = *userName
		}
		if userPhone != nil {
			user.Phone = *userPhone
		}
		farm.Leader = &user
	}
	return farm, err
}

func (r *farmRepo) UpdateFarm(ctx context.Context, farm model.Farm, id int) (model.Farm, error) {

	query := `UPDATE farms SET name = $1, location = $2, leader_id = $3 WHERE id = $4 RETURNING id, name, location, leader_id, updated_at`
	err := r.db.QueryRow(ctx, query, farm.Name, farm.Location, farm.LeaderID, id).Scan(
		&farm.ID, &farm.Name, &farm.Location, &farm.LeaderID, &farm.UpdatedAt,
	)

	return farm, err
}

func (r *farmRepo) DeleteFarm(ctx context.Context, id int) (model.Farm, error) {

	query := `DELETE FROM farms WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)
	return model.Farm{}, err
}

func (r *farmRepo) GetAllFarms(ctx context.Context, limit, offset int) ([]model.Farm, int, error) {

	// ambil total keseluruhan data
	var total int
	countQuery := `SELECT COUNT(*) FROM farms`

	err := r.db.QueryRow(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	query := `
		SELECT 
			f.id, 
			f.name, 
			f.location, 
			f.leader_id, 
			u.id,
			u.name, 
			u.phone 
		FROM 
			farms AS f 
		LEFT JOIN users AS u 
		ON f.leader_id=u.id
		LIMIT $1 OFFSET $2
		`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var farms []model.Farm

	for rows.Next() {
		var farm model.Farm
		var user model.User

		// Gunakan pointer untuk nullable fields dari LEFT JOIN
		var userID *int
		var userName *string
		var userPhone *string

		err := rows.Scan(
			&farm.ID,
			&farm.Name,
			&farm.Location,
			&farm.LeaderID,
			&userID,
			&userName,
			&userPhone,
		)

		if err != nil {
			return nil, 0, err
		}

		// Hanya assign leader jika data user valid (tidak NULL dari LEFT JOIN)
		if farm.LeaderID != nil && userID != nil {
			user.ID = *userID
			if userName != nil {
				user.Name = *userName
			}
			if userPhone != nil {
				user.Phone = *userPhone
			}
			farm.Leader = &user
		}

		farms = append(farms, farm)
	}

	return farms, total, nil

}
