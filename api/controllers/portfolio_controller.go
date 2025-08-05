package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"a2sv_stocet_learning_path/internal/domain/contracts/usecase"

	"a2sv_stocet_learning_path/internal/domain/entities"
)

type PortfolioController struct {
	Usecase usecase.PortfolioUsecase
}

func NewPortfolioController(u usecase.PortfolioUsecase) *PortfolioController {
	return &PortfolioController{
		Usecase: u,
	}
}

func (pc *PortfolioController) Recommend(c *gin.Context) {
	var req entities.RecommendationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result, err := pc.Usecase.RecommendPortfolio(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
