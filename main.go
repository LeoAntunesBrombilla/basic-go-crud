package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
)

type Item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var items []Item

func main() {

	router := gin.Default()

	router.POST("/create", createItem)
	router.GET("/read/:id", readItem)
	router.GET("/read", readItems)
	router.DELETE("/delete/:id", deleteItem)
	router.PUT("/update/:id", updateItem)

	router.Run(":8080")
}

func createItem(c *gin.Context) {
	var newItem Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	items = append(items, newItem)
	c.JSON(http.StatusOK, newItem)
}

func readItem(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ta faltando um parametro ai amigo"})
		return
	}

	for _, item := range items {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Nao achei"})
}

func readItems(c *gin.Context) {

	if len(items) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nao ha nenhum item na lista"})
		return
	}

	c.JSON(http.StatusOK, items)
}

func deleteItem(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Falta um parametro parceiro"})
		return
	}

	for _, item := range items {
		if item.ID == id {
			idxToRemove, err := strconv.Atoi(id)
			if err != nil {
				return
			}
			items = append(items[:idxToRemove], items[idxToRemove+1:]...)
		}
	}

}

func updateItem(c *gin.Context) {
	id := c.Param("id")
	var newItem Item

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Faltou o parametro de id"})
		return
	}

	for _, item := range items {
		if item.ID == id {
			idxToUpdate, _ := strconv.Atoi(id)
			if err := c.ShouldBindJSON(&newItem); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			items[idxToUpdate] = newItem
			c.JSON(http.StatusOK, newItem)
		}

	}
}
