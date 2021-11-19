package service

import (
	"context"

	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/repository"
	"gorm.io/gorm"
)

type IUserService interface {
	Register(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user_id *uint, user *model.User) error
}

type UserService struct {
	UserRepo repository.IUserRepository
}

func NewUserService(user_repo repository.IUserRepository) IUserService {
	return &UserService{
		UserRepo: user_repo,
	}
}

func (s *UserService) Register(ctx context.Context, user *model.User) error {
	if err := s.UserRepo.Store(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) Update(ctx context.Context, user_id *uint, user *model.User) error {
	is_exists, err := s.UserRepo.IsExists(ctx, user_id)
	if err != nil {
		return err
	}
	if !is_exists {
		return gorm.ErrRecordNotFound
	}
	if err := s.UserRepo.Update(ctx, user_id, user); err != nil {
		return err
	}
	new_user, err := s.UserRepo.FindById(ctx, user_id)
	if err != nil {
		return err
	}
	*user = *new_user

	return nil
}
