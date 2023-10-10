package boostrap

import (
	"jagch/auth-go/internal/app/token"
	"jagch/auth-go/model"
	"jagch/auth-go/platform/api/handler"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Run() error {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("error loading .env file")
	}

	// Engine
	var api *gin.Engine

	// Logger
	logger, err := newLogger()
	if err != nil {
		log.Fatal("error loading logger")
	}

	// Database
	db := newDatabase()

	// More Deps
	tokenManager := token.NewTokenJWT(os.Getenv("JWT_SECRETKEY"))

	// router
	handler.InitRoutes(model.RouterSpecification{
		Api:          api,
		Logger:       logger,
		DB:           db,
		TokenManager: tokenManager,
	})

	return nil
}
