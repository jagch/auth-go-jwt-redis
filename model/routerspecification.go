package model

import (
	"database/sql"
	"jagch/auth-go/internal/app/token"

	"github.com/gin-gonic/gin"
)

type RouterSpecification struct {
	Api          *gin.Engine
	Logger       Logger
	DB           *sql.DB
	TokenManager token.TokenManager
}
