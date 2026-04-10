package usecase

import (
	"context"
	"farming/internal/dto"
	"farming/internal/model"
	"farming/internal/repository"
	"farming/pkg/utils"
)

type FarmUsecase interface {
	CreateFarm(ctx context.Context, req dto.CreateFarmRequest) (dto.FarmResponse, error)
	GetFarmById(ctx context.Context, id int) (dto.FarmResponse, error)
	UpdateFarm(ctx context.Context, id int, req dto.UpdateFarmRequest) (dto.FarmResponse, error)
	DeleteFarm(ctx context.Context, id int) error
	GetAllFarms(ctx context.Context) ([]dto.FarmResponse, error)
}

type farmUsecase struct {
	farmRepo repository.FarmRepository
	userRepo repository.UserRepository
}

func NewFarmUsecase(farmRepo repository.FarmRepository, userRepo repository.UserRepository) FarmUsecase {
	return &farmUsecase{
		farmRepo: farmRepo,
		userRepo: userRepo,
	}
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

	// ambil ulang dengan join
	fullFarm, err := u.farmRepo.GetFarmById(ctx, createdFarm.ID)
	if err != nil {
		return dto.FarmResponse{}, err
	}

	return utils.ToFarmResponse(fullFarm), nil

}

func (u *farmUsecase) GetFarmById(ctx context.Context, id int) (dto.FarmResponse, error) {
	farm, err := u.farmRepo.GetFarmById(ctx, id)

	if err != nil {
		return dto.FarmResponse{}, err
	}

	return utils.ToFarmResponse(farm), nil
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

	fullFarm, err := u.farmRepo.GetFarmById(ctx, updatedFarm.ID)
	if err != nil {
		return dto.FarmResponse{}, err
	}

	return utils.ToFarmResponse(fullFarm), nil
}

func (u *farmUsecase) DeleteFarm(ctx context.Context, id int) error {
	_, err := u.farmRepo.DeleteFarm(ctx, id)
	return err
}

func (u *farmUsecase) GetAllFarms(ctx context.Context) ([]dto.FarmResponse, error) {

	farms, err := u.farmRepo.GetAllFarms(ctx)
	if err != nil {
		return nil, err
	}

	var result []dto.FarmResponse

	for _, farm := range farms {
		result = append(result, utils.ToFarmResponse(farm))
	}

	return result, nil
}
