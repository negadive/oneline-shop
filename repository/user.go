package repository

import (
	"context"

	"github.com/negadive/oneline/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Store(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user_id *uint, user *model.User) error
	FindById(ctx context.Context, user_id *uint) (*model.User, error)
	IsExists(ctx context.Context, user_id *uint) (bool, error)
}

type UserRepository struct {
	DBCon *gorm.DB
}

func NewUserRepository(DBCon *gorm.DB) IUserRepository {
	return &UserRepository{
		DBCon: DBCon,
	}
}

func (repo *UserRepository) Store(ctx context.Context, user *model.User) error {
	result := repo.DBCon.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *UserRepository) Update(ctx context.Context, user_id *uint, user *model.User) error {
	if err := repo.DBCon.WithContext(ctx).Model(&model.User{}).Where("id = ?", *user_id).Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) IsExists(ctx context.Context, user_id *uint) (bool, error) {
	var count int64
	if err := r.DBCon.WithContext(ctx).Model(&model.User{}).Where("id = ?", user_id).Count(&count).Error; err != nil {
		return false, gorm.ErrRecordNotFound
	}

	return count > 0, nil
}

func (repo *UserRepository) FindById(ctx context.Context, user_id *uint) (*model.User, error) {
	user := model.User{}
	result := repo.DBCon.WithContext(ctx).First(&user, user_id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
