package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"a2sv_stocet_learning_path/internal/domain/contracts/usecase"
)

type EquityController struct {
	Usecase usecase.EquityUsecase
}

func NewEquityController(u usecase.EquityUsecase) *EquityController {
	return &EquityController{Usecase: u}
}

func (c *EquityController) GetEquities(ctx *gin.Context) {
	filterType := ctx.Query("type") // gainer or loser
	result := c.Usecase.GetEquities(filterType)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}
