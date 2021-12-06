package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//GET SERVICE
	router.GET("/", rootHandler)                     //endpoint root
	router.GET("/contacts", contactsHandler)         //endpoint /contacts
	router.GET("/books/:id", booksHandler)           //endpoints using path id
	router.GET("/books/:id/:title", manyPathHandler) //endpoints using many path
	router.GET("/query", queryHandler)               //endpoints query (books?categories=romance)
	//POST SERVICE

	router.Run(":8888")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Yasir Nur Prasetya",
		"job":  "students",
	})
}

func contactsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"whatsapp": "081229844969",
		"twitter":  "@yasirnurpras_",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func manyPathHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func queryHandler(c *gin.Context) {
	categories := c.Query("categories")
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{"categories": categories, "title": title})
}
