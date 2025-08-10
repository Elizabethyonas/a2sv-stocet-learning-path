package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"a2sv_stocet_learning_path/internal/domain/contracts/usecase"
)

type EconController struct {
	Usecase usecase.EconUsecase
}

func NewEconController(u usecase.EconUsecase) *EconController {
	return &EconController{Usecase: u}
}

// GET /api/economic/indicators?from=2019&to=2024
func (ec *EconController) GetIndicators(c *gin.Context) {
	fromStr := c.Query("from")
	toStr := c.Query("to")

	if fromStr == "" && toStr == "" {
		// return all
		resp, err := ec.Usecase.GetAllIndicators()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	from := 0
	to := 0
	var err error
	if fromStr != "" {
		from, err = strconv.Atoi(fromStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid 'from' year"})
			return
		}
	}
	if toStr != "" {
		to, err = strconv.Atoi(toStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid 'to' year"})
			return
		}
	}

	resp, err := ec.Usecase.GetIndicatorsForRange(from, to)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
