package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/products/:sku", func(c *gin.Context) {
		sku := c.Params.ByName("sku")
		image, err1 := getImage(sku)
		price, err2 := getPrice(sku)
		if err1 != nil {
			log.Println("Error fetching image")
			c.JSON(500, "")
		} else if err2 != nil {
			log.Println("Error fetching price")
			c.JSON(500, "")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"image": image,
				"price": price,
			})
		}
	})
	r.Run()
}

func getImage(sku string) (interface{}, error) {
	log.Println("Fetching image")
	url := fmt.Sprintf("%v://%v:%v/products/%v", "http", "localhost", "8081", sku)
	resp, err := http.Get(url)
	if err != nil {
		log.Panic(err)
		return "", err
	}
	log.Println("Hooray")
	defer resp.Body.Close()
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result["imagePath"], nil
}

func getPrice(sku string) (interface{}, error) {
	log.Println("Fetching price")
	url := fmt.Sprintf("%v://%v:%v/products/%v", "http", "localhost", "8082", sku)
	resp, err := http.Get(url)
	if err != nil {
		log.Panic(err)
		return 0, err
	}
	log.Println("Hooray")
	defer resp.Body.Close()
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result["price"], nil
}
