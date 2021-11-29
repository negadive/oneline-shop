package authorizer

import (
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/repository"
)

type IUserAuthorizer interface {
	CanEdit(user *model.User, user_id *uint) bool
}

type UserAuthorizer struct {
	userRepo repository.IUserRepository
}

func NewUserAuthorizer(userRepo repository.IUserRepository) IUserAuthorizer {
	return &UserAuthorizer{userRepo: userRepo}
}

func (ua *UserAuthorizer) CanEdit(user *model.User, user_id *uint) bool {
	return user.ID == *user_id
}
