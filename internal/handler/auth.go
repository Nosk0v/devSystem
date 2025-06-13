package handler

import (
	"devSystem/internal/usecase"
	"devSystem/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
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

// GetDepartments godoc
// @Summary Получить список департаментов
// @Description Возвращает все доступные департаменты
// @Tags auth
// @Produce json
// @Success 200 {array} models.Department
// @Failure 500 {object} ErrorResponse
// @Router /auth/departments [get]
func (h *Handler) getDepartments(c *gin.Context) {
	departments, err := h.usecases.GetDepartments()
	if err != nil {
		logrus.WithError(err).Error("Ошибка при получении департаментов")
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}
	h.sendResponseSuccess(c, departments, usecase.Success)
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

	code, err := h.usecases.CreateRegistrationCode(input.OrganizationID, input.IsAdmin, input.DepartmentID)
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

// GetOrganizationUsers godoc
// @Summary Получить пользователей организации
// @Description Возвращает список всех пользователей в организации (по JWT или query).
// @Tags users
// @Accept json
// @Produce json
// @Param organization_id query int false "ID организации (только для SuperAdmin)"
// @Success 200 {array} models.UserResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/organization [get]
func (h *Handler) getOrganizationUsers(c *gin.Context) {
	log.Println("Handling getOrganizationUsers...")

	claims, err := h.GetJWTClaims(c)
	if err != nil {
		log.Printf("Failed to get JWT claims: %v\n", err)
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Unauthorized"})
		return
	}
	log.Printf("Extracted claims: %+v\n", claims)

	var orgID int
	if claims.Role == usecase.RoleSuperAdmin {
		orgIDStr := c.Query("organization_id")
		log.Printf("SuperAdmin request with organization_id: %s\n", orgIDStr)

		if orgIDStr == "" {
			log.Println("Missing organization_id in query params")
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "organization_id is required for SuperAdmin"})
			return
		}

		orgID, err = strconv.Atoi(orgIDStr)
		if err != nil || orgID <= 0 {
			log.Printf("Invalid organization_id format: %s, err: %v\n", orgIDStr, err)
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid organization_id"})
			return
		}
	} else {
		if claims.OrganizationID == nil {
			log.Println("Organization ID missing from token for non-superadmin")
			c.JSON(http.StatusForbidden, ErrorResponse{Error: "Organization ID missing from token"})
			return
		}
		orgID = *claims.OrganizationID
		log.Printf("Non-superadmin, using orgID from token: %d\n", orgID)
	}

	log.Printf("Fetching users for orgID: %d\n", orgID)
	users, err := h.usecases.GetUsersByOrganization(orgID)
	if err != nil {
		log.Printf("Failed to get users from usecase for orgID=%d: %v\n", orgID, err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to fetch users"})
		return
	}

	log.Printf("Successfully fetched %d users for orgID=%d\n", len(users), orgID)
	c.JSON(http.StatusOK, users)
}

// DeleteUser godoc
// @Summary Удаление пользователя
// @Description Удаляет пользователя по email. Только для SuperAdmin.
// @Tags users
// @Param email path string true "Email пользователя"
// @Success 200 {object} map[string]string "User deleted successfully"
// @Failure 400 {object} ErrorResponse "Email is required"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 403 {object} ErrorResponse "Only super admin can delete users"
// @Failure 500 {object} ErrorResponse "Failed to delete user"
// @Router /auth/users/{email} [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Email is required"})
		return
	}

	claims, err := h.GetJWTClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Unauthorized"})
		return
	}

	if claims.Role != usecase.RoleSuperAdmin {
		c.JSON(http.StatusForbidden, ErrorResponse{Error: "Only super admin can delete users"})
		return
	}

	if err := h.usecases.DeleteUser(email, claims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
