package handler

import (
	"devSystem/internal/usecase"
	"devSystem/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

// SignIn godoc
// @Summary Вход в систему
// @Description Авторизация пользователя по email и паролю.
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.SignInInput true "Данные для входа"
// @Success 200
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input models.SignInInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.Error(err.Error())
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	output, processStatus := h.usecases.SignIn(&input)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	h.sendResponseSuccess(c, output, processStatus)
}

// DeleteOrganization godoc
// @Summary Удаление организации
// @Description Удаляет организацию и все связанные с ней записи
// @Tags auth
// @Param id path int true "ID организации"
// @Success 200 {object} gin.H
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/organization/{id} [delete]
func (h *Handler) deleteOrganization(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	err = h.usecases.DeleteOrganization(id)
	if err != nil {
		logrus.WithError(err).Error("Ошибка при удалении организации")
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	h.sendResponseSuccess(c, gin.H{"message": "Организация удалена"}, usecase.Success)
}

// Refresh godoc
// @Summary Обновить access токен
// @Description Обновление access токена с использованием refresh токена. Refresh токен остаётся прежним.
// @Tags auth
// @Accept json
// @Produce json
// @Success 200
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/refresh [post]
func (h *Handler) refresh(c *gin.Context) {
	var request struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, ErrorResponse{Error: "Invalid JSON body"})
		return
	}

	logrus.Infof("Received refresh token: %v", request.RefreshToken)

	if request.RefreshToken == "" {
		c.JSON(401, ErrorResponse{Error: "Refresh token is missing"})
		return
	}

	output, processStatus := h.usecases.Refresh(request.RefreshToken)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	h.sendResponseSuccess(c, output, processStatus)
}

// SignUp godoc
// @Summary Регистрация пользователя
// @Description Создание нового аккаунта
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.SignUpInput true "Данные для регистрации"
// @Success 201 {object} models.SignUpOutput
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse "Пользователь уже существует"
// @Failure 500 {object} ErrorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input models.SignUpInput

	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.WithError(err).Warn("Некорректный JSON при регистрации")
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	processStatus := h.usecases.SignUp(&input)

	switch processStatus {
	case usecase.ResourceCreated:
		h.sendResponseSuccess(c, models.SignUpOutput{Message: "Регистрация прошла успешно"}, processStatus)
	case usecase.ResourceAlreadyExist:
		h.sendResponseSuccess(c, nil, processStatus)
	case usecase.CodeNotFound:
		h.sendResponseSuccess(c, nil, processStatus)
	default:
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
	}
}

// GetOrganizations godoc
// @Summary Получить список организаций
// @Description Возвращает все доступные организации
// @Tags auth
// @Produce json
// @Success 200 {array} models.Organization
// @Failure 500 {object} ErrorResponse
// @Router /auth/organizations [get]
func (h *Handler) getOrganizations(c *gin.Context) {
	orgs, err := h.usecases.GetOrganizations()
	if err != nil {
		logrus.WithError(err).Error("Ошибка при получении организаций")
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}
	h.sendResponseSuccess(c, orgs, usecase.Success)
}

// CreateOrganization godoc
// @Summary Создание организации
// @Description Создаёт новую организацию с уникальным префиксом
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.CreateOrganizationInput true "Данные организации"
// @Success 201 {object} models.CreateOrganizationOutput
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/organization [post]
func (h *Handler) createOrganization(c *gin.Context) {
	var input models.CreateOrganizationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.WithError(err).Warn("Некорректный JSON при создании организации")
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	status := h.usecases.CreateOrganizationWithPrefix(input)
	h.sendResponseSuccess(c, status, usecase.Success)
}

// CreateRegistrationCode godoc
// @Summary Создание регистрационного кода
// @Description Генерация одноразового кода для регистрации
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.CreateRegistrationCodeInput true "Параметры создания кода"
// @Success 200 {object} models.RegistrationCodeOutput
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/registration-code [post]
func (h *Handler) createRegistrationCode(c *gin.Context) {
	var input models.CreateRegistrationCodeInput

	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.WithError(err).Warn("Неверный JSON при создании регистрационного кода")
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	code, err := h.usecases.CreateRegistrationCode(input.OrganizationID, input.IsAdmin)
	if err != nil {
		logrus.WithError(err).Error("Ошибка при создании регистрационного кода")
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	h.sendResponseSuccess(c, models.RegistrationCodeOutput{Code: code}, usecase.Success)
}

// DeleteRegistrationCode godoc
// @Summary Удалить регистрационный код
// @Description Удаление регистрационного кода по строковому ID
// @Tags auth
// @Param code path string true "Регистрационный код"
// @Success 200
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/registration-code/{code} [delete]
func (h *Handler) deleteRegistrationCode(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	err := h.usecases.DeleteRegistrationCode(code)
	if err != nil {
		logrus.WithError(err).Error("Ошибка при удалении регистрационного кода")
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	h.sendResponseSuccess(c, gin.H{"message": "Код удалён"}, usecase.Success)
}
