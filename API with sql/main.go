package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "localhost"
	DB_NAME     = "dbbook"
)

var (
	db  *sql.DB
	err error
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

func main() {
	psqlUrl := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`, DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err = sql.Open("postgres", psqlUrl)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to postgres :)")

	router := gin.Default()

	router.GET("/books", GetBook)
	router.POST("/books", CreateBook)
	// router.PUT("/employees/:id", UpdateBook)
	// router.DELETE("/employees/:id", DeleteBook)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

func CreateBook(c *gin.Context) {
	book := Book{}

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO books (title, author, "desc")
		VALUES ($1, $2, $3)
		RETURNING *
	`

	err = db.QueryRow(query, book.Title, book.Author, book.Desc).
		Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func GetBook(c *gin.Context) {
	books := []Book{}

	query := "SELECT * FROM books"

	rows, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		book := Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

// func UpdateBook(c *gin.Context) {
// 	book := Book{}

// 	err := c.ShouldBindJSON(&book)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	id := c.Param("id")

// 	query := `
