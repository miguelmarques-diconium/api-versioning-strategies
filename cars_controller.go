package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

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
