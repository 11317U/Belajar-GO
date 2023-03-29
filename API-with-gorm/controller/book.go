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
}

func (c *Controllers) GetBook(ctx *gin.Context) {
	var getbook []models.Books

	err := c.masterDB.Find(&getbook).Error

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, getbook)

}

func (c *Controllers) DeleteBook(ctx *gin.Context) {
	bookmodel := models.Books{}
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"massage": "Id not found",
		})
		return
	}
	err := c.masterDB.Delete(bookmodel, id).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"massage": "delete cannot do",
		})
		// return err
	}
	ctx.JSON(http.StatusOK, gin.H{
		"massage": "deleted",
	})
	// return nil

}

func (c *Controllers) GetBookid(ctx *gin.Context) {
	getidbook := models.Books{}
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"massage": "Id cannot empty",
		})
		return
	}
	err := c.masterDB.Where("id = ?", id).First(getidbook)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"massage": "Id not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, getidbook)
	// return
}
