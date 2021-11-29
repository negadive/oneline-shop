package service

import (
	"context"

	"github.com/negadive/oneline/authorizer"
	"github.com/negadive/oneline/customErrors"
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/repository"
	"gorm.io/gorm"
)

type IUserService interface {
	Register(ctx context.Context, user *model.User) error
	Update(ctx context.Context, actorId *uint, user_id *uint, user *model.User) error
}

type UserService struct {
	dBCon       *gorm.DB
	userAuthzer authorizer.IUserAuthorizer
	userRepo    repository.IUserRepository
}

func NewUserService(
	dbCon *gorm.DB,
	userAuthzer authorizer.IUserAuthorizer,
	userRepo repository.IUserRepository,
) IUserService {
	return &UserService{
		dBCon:       dbCon,
		userAuthzer: userAuthzer,
		userRepo:    userRepo,
	}
}

func (s *UserService) Register(ctx context.Context, user *model.User) error {
	tx := s.dBCon.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	if err := s.userRepo.Store(ctx, tx, user); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (s *UserService) Update(ctx context.Context, actorId *uint, user_id *uint, user *model.User) error {
	tx := s.dBCon.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	oldUser, err := s.userRepo.FindById(ctx, tx, user_id)
	if err != nil {
		tx.Rollback()
		return err
	}
	if !s.userAuthzer.CanEdit(oldUser, actorId) {
		tx.Rollback()
		return customErrors.NewForbiddenUser("not user data")
	}
	if err := s.userRepo.Update(ctx, tx, user_id, user); err != nil {
		tx.Rollback()
		return err
	}
	new_user, err := s.userRepo.FindById(ctx, tx, user_id)
	if err != nil {
		tx.Rollback()
		return err
	}
	*user = *new_user

	return nil
}
