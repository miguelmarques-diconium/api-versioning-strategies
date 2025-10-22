package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
