package usecase

import (
	"devSystem/models"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
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

	accessToken, err := u.services.JWTToken.GenerateAccessToken(account.Email, account.Role, account.OrganizationID, account.DepartmentID)
	if err != nil {
		logrus.Error("ошибка генерации Access токена: ", err)
		return nil, InternalServerError
	}

	refreshToken, err := u.services.JWTToken.GenerateRefreshToken(account.Email, account.Role, account.OrganizationID, account.DepartmentID)
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
	// Получить информацию по коду
	codeInfo, err := u.services.Auth.GetRegistrationCodeInfo(input.Code)
	if err != nil {
		logrus.Error("не удалось получить информацию по коду: ", err)
		return CodeNotFound
	}

	if codeInfo.IsUsed {
		logrus.Warn("код уже использован")
		return CodeAlreadyUsed
	}

	if codeInfo.ExpiresAt != nil && codeInfo.ExpiresAt.Before(time.Now()) {
		logrus.Warn("срок действия кода истёк")
		return ResourceExpired
	}

	if codeInfo.Code == "" || codeInfo.OrganizationID == 0 {
		logrus.Warn("Регистрационный код невалиден или не содержит организацию")
		return CodeNotFound
	}

	// Заполняем роль и организацию на основе кода
	if codeInfo.Role == 0 {
		input.Role = 0 // админ
	} else {
		input.Role = 1 // обычный
	}
	input.OrganizationID = codeInfo.OrganizationID
	input.DepartmentID = codeInfo.DepartmentID

	// Регистрируем
	err = u.services.Auth.SignUp(input)
	if err != nil {
		if strings.Contains(err.Error(), "уже зарегистрирован") {
			return ResourceAlreadyExist
		}
		logrus.Error("Ошибка при создании аккаунта: ", err)
		return InternalServerError
	}

	// Помечаем код как использованный
	err = u.services.Auth.MarkRegistrationCodeAsUsed(input.Code)
	if err != nil {
		logrus.Warn("Не удалось пометить код как использованный: ", err)
		// Не критично, регистрацию не прерываем
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

func (u *Usecase) GetOrganizations() ([]models.Organization, error) {
	return u.services.Auth.GetOrganizations()
}

func (u *Usecase) GetDepartments() ([]models.Department, error) {
	return u.services.Auth.GetDepartments()
}

func (u *Usecase) DeleteRegistrationCode(code string) error {
	return u.services.Auth.DeleteRegistrationCode(code)
}

func (u *Usecase) GetRegistrationCodeInfo(code string) (*models.RegistrationCode, error) {
	return u.services.Auth.GetRegistrationCodeInfo(code)
}

func (u *Usecase) CreateRegistrationCode(orgID int, isAdmin bool, departmentID *int) (string, error) {
	prefix, err := u.services.Auth.GetPrefixByOrgID(orgID)
	if err != nil {
		logrus.WithError(err).Error("не удалось получить префикс по организации")
		return "", err
	}

	code, err := u.services.Auth.CreateRegistrationCode(prefix, isAdmin, departmentID)
	if err != nil {
		logrus.WithError(err).Error("не удалось создать код регистрации")
		return "", err
	}

	return code, nil
}

func (u *Usecase) CreateOrganizationWithPrefix(input models.CreateOrganizationInput) error {
	return u.services.Auth.CreateOrganizationWithPrefix(input.Name, input.Prefix)
}

func (u *Usecase) DeleteOrganization(organizationID int) error {
	return u.services.Auth.DeleteOrganization(organizationID)
}

func (u *Usecase) GetUsersByOrganization(orgID int) ([]models.UserResponse, error) {
	return u.services.Auth.GetUsersByOrganization(orgID)
}

func (u *Usecase) DeleteUser(email string, byRole int) error {
	if byRole != RoleSuperAdmin {
		return fmt.Errorf("only super admin can delete users")
	}
	return u.services.Auth.DeleteUser(email)
}
