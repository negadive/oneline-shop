package repository

import (
	"context"

	"github.com/negadive/oneline/customErrors"
	"github.com/negadive/oneline/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Store(ctx context.Context, tx *gorm.DB, user *model.User) error
	Update(ctx context.Context, tx *gorm.DB, userId *uint, user *model.User) error
	FindById(ctx context.Context, tx *gorm.DB, userId *uint) (*model.User, error)
	IsExists(ctx context.Context, tx *gorm.DB, userId *uint) (bool, error)
}

type UserRepository struct {
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) Store(ctx context.Context, tx *gorm.DB, user *model.User) error {
	result := tx.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *UserRepository) Update(ctx context.Context, tx *gorm.DB, userId *uint, user *model.User) error {
	if err := tx.WithContext(ctx).Model(&model.User{}).Where("id = ?", *userId).Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) IsExists(ctx context.Context, tx *gorm.DB, userId *uint) (bool, error) {
	var count int64
	if err := tx.WithContext(ctx).Model(&model.User{}).Where("id = ?", userId).Count(&count).Error; err != nil {
		return false, customErrors.NewNotFoundError("user")
	}

	return count > 0, nil
}

func (repo *UserRepository) FindById(ctx context.Context, tx *gorm.DB, userId *uint) (*model.User, error) {
	user := model.User{}
	result := tx.WithContext(ctx).First(&user, userId)
	if result.Error != nil {
		return nil, customErrors.NewNotFoundError("user")
	}

	return &user, nil
}
