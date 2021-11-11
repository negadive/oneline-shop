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

func (c *UserService) UpdateUser(_user *schema.UserUpdateReq, user_id int) (*model.User, error) {
	var count int64
	if c.DBCon.Model(&model.User{}).Where("id = ?", user_id).Count(&count); count < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	user := model.User{}
	if err := c.DBCon.Model(&model.User{}).Where("id = ?", user_id).Updates(model.User{Name: _user.Name, Password: _user.Password}).Error; err != nil {
		return nil, err
	}
	if err := c.DBCon.Where("id = ?", user_id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
