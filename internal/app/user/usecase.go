package user

import (
	"jagch/auth-go/internal/app/token"
	"jagch/auth-go/internal/domain"
	"jagch/auth-go/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepository domain.UserRepository
	tokenManager   token.TokenManager
}

func NewUserUseCase(userRepository domain.UserRepository, tokenManager token.TokenManager) UserUseCase {
	return UserUseCase{
		userRepository: userRepository,
		tokenManager:   tokenManager,
	}
}

func (au UserUseCase) Create(c *gin.Context, ID, email, password string) (model.UserCreateResponse, error) {
	return au.userRepository.Create(c, ID, email, password)
}

func (au UserUseCase) Auth(c *gin.Context, email, password string) (model.UserAuthResponse, error) {
	user, err := au.userRepository.GetByEmail(c, email)
	if err != nil {
		return model.UserAuthResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email y/o password are incorrect"})
	}

	tokenString, err := au.tokenManager.Create(user.ID)
	if err != nil {
		return model.UserAuthResponse{}, err
	}

	return model.UserAuthResponse{
		ID:    user.ID,
		Email: user.Email,
		Token: tokenString,
	}, nil
}
