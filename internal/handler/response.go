package handler

import (
	"devSystem/clerr"
	"devSystem/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	accessDenied          = "accessDenied"
	tariffExpired         = "tariffExpired"
	generalException      = "generalException"
	itemNotFound          = "itemNotFound"
	invalidRequest        = "invalidRequest"
	resourceModified      = "resourceModified"
	unAuthenticated       = "unAuthenticated"
	sendEmail             = "sendEmail"
	sendEmailError        = "smtp server error"
	conditionNotMet       = "conditionNotMet"
	resourceDeleted       = "resourceDeleted"
	notConfirmed          = "notConfirmed"
	timeoutExpired        = "timeoutExpired"
	headerException       = "headerException"
	sendEmailErrorStatus  = 506
	headerExceptionStatus = 432
	resourceInTrash       = 434
)

func (h *Handler) sendResponseSuccess(c *gin.Context, successResponse any, err usecase.ErrorCode) {
	if successResponse == nil {
		if err == usecase.NoContent {
			c.Status(http.StatusNoContent)
			return
		}
		code, response := getFailedResponse(err)
		c.AbortWithStatusJSON(code, struct {
			Code    string `json:"code"`
			Message any    `json:"message"`
		}{
			Code:    response.ErrorCode.String(),
			Message: response.Message,
		})
		return
	}
	c.JSON(http.StatusOK, successResponse)
}

func (h *Handler) sendResponseCreated(c *gin.Context, successResponse any, err usecase.ErrorCode) {
	if successResponse == nil {
		if err == usecase.NoContent {
			c.Status(http.StatusNoContent)
			return
		}
		code, response := getFailedResponse(err)
		c.AbortWithStatusJSON(code, struct {
			Code    string `json:"code"`
			Message any    `json:"message"`
		}{
			Code:    response.ErrorCode.String(),
			Message: response.Message,
		})
		return
	}
	c.JSON(http.StatusCreated, successResponse)
}

func getFailedResponse(err usecase.ErrorCode) (int, usecase.FailedResponseBody) {
	failedResponse, isFound := usecase.ErrorCodeToFailedResponse[err]
	if !isFound {
		logrus.Error("the specified error code not found")
		return http.StatusInternalServerError, usecase.FailedResponseBody{
			ErrorCode: err,
			Message:   clerr.ErrorServer.Error(),
		}
	}

	return int(failedResponse.HttpCode), usecase.FailedResponseBody{
		ErrorCode: err,
		Message:   failedResponse.Message,
	}
}
