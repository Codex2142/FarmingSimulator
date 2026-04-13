package repository

import (
	"context"
	"farming/internal/model"

	"github.com/jackc/pgx/v5"
)

// CycleRepository mendefinisikan kontrak untuk operasi cycle di database
type CycleRepository interface {
	CreateCycle(ctx context.Context, cycle model.Cycle) (model.Cycle, error)
	GetCycleById(ctx context.Context, id int) (model.Cycle, error)
	UpdateCycle(ctx context.Context, cycle model.Cycle, id int) (model.Cycle, error)
	DeleteCycle(ctx context.Context, id int) (model.Cycle, error)
	GetAllCycles(ctx context.Context, limit, offset int) ([]model.Cycle, int, error)
}

// Implementasi CycleRepository menggunakan pgx.Conn
type cycleRepo struct {
	db *pgx.Conn
}

// Constructor NewCycleRepo
func NewCycleRepo(db *pgx.Conn) CycleRepository {
	return &cycleRepo{db: db}
}

// CreateCycle membuat cycle baru
func (r *cycleRepo) CreateCycle(ctx context.Context, cycle model.Cycle) (model.Cycle, error) {
	query := `
		INSERT INTO cycles(sub_place_id, commodity_type, commodity_name, start_date, end_date, status) 
		VALUES($1, $2, $3, $4, $5, $6) 
		RETURNING id, sub_place_id, commodity_type, commodity_name, start_date, end_date, status, created_at, updated_at
	`

	err := r.db.QueryRow(ctx, query, cycle.SubPlaceID, cycle.CommodityType, cycle.CommodityName, cycle.StartDate, cycle.EndDate, cycle.Status).Scan(
		&cycle.ID, &cycle.SubPlaceID, &cycle.CommodityType, &cycle.CommodityName, &cycle.StartDate, &cycle.EndDate, &cycle.Status, &cycle.CreatedAt, &cycle.UpdatedAt,
	)
	return cycle, err
}

// GetCycleById mengambil cycle berdasarkan ID
func (r *cycleRepo) GetCycleById(ctx context.Context, id int) (model.Cycle, error) {
	var cycle model.Cycle
	query := `
		SELECT id, sub_place_id, commodity_type, commodity_name, start_date, end_date, status, created_at, updated_at 
		FROM cycles 
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&cycle.ID, &cycle.SubPlaceID, &cycle.CommodityType, &cycle.CommodityName, &cycle.StartDate, &cycle.EndDate, &cycle.Status, &cycle.CreatedAt, &cycle.UpdatedAt,
	)

	return cycle, err
}

// UpdateCycle mengupdate cycle
func (r *cycleRepo) UpdateCycle(ctx context.Context, cycle model.Cycle, id int) (model.Cycle, error) {
	query := `
		UPDATE cycles 
		SET commodity_type = $1, commodity_name = $2, start_date = $3, end_date = $4, status = $5 
		WHERE id = $6 
		RETURNING id, sub_place_id, commodity_type, commodity_name, start_date, end_date, status, created_at, updated_at
	`

	err := r.db.QueryRow(ctx, query, cycle.CommodityType, cycle.CommodityName, cycle.StartDate, cycle.EndDate, cycle.Status, id).Scan(
		&cycle.ID, &cycle.SubPlaceID, &cycle.CommodityType, &cycle.CommodityName, &cycle.StartDate, &cycle.EndDate, &cycle.Status, &cycle.CreatedAt, &cycle.UpdatedAt,
	)

	return cycle, err
}

// DeleteCycle menghapus cycle
func (r *cycleRepo) DeleteCycle(ctx context.Context, id int) (model.Cycle, error) {
	query := `DELETE FROM cycles WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)
	return model.Cycle{}, err
}

// GetAllCycles mengambil semua cycle dengan pagination
func (r *cycleRepo) GetAllCycles(ctx context.Context, limit, offset int) ([]model.Cycle, int, error) {
	// Hitung total
	var total int
	countQuery := `SELECT COUNT(*) FROM cycles`

	err := r.db.QueryRow(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Query dengan pagination
	query := `
		SELECT id, sub_place_id, commodity_type, commodity_name, start_date, end_date, status, created_at, updated_at 
		FROM cycles 
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var cycles []model.Cycle

	for rows.Next() {
		var cycle model.Cycle

		err := rows.Scan(
			&cycle.ID, &cycle.SubPlaceID, &cycle.CommodityType, &cycle.CommodityName, &cycle.StartDate, &cycle.EndDate, &cycle.Status, &cycle.CreatedAt, &cycle.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}

		cycles = append(cycles, cycle)
	}

	return cycles, total, nil
}
