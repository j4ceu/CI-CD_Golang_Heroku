package services

import (
	"CICD_GolangtoHeroku/dto"
	"CICD_GolangtoHeroku/middlewares"
	"CICD_GolangtoHeroku/models"
	"CICD_GolangtoHeroku/repositories"
)

type UserService interface {
	CreateUser(user models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	LoginUser(email string, password string) (dto.UserResponse, error)
}

type userService struct {
	userRepo repositories.UserRepo
}

func NewUserService(userRepo repositories.UserRepo) *userService {
	return &userService{userRepo}
}

func (u *userService) CreateUser(user models.User) (models.User, error) {

	user, err := u.userRepo.CreateUser(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userService) GetAllUsers() ([]models.User, error) {
	users, err := u.userRepo.GetAllUsers()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (u *userService) LoginUser(email string, password string) (dto.UserResponse, error) {
	user, err := u.userRepo.FindByEmailAndPassword(email, password)
	if err != nil {
		return dto.UserResponse{}, err
	}

	token, err := middlewares.CreateToken(user.ID, user.Email)
	if err != nil {
		return dto.UserResponse{}, err
	}

	userResponse := dto.UserResponse{
		Email: user.Email,
		Token: token,
	}

	return userResponse, nil
}
