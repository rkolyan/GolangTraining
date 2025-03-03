package service

import (
	Domains "Backend/internal/business/domains"
	"github.com/golang-jwt/jwt"
	"time"
)

type AuthService struct {
	repository Domains.IAuthRepository
}

func (service *AuthService) Login(data *Domains.LoginData) (*Domains.AuthTokens, error) {
	usr, err := service.repository.GetAuthData(data)
	//TODO: err
	//TODO:
	token1 := jwt.NewWithClaims(&jwt.SigningMethodHMAC{}, jwt.MapClaims{"id": usr.ID, "exp": time.Now().Add(time.Hour * 1).Unix()})
	token2 := jwt.NewWithClaims(&jwt.SigningMethodHMAC{}, jwt.MapClaims{"id": usr.ID, "exp": time.Now().Add(time.Hour * 10).Unix()})
	//TODO: Сохранить случайную соль в Redis?
	token1String, err := token1.SignedString([]byte("your_secret_key_here"))
	//TODO: ERR
	token2String, err := token2.SignedString([]byte("your_secret_key_here"))
	//TODO: err
	//TODO: Используй defer для отложенного удаления ресурсов
	return &Domains.AuthTokens{AccessToken: token1String, RefreshToken: token2String}, err
}

func ()  {
	
}
