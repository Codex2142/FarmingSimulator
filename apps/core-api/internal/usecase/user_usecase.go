package usecase

import (
	"context"
	"farming/internal/dto"
	"farming/internal/model"
	"farming/internal/repository"
	"farming/pkg/utils"
)

// ====================================================================
// Interface UserUsecase
// Mendefinisikan kontrak logika bisnis untuk user
type UserUsecase interface {

	// Membuat user baru
	CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.UserResponse, error)

	// Mengambil user berdasarkan ID
	GetUserById(ctx context.Context, id int) (dto.UserResponse, error)

	// Update Users
	UpdateUser(ctx context.Context, id int, req dto.UpdateUserRequest) (dto.UserResponse, error)

	// delete User
	DeleteUser(ctx context.Context, id int) error

	GetAllUsers(ctx context.Context, page, limit int) ([]dto.UserResponse, dto.PaginationResponse, error)
}

// ====================================================================
// Struct userUsecase
// Implementasi UserUsecase yang bergantung pada repository
type userUsecase struct {
	userRepo repository.UserRepository // dependency repository user
}

// ====================================================================
// Constructor NewUserUsecase
func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

// ====================================================================
// Implementasi CreateUser
func (u *userUsecase) CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.UserResponse, error) {

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return dto.UserResponse{}, err
	}
	// Membuat model.User dari request DTO
	user := model.User{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: hashedPassword,
	}

	// Memanggil repository untuk menyimpan user ke database
	createdUser, err := u.userRepo.CreateUser(ctx, user)
	if err != nil {

		// Jika gagal, kembalikan error
		return dto.UserResponse{}, err
	}

	// Mengubah model.User ke DTO UserResponse untuk dikirim ke handler
	return dto.UserResponse{
		ID:    createdUser.ID,
		Name:  createdUser.Name,
		Phone: createdUser.Phone,
	}, nil
}

// ====================================================================
// Implementasi GetUser
func (u *userUsecase) GetUserById(ctx context.Context, id int) (dto.UserResponse, error) {

	// Memanggil repository untuk mengambil user dari database
	user, err := u.userRepo.GetUserById(ctx, id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// Mengubah model.User ke DTO UserResponse
	return dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
	}, nil
}

// ====================================================================
// Memperbarui User
func (u *userUsecase) UpdateUser(ctx context.Context, id int, req dto.UpdateUserRequest) (dto.UserResponse, error) {

	// Membuat struct berdasarkan nilai dari request
	user := model.User{
		Name:  req.Name,
		Phone: req.Phone,
	}

	// Memanggil repository untuk mengambil user dari database
	updatedUser, err := u.userRepo.UpdateUser(ctx, user, id)

	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		ID:    updatedUser.ID,
		Name:  updatedUser.Name,
		Phone: updatedUser.Phone,
	}, nil
}

// ====================================================================
// Menghapus User
func (u *userUsecase) DeleteUser(ctx context.Context, id int) error {
	_, err := u.userRepo.DeleteUser(ctx, id)
	return err
}

func (u *userUsecase) GetAllUsers(ctx context.Context, page, limit int) ([]dto.UserResponse, dto.PaginationResponse, error) {

	// Helper calculate pagination
	limit, offset := utils.CalculatePagination(page, limit)

	users, total, err := u.userRepo.GetAllUsers(ctx, limit, offset)
	if err != nil {
		return nil, dto.PaginationResponse{}, err
	}

	var result []dto.UserResponse
	for _, user := range users {
		result = append(result, dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Phone: user.Phone,
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
