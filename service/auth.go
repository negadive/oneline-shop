package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/schema"
	"gorm.io/gorm"
)

type AuthService struct {
	DBCon *gorm.DB
}

func (s *AuthService) CreateToken(user map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user["name"]
	claims["id"] = user["id"]
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString(([]byte("SECRET")))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *AuthService) Login(data *schema.LoginReq) (string, error) {
	user := map[string]interface{}{}
	result := s.DBCon.Model(&model.User{}).Where(&model.User{Email: data.Email, Password: data.Password}).First(&user)
	if result.Error != nil {
		return "", result.Error
	}

	token, err := s.CreateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
