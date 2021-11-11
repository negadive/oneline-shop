package service

import (
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/schema"
	"gorm.io/gorm"
)

type UserService struct {
	DBCon *gorm.DB
}

func (c *UserService) Register(_user *schema.UserRegisterReq) (*model.User, error) {
	user := model.User{Email: _user.Email, Name: _user.Name, Password: _user.Password}
	result := c.DBCon.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
