package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}

var products []Product

func loadProducts() {
	// Load products from a JSON file
	file, _ := os.Open("products.json")
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &products)
}

func saveProducts() {
	// Save products to a JSON file
	file, _ := os.Create("products.json")
	defer file.Close()
	bytes, _ := json.MarshalIndent(products, "", "  ")
	file.Write(bytes)
}

func main() {
	r := gin.Default()

	loadProducts()

	r.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, products)
	})

	r.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, product := range products {
			if fmt.Sprintf("%d", product.ID) == id {
				c.JSON(http.StatusOK, product)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	})

	r.POST("/products", func(c *gin.Context) {
		var newProduct Product
		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newProduct.ID = len(products) + 1
		products = append(products, newProduct)
		saveProducts()
		c.JSON(http.StatusCreated, newProduct)
	})

	r.PUT("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedProduct Product
		if err := c.ShouldBindJSON(&updatedProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, product := range products {
			if fmt.Sprintf("%d", product.ID) == id {
				updatedProduct.ID = product.ID
				products[i] = updatedProduct
				saveProducts()
				c.JSON(http.StatusOK, updatedProduct)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	})

	r.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, product := range products {
			if fmt.Sprintf("%d", product.ID) == id {
				products = append(products[:i], products[i+1:]...)
				saveProducts()
				c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	})

	r.Run()
}