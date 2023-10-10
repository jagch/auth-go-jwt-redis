package user

import (
	"database/sql"
	"jagch/auth-go/model"

	"github.com/gin-gonic/gin"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (ur UserRepository) Create(c *gin.Context, ID, email, password string) (model.UserCreateResponse, error) {
	return model.UserCreateResponse{}, nil
}
func (ur UserRepository) GetByEmail(c *gin.Context, email string) (model.User, error) {
	return model.User{}, nil
}
