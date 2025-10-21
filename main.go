package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type car struct {
	VIN   string `json:"vin"`
	Brand string `json:"brand"`
	Model string `json:"model"`
}

func main() {
	r := gin.Default()
	errProxy := r.SetTrustedProxies(nil)
	if errProxy != nil {
		return
	}

	pathStrategyV1(r)
	pathStrategyV2(r)
	headerStrategy(r)
	queryParameterStrategy(r)

	err := r.Run(":8080")
	if err != nil {
		log.Println("Failed to start server:", err)
		return
	}
}

func pathStrategyV1(r *gin.Engine) {
	r.GET("/v1/cars", func(c *gin.Context) {
		cars := []string{"1HGCM82633A123456", "JH4KA9650MC012345", "2FTRX18W1XCA98765"}
		c.JSON(http.StatusOK, gin.H{"cars": cars})
	})
}

func pathStrategyV2(r *gin.Engine) {
	r.GET("/v2/cars", func(c *gin.Context) {
		cars := []car{
			{VIN: "1HGCM82633A123456", Brand: "Honda", Model: "Accord"},
			{VIN: "JH4KA9650MC012345", Brand: "Acura", Model: "Legend"},
			{VIN: "2FTRX18W1XCA98765", Brand: "Ford", Model: "F-150"},
		}
		c.JSON(http.StatusOK, gin.H{"cars": cars})
	})
}

func headerStrategy(r *gin.Engine) {
	r.GET("/cars", func(c *gin.Context) {
		acceptHeader := c.GetHeader("Accept")
		version := pickVersionFromAccept(acceptHeader)
		if version == "2" {
			cars := []car{
				{VIN: "1HGCM82633A123456", Brand: "Honda", Model: "Accord"},
				{VIN: "JH4KA9650MC012345", Brand: "Acura", Model: "Legend"},
				{VIN: "2FTRX18W1XCA98765", Brand: "Ford", Model: "F-150"},
			}
			c.JSON(http.StatusOK, gin.H{"cars": cars})
			return
		}

		if version == "1" {
			cars := []string{"1HGCM82633A123456", "JH4KA9650MC012345", "2FTRX18W1XCA98765"}
			c.JSON(http.StatusOK, gin.H{"cars": cars})
			return
		}

		// We could also return an error instead of defaulting to latest
	})
}

func queryParameterStrategy(r *gin.Engine) {
	r.GET("/carsQuery", func(c *gin.Context) {
		version := c.Query("version")
		if version == "2" || version == "" {
			cars := []car{
				{VIN: "1HGCM82633A123456", Brand: "Honda", Model: "Accord"},
				{VIN: "JH4KA9650MC012345", Brand: "Acura", Model: "Legend"},
				{VIN: "2FTRX18W1XCA98765", Brand: "Ford", Model: "F-150"},
			}
			c.JSON(http.StatusOK, gin.H{"cars": cars})
			return
		}

		if version == "1" {
			cars := []string{"1HGCM82633A123456", "JH4KA9650MC012345", "2FTRX18W1XCA98765"}
			c.JSON(http.StatusOK, gin.H{"cars": cars})
			return
		}
	})
}

func pickVersionFromAccept(acceptHeaderValue string) string {

	// Default to latest if no Accept header is provided
	if acceptHeaderValue == "" {
		return "2"
	}

	if strings.Contains(acceptHeaderValue, "version=2") || strings.Contains(acceptHeaderValue, ".v2+json") {
		return "2"
	}
	return "1"
}
