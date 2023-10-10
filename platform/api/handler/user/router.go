package user

import (
	userUseCase "jagch/auth-go/internal/app/user"
	"jagch/auth-go/model"
	userRepository "jagch/auth-go/platform/repository/postgres/user"

	"github.com/gin-gonic/gin"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(handler, specification.Api)
}

func buildHandler(specification model.RouterSpecification) userHandler {
	useCase := userUseCase.NewUserUseCase(userRepository.NewUserRepository(specification.DB), specification.TokenManager)

	return newUserHandler(useCase)
}

func publicRoutes(h userHandler, api *gin.Engine, middlewares ...gin.HandlerFunc) {
	routes := api.Group("v1/user", middlewares...)

	routes.POST("", h.Create)
	routes.POST("/auth", h.Auth)
}
