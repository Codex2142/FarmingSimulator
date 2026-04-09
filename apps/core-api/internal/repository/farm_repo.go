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
}

type farmRepo struct {
	db *pgx.Conn
}

func NewFarmRepo(db *pgx.Conn) FarmRepository {
	return &farmRepo{db: db}
}

func (r *farmRepo) CreateFarm(ctx context.Context, farm model.Farm) (model.Farm, error) {
	query := `INSERT INTO farms (name, location, leader_id) VALUES($1, $2, $3) RETURNING id, name, location, leader_id, created_at`

	// Scan digunakan untuk menyalin hasil RETURNING ke struct farm
	err := r.db.QueryRow(ctx, query, farm.Name, farm.Location, farm.LeaderID).Scan(
		&farm.ID, &farm.Name, &farm.Location, &farm.LeaderID,
	)

	return farm, err
}

func (r *farmRepo) GetFarmById(ctx context.Context, id int) (model.Farm, error) {
	var farm model.Farm
	query := `SELECT id, name, location, leader_id, created_at FROM farms WHERE id=$1`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&farm.ID, &farm.Name, &farm.Location, &farm.LeaderID,
	)

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
