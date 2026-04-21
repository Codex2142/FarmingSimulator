package usecase

import (
	"context"
	"farming/internal/dto"
	"farming/internal/model"
	"farming/internal/repository"
	"farming/pkg/utils"
)

// ActivityUsecase mendefinisikan kontrak logika bisnis untuk activity
type ActivityUsecase interface {
	CreateActivity(ctx context.Context, req dto.CreateActivityRequest) (dto.ActivityResponse, error)
	GetActivityById(ctx context.Context, id int) (dto.ActivityResponse, error)
	UpdateActivity(ctx context.Context, id int, req dto.UpdateActivityRequest) (dto.ActivityResponse, error)
	DeleteActivity(ctx context.Context, id int) error
	GetAllActivities(ctx context.Context, page, limit int) ([]dto.ActivityResponse, dto.PaginationResponse, error)
}

// Implementasi ActivityUsecase
type activityUsecase struct {
	activityRepo repository.ActivityRepository
	cycleRepo    repository.CycleRepository
	userRepo     repository.UserRepository
}

// Constructor NewActivityUsecase
func NewActivityUsecase(
	activityRepo repository.ActivityRepository,
	cycleRepo repository.CycleRepository,
	userRepo repository.UserRepository,
) ActivityUsecase {
	return &activityUsecase{
		activityRepo: activityRepo,
		cycleRepo:    cycleRepo,
		userRepo:     userRepo,
	}
}

// CreateActivity membuat activity baru
func (u *activityUsecase) CreateActivity(ctx context.Context, req dto.CreateActivityRequest) (dto.ActivityResponse, error) {
	// Validasi cycle ada
	_, err := u.cycleRepo.GetCycleById(ctx, req.CycleID)
	if err != nil {
		return dto.ActivityResponse{}, err
	}

	// Validasi user/createdBy ada
	_, err = u.userRepo.GetUserById(ctx, req.CreatedBy)
	if err != nil {
		return dto.ActivityResponse{}, err
	}

	// Buat model
	activity := model.Activity{
		CycleID:     req.CycleID,
		Type:        req.Type,
		Description: req.Description,
		Quantity:    req.Quantity,
		Unit:        req.Unit,
		CreatedBy:   req.CreatedBy,
	}

	// Save ke database
	createdActivity, err := u.activityRepo.CreateActivity(ctx, activity)
	if err != nil {
		return dto.ActivityResponse{}, err
	}

	// Convert ke response
	return dto.ActivityResponse{
		ID:          createdActivity.ID,
		CycleID:     createdActivity.CycleID,
		Type:        createdActivity.Type,
		Description: createdActivity.Description,
		Quantity:    createdActivity.Quantity,
		Unit:        createdActivity.Unit,
		CreatedBy:   createdActivity.CreatedBy,
		CreatedAt:   createdActivity.CreatedAt.String(),
	}, nil
}

// GetActivityById mengambil activity berdasarkan ID
func (u *activityUsecase) GetActivityById(ctx context.Context, id int) (dto.ActivityResponse, error) {
	activity, err := u.activityRepo.GetActivityById(ctx, id)
	if err != nil {
		return dto.ActivityResponse{}, err
	}

	return dto.ActivityResponse{
		ID:          activity.ID,
		CycleID:     activity.CycleID,
		Type:        activity.Type,
		Description: activity.Description,
		Quantity:    activity.Quantity,
		Unit:        activity.Unit,
		CreatedBy:   activity.CreatedBy,
		CreatedAt:   activity.CreatedAt.String(),
	}, nil
}

// UpdateActivity mengupdate activity
func (u *activityUsecase) UpdateActivity(ctx context.Context, id int, req dto.UpdateActivityRequest) (dto.ActivityResponse, error) {
	activity := model.Activity{
		Type:        req.Type,
		Description: req.Description,
		Quantity:    req.Quantity,
		Unit:        req.Unit,
	}

	updatedActivity, err := u.activityRepo.UpdateActivity(ctx, activity, id)
	if err != nil {
		return dto.ActivityResponse{}, err
	}

	return dto.ActivityResponse{
		ID:          updatedActivity.ID,
		CycleID:     updatedActivity.CycleID,
		Type:        updatedActivity.Type,
		Description: updatedActivity.Description,
		Quantity:    updatedActivity.Quantity,
		Unit:        updatedActivity.Unit,
		CreatedBy:   updatedActivity.CreatedBy,
		CreatedAt:   updatedActivity.CreatedAt.String(),
	}, nil
}

// DeleteActivity menghapus activity
func (u *activityUsecase) DeleteActivity(ctx context.Context, id int) error {
	_, err := u.activityRepo.DeleteActivity(ctx, id)
	return err
}

// GetAllActivities mengambil semua activity dengan pagination
func (u *activityUsecase) GetAllActivities(ctx context.Context, page, limit int) ([]dto.ActivityResponse, dto.PaginationResponse, error) {
	limit, offset := utils.CalculatePagination(page, limit)

	activities, total, err := u.activityRepo.GetAllActivities(ctx, limit, offset)
	if err != nil {
		return nil, dto.PaginationResponse{}, err
	}

	var result []dto.ActivityResponse

	for _, activity := range activities {
		result = append(result, dto.ActivityResponse{
			ID:          activity.ID,
			CycleID:     activity.CycleID,
			Type:        activity.Type,
			Description: activity.Description,
			Quantity:    activity.Quantity,
			Unit:        activity.Unit,
			CreatedBy:   activity.CreatedBy,
			CreatedAt:   activity.CreatedAt.String(),
		})
	}

	totalPages := (total + limit - 1) / limit

	meta := dto.PaginationResponse{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}

	return result, meta, nil
}
