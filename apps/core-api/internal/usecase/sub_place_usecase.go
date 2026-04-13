package usecase

import (
	"context"
	"farming/internal/dto"
	"farming/internal/model"
	"farming/internal/repository"
	"farming/pkg/constants"
	"farming/pkg/utils"
)

// SubPlaceUsecase mendefinisikan kontrak logika bisnis untuk sub_place
type SubPlaceUsecase interface {
	CreateSubPlace(ctx context.Context, req dto.CreateSubPlaceRequest) (dto.SubPlaceResponse, error)
	GetSubPlaceById(ctx context.Context, id int) (dto.SubPlaceResponse, error)
	UpdateSubPlace(ctx context.Context, id int, req dto.UpdateSubPlaceRequest) (dto.SubPlaceResponse, error)
	DeleteSubPlace(ctx context.Context, id int) error
	GetAllSubPlaces(ctx context.Context, page, limit int) ([]dto.SubPlaceResponse, dto.PaginationResponse, error)
}

// Implementasi SubPlaceUsecase
type subPlaceUsecase struct {
	subPlaceRepo repository.SubPlaceRepository
	farmRepo     repository.FarmRepository
}

// Constructor NewSubPlaceUsecase
func NewSubPlaceUsecase(subPlaceRepo repository.SubPlaceRepository, farmRepo repository.FarmRepository) SubPlaceUsecase {
	return &subPlaceUsecase{
		subPlaceRepo: subPlaceRepo,
		farmRepo:     farmRepo,
	}
}

// CreateSubPlace membuat sub_place baru
func (u *subPlaceUsecase) CreateSubPlace(ctx context.Context, req dto.CreateSubPlaceRequest) (dto.SubPlaceResponse, error) {
	// Validasi farm ada
	_, err := u.farmRepo.GetFarmById(ctx, req.FarmID)
	if err != nil {
		return dto.SubPlaceResponse{}, err
	}

	// Set default status jika kosong
	if req.Status == "" {
		req.Status = string(constants.Active)
	}

	// Buat model
	subPlace := model.SubPlace{
		FarmID: req.FarmID,
		Name:   req.Name,
		Type:   req.Type,
		Size:   req.Size,
		Status: req.Status,
	}

	// Save ke database
	createdSubPlace, err := u.subPlaceRepo.CreateSubPlace(ctx, subPlace)
	if err != nil {
		return dto.SubPlaceResponse{}, err
	}

	// Convert ke response
	return dto.SubPlaceResponse{
		ID:        createdSubPlace.ID,
		FarmID:    createdSubPlace.FarmID,
		Name:      createdSubPlace.Name,
		Type:      createdSubPlace.Type,
		Size:      createdSubPlace.Size,
		Status:    createdSubPlace.Status,
		CreatedAt: createdSubPlace.CreatedAt.String(),
		UpdatedAt: createdSubPlace.UpdatedAt.String(),
	}, nil
}

// GetSubPlaceById mengambil sub_place berdasarkan ID
func (u *subPlaceUsecase) GetSubPlaceById(ctx context.Context, id int) (dto.SubPlaceResponse, error) {
	subPlace, err := u.subPlaceRepo.GetSubPlaceById(ctx, id)
	if err != nil {
		return dto.SubPlaceResponse{}, err
	}

	return dto.SubPlaceResponse{
		ID:        subPlace.ID,
		FarmID:    subPlace.FarmID,
		Name:      subPlace.Name,
		Type:      subPlace.Type,
		Size:      subPlace.Size,
		Status:    subPlace.Status,
		CreatedAt: subPlace.CreatedAt.String(),
		UpdatedAt: subPlace.UpdatedAt.String(),
	}, nil
}

// UpdateSubPlace mengupdate sub_place
func (u *subPlaceUsecase) UpdateSubPlace(ctx context.Context, id int, req dto.UpdateSubPlaceRequest) (dto.SubPlaceResponse, error) {
	subPlace := model.SubPlace{
		Name:   req.Name,
		Type:   req.Type,
		Size:   req.Size,
		Status: req.Status,
	}

	updatedSubPlace, err := u.subPlaceRepo.UpdateSubPlace(ctx, subPlace, id)
	if err != nil {
		return dto.SubPlaceResponse{}, err
	}

	return dto.SubPlaceResponse{
		ID:        updatedSubPlace.ID,
		FarmID:    updatedSubPlace.FarmID,
		Name:      updatedSubPlace.Name,
		Type:      updatedSubPlace.Type,
		Size:      updatedSubPlace.Size,
		Status:    updatedSubPlace.Status,
		CreatedAt: updatedSubPlace.CreatedAt.String(),
		UpdatedAt: updatedSubPlace.UpdatedAt.String(),
	}, nil
}

// DeleteSubPlace menghapus sub_place
func (u *subPlaceUsecase) DeleteSubPlace(ctx context.Context, id int) error {
	_, err := u.subPlaceRepo.DeleteSubPlace(ctx, id)
	return err
}

// GetAllSubPlaces mengambil semua sub_place dengan pagination
func (u *subPlaceUsecase) GetAllSubPlaces(ctx context.Context, page, limit int) ([]dto.SubPlaceResponse, dto.PaginationResponse, error) {
	limit, offset := utils.CalculatePagination(page, limit)

	subPlaces, total, err := u.subPlaceRepo.GetAllSubPlaces(ctx, limit, offset)
	if err != nil {
		return nil, dto.PaginationResponse{}, err
	}

	var result []dto.SubPlaceResponse

	for _, subPlace := range subPlaces {
		result = append(result, dto.SubPlaceResponse{
			ID:        subPlace.ID,
			FarmID:    subPlace.FarmID,
			Name:      subPlace.Name,
			Type:      subPlace.Type,
			Size:      subPlace.Size,
			Status:    subPlace.Status,
			CreatedAt: subPlace.CreatedAt.String(),
			UpdatedAt: subPlace.UpdatedAt.String(),
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
