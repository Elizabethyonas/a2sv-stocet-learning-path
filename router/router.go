package router

import (
	"a2sv_stocet_learning_path/api/controllers"
	"a2sv_stocet_learning_path/internal/application/usecase"
	usecaseImpl "a2sv_stocet_learning_path/internal/application/usecase"

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

	portfolioUsecase := usecase.NewPortfolioUsecase()
	portfolioController := controllers.NewPortfolioController(portfolioUsecase)

	api = r.Group("/api")
	{
		api.POST("/portfolio/recommend", portfolioController.Recommend)
	}

	econUsecase := usecaseImpl.NewEconUsecase()
	econController := controllers.NewEconController(econUsecase)
	api = r.Group("/api")
	{
		api.GET("/economic/indicators", econController.GetIndicators)
	}


	return r
}
