package repositories

import (
	"CICD_GolangtoHeroku/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	FindByEmailAndPassword(email string, password string) (models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db}
}

func (u *userRepo) CreateUser(user models.User) (models.User, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := u.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (u *userRepo) FindByEmailAndPassword(email string, password string) (models.User, error) {
	var user models.User
	err := u.db.Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
