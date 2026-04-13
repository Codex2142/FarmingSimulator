package usecase

import (
	"context"
	"farming/internal/dto"
	"farming/internal/model"
	"farming/internal/repository"
	"farming/pkg/constants"
	"farming/pkg/utils"
	"time"
)

// CycleUsecase mendefinisikan kontrak logika bisnis untuk cycle
type CycleUsecase interface {
	CreateCycle(ctx context.Context, req dto.CreateCycleRequest) (dto.CycleResponse, error)
	GetCycleById(ctx context.Context, id int) (dto.CycleResponse, error)
	UpdateCycle(ctx context.Context, id int, req dto.UpdateCycleRequest) (dto.CycleResponse, error)
	DeleteCycle(ctx context.Context, id int) error
	GetAllCycles(ctx context.Context, page, limit int) ([]dto.CycleResponse, dto.PaginationResponse, error)
}

// Implementasi CycleUsecase
type cycleUsecase struct {
	cycleRepo    repository.CycleRepository
	subPlaceRepo repository.SubPlaceRepository
}

// Constructor NewCycleUsecase
func NewCycleUsecase(cycleRepo repository.CycleRepository, subPlaceRepo repository.SubPlaceRepository) CycleUsecase {
	return &cycleUsecase{
		cycleRepo:    cycleRepo,
		subPlaceRepo: subPlaceRepo,
	}
}

// CreateCycle membuat cycle baru
func (u *cycleUsecase) CreateCycle(ctx context.Context, req dto.CreateCycleRequest) (dto.CycleResponse, error) {
	// Validasi sub_place ada
	_, err := u.subPlaceRepo.GetSubPlaceById(ctx, req.SubPlaceID)
	if err != nil {
		return dto.CycleResponse{}, err
	}

	// Parse start_date
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return dto.CycleResponse{}, err
	}

	// Set default status jika kosong
	if req.Status == "" {
		req.Status = string(constants.Ongoing)
	}

	// Buat model
	cycle := model.Cycle{
		SubPlaceID:    req.SubPlaceID,
		CommodityType: req.CommodityType,
		CommodityName: req.CommodityName,
		StartDate:     startDate,
		Status:        req.Status,
	}

	// Save ke database
	createdCycle, err := u.cycleRepo.CreateCycle(ctx, cycle)
	if err != nil {
		return dto.CycleResponse{}, err
	}

	// Convert ke response
	return dto.CycleResponse{
		ID:            createdCycle.ID,
		SubPlaceID:    createdCycle.SubPlaceID,
		CommodityType: createdCycle.CommodityType,
		CommodityName: createdCycle.CommodityName,
		StartDate:     createdCycle.StartDate,
		EndDate:       createdCycle.EndDate,
		Status:        createdCycle.Status,
		CreatedAt:     createdCycle.CreatedAt.String(),
		UpdatedAt:     createdCycle.UpdatedAt.String(),
	}, nil
}

// GetCycleById mengambil cycle berdasarkan ID
func (u *cycleUsecase) GetCycleById(ctx context.Context, id int) (dto.CycleResponse, error) {
	cycle, err := u.cycleRepo.GetCycleById(ctx, id)
	if err != nil {
		return dto.CycleResponse{}, err
	}

	return dto.CycleResponse{
		ID:            cycle.ID,
		SubPlaceID:    cycle.SubPlaceID,
		CommodityType: cycle.CommodityType,
		CommodityName: cycle.CommodityName,
		StartDate:     cycle.StartDate,
		EndDate:       cycle.EndDate,
		Status:        cycle.Status,
		CreatedAt:     cycle.CreatedAt.String(),
		UpdatedAt:     cycle.UpdatedAt.String(),
	}, nil
}

// UpdateCycle mengupdate cycle
func (u *cycleUsecase) UpdateCycle(ctx context.Context, id int, req dto.UpdateCycleRequest) (dto.CycleResponse, error) {
	// Parse start_date
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return dto.CycleResponse{}, err
	}

	// Parse end_date jika ada
	var endDate *time.Time
	if req.EndDate != nil {
		parsed, err := time.Parse("2006-01-02", *req.EndDate)
		if err != nil {
			return dto.CycleResponse{}, err
		}
		endDate = &parsed
	}

	cycle := model.Cycle{
		CommodityType: req.CommodityType,
		CommodityName: req.CommodityName,
		StartDate:     startDate,
		EndDate:       endDate,
		Status:        req.Status,
	}

	updatedCycle, err := u.cycleRepo.UpdateCycle(ctx, cycle, id)
	if err != nil {
		return dto.CycleResponse{}, err
	}

	return dto.CycleResponse{
		ID:            updatedCycle.ID,
		SubPlaceID:    updatedCycle.SubPlaceID,
		CommodityType: updatedCycle.CommodityType,
		CommodityName: updatedCycle.CommodityName,
		StartDate:     updatedCycle.StartDate,
		EndDate:       updatedCycle.EndDate,
		Status:        updatedCycle.Status,
		CreatedAt:     updatedCycle.CreatedAt.String(),
		UpdatedAt:     updatedCycle.UpdatedAt.String(),
	}, nil
}

// DeleteCycle menghapus cycle
func (u *cycleUsecase) DeleteCycle(ctx context.Context, id int) error {
	_, err := u.cycleRepo.DeleteCycle(ctx, id)
	return err
}

// GetAllCycles mengambil semua cycle dengan pagination
func (u *cycleUsecase) GetAllCycles(ctx context.Context, page, limit int) ([]dto.CycleResponse, dto.PaginationResponse, error) {
	limit, offset := utils.CalculatePagination(page, limit)

	cycles, total, err := u.cycleRepo.GetAllCycles(ctx, limit, offset)
	if err != nil {
		return nil, dto.PaginationResponse{}, err
	}

	var result []dto.CycleResponse

	for _, cycle := range cycles {
		result = append(result, dto.CycleResponse{
			ID:            cycle.ID,
			SubPlaceID:    cycle.SubPlaceID,
			CommodityType: cycle.CommodityType,
			CommodityName: cycle.CommodityName,
			StartDate:     cycle.StartDate,
			EndDate:       cycle.EndDate,
			Status:        cycle.Status,
			CreatedAt:     cycle.CreatedAt.String(),
			UpdatedAt:     cycle.UpdatedAt.String(),
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
