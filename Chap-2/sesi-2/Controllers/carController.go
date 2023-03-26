package controllers

import (
	"net/http"
	models "sesi-2/Models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBook(ctx *gin.Context) {
	var newBook models.Book

	err := ctx.ShouldBindJSON(&newBook)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)

		return
	}

	newBookID := len(models.Bookdata) + 1
	newBook.BookID = newBookID
	models.Bookdata = append(models.Bookdata, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"Book": newBook,
	})
}

func UpadateBook(ctx *gin.Context) {
	var updatedBook models.Book
	bookID := ctx.Param("bookID")

	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = ctx.ShouldBindJSON(&updatedBook)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)

		return
	}

	isBookFound := false
	for _, book := range models.Bookdata {
		if book.BookID == bookIDInt {
			updatedBook.BookID = bookIDInt
			book = updatedBook
			isBookFound = true
			break
		}
	}
	if isBookFound {
		ctx.JSON(http.StatusOK, gin.H{
			"book": updatedBook,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Massage": "id is not valid",
	})
}
