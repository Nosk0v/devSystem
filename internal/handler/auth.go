package handler

import (
	"devSystem/internal/usecase"
	"devSystem/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
