package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"a2sv_stocet_learning_path/internal/domain/contracts/usecase"
)

type BondController struct {
	Usecase usecase.BondUsecase
}

func NewBondController(u usecase.BondUsecase) *BondController {
	return &BondController{Usecase: u}
}

func (c *BondController) SearchBonds(ctx *gin.Context) {
	minCoupon, _ := strconv.ParseFloat(ctx.Query("minCoupon"), 64)
	maxCoupon, _ := strconv.ParseFloat(ctx.Query("maxCoupon"), 64)

	var maturityAfter, maturityBefore time.Time
	var err error

	if after := ctx.Query("maturityAfter"); after != "" {
		maturityAfter, err = time.Parse("2006-01-02", after)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maturityAfter format. Use YYYY-MM-DD"})
			return
		}
	}
	if before := ctx.Query("maturityBefore"); before != "" {
		maturityBefore, err = time.Parse("2006-01-02", before)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maturityBefore format. Use YYYY-MM-DD"})
			return
		}
	}

	results := c.Usecase.SearchBonds(minCoupon, maxCoupon, maturityAfter, maturityBefore)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   results,
	})
}
