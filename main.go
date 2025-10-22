package main

import (
	"github.com/gin-gonic/gin"
	"log"
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
