package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/products/:sku", func(c *gin.Context) {
		sku := c.Params.ByName("sku")
		image, err := getImage(sku)
		if err != nil {
			c.JSON(http.StatusNotFound, "unknown SKU")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"imagePath": image,
			})
		}
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	r.Run("127.0.0.1:" + port)
}

func init() {
	db["IPOD-BLK"] = "https://i.ebayimg.com/images/i/162486149300-0-1/s-l1000.jpg"
	db["IPOD-RED"] = "https://static.decalgirl.com/assets/items/ipc/800/ipc-ss-red.jpg"
	db["IPOD-WHT"] = "https://gladallover.files.wordpress.com/2011/12/ipod-classic-white.jpg"
}

func getImage(sku string) (string, error) {
	if sku == "" {
		return "", errors.New("empty sku")
	}
	image := db[sku]
	if image == "" {
		return "", errors.New("invalid sku")
	}
	return image, nil
}
