package usecase

import (
	"devSystem/models"
	"github.com/sirupsen/logrus"
)

func (u *Usecase) SignIn(input *models.SignInInput) (*models.SignInOutput, ErrorCode) {
	account, err := u.services.Account.Get(input.Email)
	if err != nil {
		logrus.Error(err)
		return nil, Unauthorized
	}
	err = u.services.Auth.SignIn(input, account.Password)
	if err != nil {
		logrus.Error(err.Error())
		return nil, Unauthorized
	}

	accessToken, err := u.services.JWTToken.GenerateAccessToken(account.Email)
	if err != nil {
		logrus.Error("ошибка генерации Access токена: ", err)
		return nil, InternalServerError
	}

	refreshToken, err := u.services.JWTToken.GenerateRefreshToken(account.Email)
	if err != nil {
		logrus.Error("ошибка генерации Refresh токена: ", err)
		return nil, InternalServerError
	}

	output := &models.SignInOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return output, Success
}

func (u *Usecase) ParseToken(token string) (*models.JWTClaims, ErrorCode) {
	claims, err := u.services.JWTToken.ParseToken(token)
	if err != nil {
		logrus.Error(err)
		return nil, InternalServerError
	}

	return claims, Success
}
