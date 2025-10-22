package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func pathStrategyV1(r *gin.Engine) {
	r.GET("/v1/cars", func(c *gin.Context) {
		cars := []string{"1HGCM82633A123456", "JH4KA9650MC012345", "2FTRX18W1XCA98765"}
		c.JSON(http.StatusOK, gin.H{"cars": cars})
	})
}
