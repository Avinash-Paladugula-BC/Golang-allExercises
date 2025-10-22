package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"errors"
)

type Book struct {
	id       string `json:"id`
	title    string `json:"title`
	author   string `json:"author`
	quantity int    `json:"quantity`
}

var books = []Book{
{id: "1", title: "In Search of Lost Time", author: "Marcel Proust", quantity: 2},
	{id: "2", title: "The Great Gatsby", author: "F. Scott Fitzgerald", quantity: 5},
	{id: "3", title: "War and Peace", author: "Leo Tolstoy", quantity: 6},
}

// get endpoint
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "BOok not found"})
		return
	}

	if book.quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not available"})
		return
	}
	book.quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "BOok not found"})
		return
	}
	book.quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"messsage": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*Book, error) {
	for _, b := range books {
		if b.id == id {
			return &b, nil
		}
	}
	return nil, errors.New("book not found")
}

// post
func createBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	// curl localhost:8080/checkout?id=2 --request "PATCH"
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return",returnBook)
	router.Run("localhost:8080")
}
