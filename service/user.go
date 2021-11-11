package service

import (
	"github.com/negadive/oneline/db"
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/schema"
)

func Register(_user *schema.UserRegisterReq) (*model.User, error) {
	_db := db.GetDb()

	user := model.User{Email: _user.Email, Name: _user.Name, Password: _user.Password}
	result := _db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
