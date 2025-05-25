package handler

import (
	"devSystem/internal/usecase"
	"devSystem/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

const (
	errorGetAuthKeyFromHeader = "unable to get authorization token from header"
	errorHeaders              = "an error occurred while working with headers"
	errorDocumentation        = "documentation is only available in development mode"
)

type Middlewares interface {
	OnlyDevelopModeMiddleware(c *gin.Context)
	UserIdentityMiddleware(c *gin.Context)
}

type RequestHeaders struct {
	Authorization   string `header:"Token"`
	ClientRequestId string `header:"Client-Request-ID" binding:"required,uuid"`
}

type ResponseHeaders struct {
	RequestTimeString string `header:"Timestamp" binding:"required,datetime=Mon, 02 Jan 2006 15:04:05 MST"`
	RequestTime       time.Time
	ClientRequestId   string `header:"Client-Request-ID" binding:"required,uuid"`
	RequestId         string `header:"Request-ID" binding:"required,uuid"`
}

func (h *Handler) OnlyDevelopModeMiddleware(c *gin.Context) {
	if gin.Mode() != gin.DebugMode {
		c.AbortWithStatusJSON(http.StatusForbidden, errorDocumentation)
	}
}

func (h *Handler) UserIdentityMiddleware(c *gin.Context) {
	if c.Request.Method == http.MethodOptions {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return
	}

	token := strings.TrimSpace(authHeader[len(bearerPrefix):])
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is empty"})
		return
	}

	// Логирование для отладки
	logrus.Infof("Authorization token received: %v", token)

	claims, err := h.usecases.ParseToken(token)
	if err != usecase.Success {
		logrus.Errorf("Error parsing token: %v", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	logrus.Infof("Token successfully parsed. Claims: %+v", claims)

	c.Set("claims", claims)

	setUserContext(c, claims.Email)

	c.Next()
}

func (h *Handler) GetJWTClaims(c *gin.Context) (*models.JWTClaims, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return nil, errors.New("Missing Authorization header")
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return nil, errors.New("Invalid Authorization header")
	}

	token := strings.TrimSpace(authHeader[len(bearerPrefix):])
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is empty"})
		return nil, errors.New("Invalid Authorization header")
	}

	claims, err := h.usecases.ParseToken(token)
	if err != usecase.Success {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return nil, errors.New("Unauthorized")
	}

	return claims, nil
}

func setUserContext(c *gin.Context, email string) {
	c.Set(gin.AuthUserKey, email)
}

func (h *Handler) DEPRECATED(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusGone, nil)
}
