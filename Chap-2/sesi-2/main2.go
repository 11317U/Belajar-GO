package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var books = []Book{}

func main() {
	router := gin.Default()
	// get All
	router.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, books)
	})

	// Get by id
	router.GET("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, Book := range books {
			if fmt.Sprintf("%v", Book.ID) == id {
				c.JSON(http.StatusOK, Book)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Books not found"})
	})

	//Add
	router.POST("/books", func(c *gin.Context) {
		var newBook Book
		if err := c.ShouldBindJSON(&newBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newBook.ID = len(books) + 1
		books = append(books, newBook)
		c.JSON(http.StatusCreated, gin.H{
			"Created Book ": newBook,
		})
	})

	// update
	router.PUT("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedBook Book
		if err := c.ShouldBindJSON(&updatedBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, book := range books {
			if fmt.Sprintf("%v", book.ID) == id {
				updatedBook.ID = book.ID
				books[i] = updatedBook
				c.JSON(http.StatusOK, gin.H{"message": "updated"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
	})

	// Delete
	router.DELETE("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, book := range books {
			if fmt.Sprintf("%v", book.ID) == id {
				books = append(books[:i], books[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Books not found"})
	})

	router.Run(":3000")
}
