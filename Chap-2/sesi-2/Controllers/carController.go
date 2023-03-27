package controllers

import (
	"fmt"
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
	for i, book := range models.Bookdata {
		if book.BookID == bookIDInt {
			updatedBook.BookID = bookIDInt
			book = updatedBook
			models.Bookdata[i] = book
			isBookFound = true
			break
		}
	}
	if isBookFound {
		ctx.JSON(http.StatusOK, gin.H{
			"Book": updatedBook,
		})
		return
	}
	ctx.JSON(http.StatusNotFound, gin.H{
		"Massage": "id is not valid",
	})

}

func GetBook(ctx *gin.Context) {
	var getbook models.Book
	bookID := ctx.Param("bookID")

	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	isBookFound := false
	for _, book := range models.Bookdata {
		if book.BookID == bookIDInt {
			isBookFound = true
			getbook = book
			break
		}
	}
	if isBookFound {
		ctx.JSON(http.StatusOK, gin.H{
			"Book": getbook,
		})
		return
	}
	ctx.JSON(http.StatusNotFound, gin.H{
		"Massage": fmt.Sprintf("id %s is not valid", bookID),
	})

}

func GetAllBook(ctx *gin.Context) {
	// var getallbook models.Book
	ctx.JSON(http.StatusOK, gin.H{
		"Book": models.Bookdata,
	})

}

func DeleteBook(ctx *gin.Context) {
	// var getbook models.Book
	bookID := ctx.Param("bookID")

	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	bookIndexToDelete := -1
	for i, book := range models.Bookdata {
		if book.BookID == bookIDInt {
			bookIndexToDelete = i
			break
		}
	}
	if bookIndexToDelete < 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Massage": fmt.Sprintf("id %s is not valid", bookID),
		})
		return
	}

	Length := len(models.Bookdata)
	copy(models.Bookdata[bookIndexToDelete:], models.Bookdata[bookIndexToDelete+1:])
	models.Bookdata[Length-1] = models.Book{}
	models.Bookdata = models.Bookdata[:Length]

	ctx.JSON(http.StatusOK, gin.H{
		"Massage": fmt.Sprintf("deleted %s", bookID),
	})

}
