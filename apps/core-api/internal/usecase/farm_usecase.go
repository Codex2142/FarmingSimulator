package usecase

import (
	"context"
	"farming/internal/dto"
	"farming/internal/model"
	"farming/internal/repository"
)

type FarmUsecase interface {
	CreateFarm(ctx context.Context, req dto.CreateFarmRequest) (dto.FarmResponse, error)
	GetFarmById(ctx context.Context, id int) (dto.FarmResponse, error)
	UpdateFarm(ctx context.Context, id int, req dto.UpdateFarmRequest) (dto.FarmResponse, error)
	DeleteFarm(ctx context.Context, id int) error
}

type farmUsecase struct {
	farmRepo repository.FarmRepository
	userRepo repository.UserRepository
}

func NewFarmUsecase(farmRepo repository.FarmRepository) FarmUsecase {
	return &farmUsecase{farmRepo: farmRepo}
}

func (u *farmUsecase) CreateFarm(ctx context.Context, req dto.CreateFarmRequest) (dto.FarmResponse, error) {

	// pengecekan relasi ke user
	_, err := u.userRepo.GetUserById(ctx, req.LeaderID)
	if err != nil {
		return dto.FarmResponse{}, err
	}

	// Init model untuk di insert ke DB
	farm := model.Farm{
		Name:     req.Name,
		Location: req.Location,
		LeaderID: &req.LeaderID,
	}

	createdFarm, err := u.farmRepo.CreateFarm(ctx, farm)
	if err != nil {
		return dto.FarmResponse{}, err
	}

	return dto.FarmResponse{
		ID:       createdFarm.ID,
		Name:     createdFarm.Name,
		Location: createdFarm.Location,
		LeaderID: *createdFarm.LeaderID,
	}, nil

}

func (u *farmUsecase) GetFarmById(ctx context.Context, id int) (dto.FarmResponse, error) {
	farm, err := u.farmRepo.GetFarmById(ctx, id)

	if err != nil {
		return dto.FarmResponse{}, err
	}

	return dto.FarmResponse{
		ID:       farm.ID,
		Name:     farm.Name,
		Location: farm.Location,
		LeaderID: *farm.LeaderID,
	}, nil
}

func (u *farmUsecase) UpdateFarm(ctx context.Context, id int, req dto.UpdateFarmRequest) (dto.FarmResponse, error) {

	// pengecekan relasi ke user
	_, err := u.userRepo.GetUserById(ctx, req.LeaderID)
	if err != nil {
		return dto.FarmResponse{}, err
	}

	farm := model.Farm{
		Name:     req.Name,
		Location: req.Location,
		LeaderID: &req.LeaderID,
	}

	updatedFarm, err := u.farmRepo.UpdateFarm(ctx, farm, id)

	if err != nil {
		return dto.FarmResponse{}, err
	}

	return dto.FarmResponse{
		ID:       updatedFarm.ID,
		Name:     updatedFarm.Name,
		Location: updatedFarm.Location,
		LeaderID: *updatedFarm.LeaderID,
	}, nil
}

func (u *farmUsecase) DeleteFarm(ctx context.Context, id int) error {
	_, err := u.farmRepo.DeleteFarm(ctx, id)
	return err
}
