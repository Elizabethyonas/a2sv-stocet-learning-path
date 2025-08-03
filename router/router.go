package router

import (
	"github.com/gin-gonic/gin"
	"a2sv_stocet_learning_path/api/controllers"
	"a2sv_stocet_learning_path/internal/application/usecase"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Equities endpoint
	equityUsecase := usecase.NewEquityUsecase()
	equityController := controllers.NewEquityController(equityUsecase)
	r.GET("/api/equities", equityController.GetEquities)

	return r
}
