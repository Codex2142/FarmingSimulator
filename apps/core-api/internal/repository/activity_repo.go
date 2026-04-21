package repository

import (
	"context"
	"farming/internal/model"

	"github.com/jackc/pgx/v5"
)

// ActivityRepository mendefinisikan kontrak untuk operasi activity di database
type ActivityRepository interface {
	CreateActivity(ctx context.Context, activity model.Activity) (model.Activity, error)
	GetActivityById(ctx context.Context, id int) (model.Activity, error)
	UpdateActivity(ctx context.Context, activity model.Activity, id int) (model.Activity, error)
	DeleteActivity(ctx context.Context, id int) (model.Activity, error)
	GetAllActivities(ctx context.Context, limit, offset int) ([]model.Activity, int, error)
}

// Implementasi ActivityRepository menggunakan pgx.Conn
type activityRepo struct {
	db *pgx.Conn
}

// Constructor NewActivityRepo
func NewActivityRepo(db *pgx.Conn) ActivityRepository {
	return &activityRepo{db: db}
}

// CreateActivity membuat activity baru
func (r *activityRepo) CreateActivity(ctx context.Context, activity model.Activity) (model.Activity, error) {
	query := `
		INSERT INTO activities(cycle_id, type, description, quantity, unit, created_by) 
		VALUES($1, $2, $3, $4, $5, $6) 
		RETURNING id, cycle_id, type, description, quantity, unit, created_by, created_at
	`

	err := r.db.QueryRow(ctx, query, activity.CycleID, activity.Type, activity.Description, activity.Quantity, activity.Unit, activity.CreatedBy).Scan(
		&activity.ID, &activity.CycleID, &activity.Type, &activity.Description, &activity.Quantity, &activity.Unit, &activity.CreatedBy, &activity.CreatedAt,
	)
	return activity, err
}

// GetActivityById mengambil activity berdasarkan ID
func (r *activityRepo) GetActivityById(ctx context.Context, id int) (model.Activity, error) {
	var activity model.Activity
	query := `
		SELECT id, cycle_id, type, description, quantity, unit, created_by, created_at 
		FROM activities 
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&activity.ID, &activity.CycleID, &activity.Type, &activity.Description, &activity.Quantity, &activity.Unit, &activity.CreatedBy, &activity.CreatedAt,
	)

	return activity, err
}

// UpdateActivity mengupdate activity
func (r *activityRepo) UpdateActivity(ctx context.Context, activity model.Activity, id int) (model.Activity, error) {
	query := `
		UPDATE activities 
		SET type = $1, description = $2, quantity = $3, unit = $4 
		WHERE id = $5 
		RETURNING id, cycle_id, type, description, quantity, unit, created_by, created_at
	`

	err := r.db.QueryRow(ctx, query, activity.Type, activity.Description, activity.Quantity, activity.Unit, id).Scan(
		&activity.ID, &activity.CycleID, &activity.Type, &activity.Description, &activity.Quantity, &activity.Unit, &activity.CreatedBy, &activity.CreatedAt,
	)

	return activity, err
}

// DeleteActivity menghapus activity
func (r *activityRepo) DeleteActivity(ctx context.Context, id int) (model.Activity, error) {
	query := `DELETE FROM activities WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)
	return model.Activity{}, err
}

// GetAllActivities mengambil semua activity dengan pagination
func (r *activityRepo) GetAllActivities(ctx context.Context, limit, offset int) ([]model.Activity, int, error) {
	// Hitung total
	var total int
	countQuery := `SELECT COUNT(*) FROM activities`

	err := r.db.QueryRow(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Query dengan pagination
	query := `
		SELECT id, cycle_id, type, description, quantity, unit, created_by, created_at 
		FROM activities 
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var activities []model.Activity

	for rows.Next() {
		var activity model.Activity

		err := rows.Scan(
			&activity.ID, &activity.CycleID, &activity.Type, &activity.Description, &activity.Quantity, &activity.Unit, &activity.CreatedBy, &activity.CreatedAt,
		)
		if err != nil {
			return nil, 0, err
		}

		activities = append(activities, activity)
	}

	return activities, total, nil
}
