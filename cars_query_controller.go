package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
