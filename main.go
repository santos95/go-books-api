package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


// we add json to allow to convert to json
// exported file thats why in upper, json lowercase - when serialized turn the ID into id and vicerversa
type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Caballo De Troya", Author: "J.J Benitez", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "La Rebelion de Lucifer", Author: "J.J Benitez", Quantity: 1},
	{ID: "4", Title: "El Enviado", Author: "J.J Benitez", Quantity: 3},
	{ID: "5", Title: "Space Odyssey", Author: "Arthur C. Clarke", Quantity: 3},
	{ID: "6", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 3},
}

func getBooks(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, books)
}

func createBooks(c *gin.Context) {

	var newBook book 

	if err := c.BindJSON(&newBook); err != nil {

		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {

	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBooks)
	router.Run("localhost:9090")
}
