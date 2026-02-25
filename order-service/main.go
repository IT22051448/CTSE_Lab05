package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	orders := []gin.H{}
	idCounter := 1

	// Getting All Orders

	r.GET("/orders", func(c *gin.Context) {
		c.JSON(http.StatusOK, orders)
	})

	// Creating an Order

	r.POST("/orders", func(c *gin.Context) {
		var order map[string]interface{}

		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		order["id"] = idCounter
		order["status"] = "PENDING"
		idCounter++

		orders = append(orders, order)
		c.JSON(http.StatusCreated, order)
	})

	// Getting an Order by ID
	
	r.GET("/orders/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		for _, order := range orders {
			if orderID, ok := order["id"].(int); ok && orderID == id {
				c.JSON(http.StatusOK, order)
				return
			}
		}

		c.Status(http.StatusNotFound)
	})

	r.Run(":8082")
}