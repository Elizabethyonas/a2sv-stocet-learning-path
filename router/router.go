package router

import (
	"a2sv_stocet_learning_path/api/controllers"
	"a2sv_stocet_learning_path/internal/application/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Equities endpoint
	equityUsecase := usecase.NewEquityUsecase()
	equityController := controllers.NewEquityController(equityUsecase)
	r.GET("/api/equities", equityController.GetEquities)

	bondUsecase := usecase.NewBondUsecase()
	bondController := controllers.NewBondController(bondUsecase)

	api := r.Group("/api/bonds")
	{
		api.GET("/search", bondController.SearchBonds)
	}

	// CMSP
	cmspUsecase := usecase.NewCMSPUsecase()
	cmspController := controllers.NewCMSPController(cmspUsecase)

	api = r.Group("/api/cmsps")
	{
		api.GET("", cmspController.GetAll)
		api.GET("/:id", cmspController.GetByID)
	}

	return r
}
