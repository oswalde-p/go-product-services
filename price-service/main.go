package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]int)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/products/:sku", func(c *gin.Context) {
		sku := c.Params.ByName("sku")
		price, err := getPrice(sku)
		if err != nil {
			c.JSON(http.StatusNotFound, "unknown SKU")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"price":          price,
				"formattedPrice": fmt.Sprintf("AU$%.2f", float64(price)/float64(100)),
			})
		}
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}
	r.Run("127.0.0.1:" + port)
}

func init() {
	db["IPOD-BLK"] = 19999
	db["IPOD-RED"] = 29999
	db["IPOD-WHT"] = 23999
}

func getPrice(sku string) (int, error) {
	if sku == "" {
		return 0, errors.New("empty sku")
	}
	price := db[sku]
	if price == 0 {
		return 0, errors.New("unknown sku")
	}
	return price, nil
}
