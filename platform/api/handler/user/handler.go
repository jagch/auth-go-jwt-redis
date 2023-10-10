package user

import (
	"fmt"
	"jagch/auth-go/internal/app/token"
	"jagch/auth-go/internal/domain"
	"jagch/auth-go/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	userUseCase  domain.UserUseCase
	tokenManager token.TokenManager
}

func newUserHandler(userUseCase domain.UserUseCase) userHandler {
	return userHandler{
		userUseCase: userUseCase,
	}
}

func (h userHandler) Create(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	// Receive the params from the request
	var userCreateRequest model.UserCreateRequest
	if err := c.BindJSON(&userCreateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Generate a password using the bcrypt package
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(userCreateRequest.Password), 8)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// UseCase
	userCreateResponse, err := h.userUseCase.Create(c, uuid.New().String(), userCreateRequest.Email, string(passwordBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, userCreateResponse)

}

func (h userHandler) Auth(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	// Receive the params from the request
	var userAuthrequest model.UserAuthRequest
	if err := c.BindJSON(&userAuthrequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// UseCase
	_, err := h.userUseCase.Auth(c, userAuthrequest.Email, userAuthrequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Welcome %s", userAuthrequest.Email)})
}
