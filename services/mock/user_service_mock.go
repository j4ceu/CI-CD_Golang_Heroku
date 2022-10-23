package mock

import (
	"CICD_GolangtoHeroku/dto"
	"CICD_GolangtoHeroku/models"

	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (u *UserServiceMock) CreateUser(user models.User) (models.User, error) {
	args := u.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (u *UserServiceMock) GetAllUsers() ([]models.User, error) {
	args := u.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (u *UserServiceMock) LoginUser(email string, password string) (dto.UserResponse, error) {
	args := u.Called(email, password)
	return args.Get(0).(dto.UserResponse), args.Error(1)
}
