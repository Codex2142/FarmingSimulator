package usecase

import (
	"context"
	"farming/internal/dto"
	"farming/internal/model"
	"farming/internal/repository"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.UserResponse, error)
	GetUser(ctx context.Context, id int) (dto.UserResponse, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.UserResponse, error) {
	user := model.User{
		Name:  req.Name,
		Phone: req.Phone,
	}
	createdUser, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return dto.UserResponse{
		ID:    createdUser.ID,
		Name:  createdUser.Name,
		Phone: createdUser.Phone,
	}, nil
}

func (u *userUsecase) GetUser(ctx context.Context, id int) (dto.UserResponse, error) {
	user, err := u.userRepo.GetbyID(ctx, id)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
	}, nil
}
