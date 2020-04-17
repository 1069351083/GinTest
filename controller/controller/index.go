package controller

import (
	"ginTest/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ToIndex(ctx *gin.Context)  {
	categoryService, err := service.QuetyAllCategoryService()
	if err!=nil {
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"categoryList":categoryService,
	})
}
