package usecase

import (
	"devSystem/models"
	"github.com/sirupsen/logrus"
	"strings"
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

	accessToken, err := u.services.JWTToken.GenerateAccessToken(account.Email, account.Role, account.OrganizationID)
	if err != nil {
		logrus.Error("ошибка генерации Access токена: ", err)
		return nil, InternalServerError
	}

	refreshToken, err := u.services.JWTToken.GenerateRefreshToken(account.Email, account.Role, account.OrganizationID)
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
func (u *Usecase) SignUp(input *models.SignUpInput) ErrorCode {
	input.Role = 1
	err := u.services.Auth.SignUp(input)
	if err != nil {
		if strings.Contains(err.Error(), "уже зарегистрирован") {
			return ResourceAlreadyExist
		}
		logrus.Error("Ошибка в сервисе при создании аккаунта: ", err)
		return InternalServerError
	}
	return ResourceCreated
}

func (u *Usecase) Refresh(refreshToken string) (*models.AccessTokenOutput, ErrorCode) {
	accessToken, err := u.services.JWTToken.GenerateAccessFromRefresh(refreshToken)
	if err != nil {
		logrus.Error("Ошибка при попытке обновить access token: ", err)
		return nil, Unauthorized
	}
	output := &models.AccessTokenOutput{
		AccessToken: accessToken,
	}

	return output, Success
}
