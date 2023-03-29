package controller

import (
	"apigorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controllers) CreateBook(ctx *gin.Context) {
	var newBook models.Books

	err := ctx.ShouldBindJSON(&newBook)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := c.masterDB.Create(&newBook).Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, newBook)
	return
}

func (c *Controllers) GetBook(ctx *gin.Context) {
	var getbook models.Books

	err := c.masterDB.Find(&getbook).Error

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, getbook)
	return

}
