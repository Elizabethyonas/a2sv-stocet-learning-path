package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"a2sv_stocet_learning_path/internal/domain/contracts/usecase"
)

type CMSPController struct {
	Usecase usecase.CMSPUsecase
}

func NewCMSPController(u usecase.CMSPUsecase) *CMSPController {
	return &CMSPController{Usecase: u}
}

func (c *CMSPController) GetAll(ctx *gin.Context) {
	typeFilter := ctx.Query("type")
	beforeStr := ctx.Query("licensedBefore")
	afterStr := ctx.Query("licensedAfter")

	var before, after time.Time
	var err error

	if beforeStr != "" {
		before, err = time.Parse("2006-01-02", beforeStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid licensedBefore format (YYYY-MM-DD)"})
			return
		}
	}
	if afterStr != "" {
		after, err = time.Parse("2006-01-02", afterStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid licensedAfter format (YYYY-MM-DD)"})
			return
		}
	}

	result := c.Usecase.GetAll(typeFilter, before, after)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}

func (c *CMSPController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := c.Usecase.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "CMSP not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}
