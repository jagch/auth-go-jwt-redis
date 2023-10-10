package domain

import (
	"jagch/auth-go/model"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	Create(c *gin.Context, ID, email, password string) (model.UserCreateResponse, error)
	GetByEmail(c *gin.Context, email string) (model.User, error)
}

type UserUseCase interface {
	Create(c *gin.Context, ID, email, password string) (model.UserCreateResponse, error)
	Auth(c *gin.Context, email, password string) (model.UserAuthResponse, error)
}
