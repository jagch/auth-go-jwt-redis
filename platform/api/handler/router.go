package handler

import (
	"jagch/auth-go/model"
	userHandler "jagch/auth-go/platform/api/handler/user"
)

func InitRoutes(specification model.RouterSpecification) {
	userHandler.NewRouter(specification)
}
