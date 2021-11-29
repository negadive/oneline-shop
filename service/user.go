package service

import (
	"context"

	"github.com/negadive/oneline/authorizer"
	"github.com/negadive/oneline/custom_errors"
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/repository"
	"gorm.io/gorm"
)

type IUserService interface {
	Register(ctx context.Context, user *model.User) error
	Update(ctx context.Context, actor_id *uint, user_id *uint, user *model.User) error
}

type UserService struct {
	DBCon       *gorm.DB
	UserAuthzer authorizer.IUserAuthorizer
	UserRepo    repository.IUserRepository
}

func NewUserService(db_con *gorm.DB, userAuthzer authorizer.IUserAuthorizer, user_repo repository.IUserRepository) IUserService {
	return &UserService{
		DBCon:       db_con,
		UserAuthzer: userAuthzer,
		UserRepo:    user_repo,
	}
}

func (s *UserService) Register(ctx context.Context, user *model.User) error {
	tx := s.DBCon.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	if err := s.UserRepo.Store(ctx, tx, user); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (s *UserService) Update(ctx context.Context, actor_id *uint, user_id *uint, user *model.User) error {
	tx := s.DBCon.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	old_user, err := s.UserRepo.FindById(ctx, tx, user_id)
	if err != nil {
		tx.Rollback()
		return err
	}
	if !s.UserAuthzer.CanEdit(old_user, actor_id) {
		tx.Rollback()
		return custom_errors.NewForbiddenUser("not user data")
	}
	if err := s.UserRepo.Update(ctx, tx, user_id, user); err != nil {
		tx.Rollback()
		return err
	}
	new_user, err := s.UserRepo.FindById(ctx, tx, user_id)
	if err != nil {
		tx.Rollback()
		return err
	}
	*user = *new_user

	return nil
}
