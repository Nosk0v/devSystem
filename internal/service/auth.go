package service

import (
	"devSystem/internal/repository"
	"devSystem/internal/utils"
	"devSystem/models"
	"fmt"
	"github.com/sirupsen/logrus"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (a AuthService) SignIn(input *models.SignInInput, accountPassword string) error {
	if err := utils.ComparePasswords(accountPassword, input.Password); err != nil {
		logrus.Warning(err.Error())
		return err
	}
	return nil
}

func (a *AuthService) SignUp(input *models.SignUpInput) error {
	exists, err := a.repo.AccountExists(input.Email)
	if err != nil {
		logrus.Errorf("Ошибка при проверке существования email: %v", err)
		return err
	}
	if exists {
		return fmt.Errorf("пользователь с таким email уже зарегистрирован")
	}

	hashedPassword, err := utils.GetPasswordHash(input.Password)
	if err != nil {
		logrus.Errorf("Ошибка хеширования пароля: %v", err)
		return err
	}
	input.Password = string(hashedPassword)

	err = a.repo.CreateAccount(input)
	if err != nil {
		logrus.Errorf("Ошибка создания аккаунта: %v", err)
		return err
	}

	return nil
}
