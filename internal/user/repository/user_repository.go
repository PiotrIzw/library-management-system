package repository

import (
	"gorm.io/gorm"
	"library-management-system/internal/user/model"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
