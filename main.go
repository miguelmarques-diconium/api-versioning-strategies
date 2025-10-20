package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func main() {
	r := gin.Default()
	errProxy := r.SetTrustedProxies(nil)
	if errProxy != nil {
		return
	}

	pathStrategy(r)

	err := r.Run(":8080")
	if err != nil {
		log.Println("Failed to start server:", err)
		return
	}
}

func pathStrategy(r *gin.Engine) {
	r.GET("/v1/cars", func(c *gin.Context) {
		cars := []string{"CarA", "CarB", "CarC"}
		c.JSON(http.StatusOK, gin.H{"cars": cars})
	})
}

func headerStrategy(r *gin.Engine) {
	r.GET("/cars", func(c *gin.Context) {
		cars := []string{"CarA", "CarB", "CarC"}
		c.JSON(http.StatusOK, gin.H{"cars": cars})
	})
}


func pickVersionFromAccept(hdrs []string) string {
	// Look for version parameter or vendor subtype with .v2 for _, h := range hdrs
	if strings.Contains(h, "version=2") || strings.Contains(h, ".v2+json") {
		return "2" }
return "1"
}
	func tempMain() {
		r := chi.NewRouter()
		r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "id")
			v := pickVersionFromAccept(r.Header.Values("Accept"))
			if v == "2" {
				w.Header().Set("Content-Type", "application/vnd.acme.users.v2+json") json.NewEncoder(w).Encode(map[string]any{"id": id, "full_name": "Alice Smith"})
				return } w.Header().Set("Content-Type", "application/vnd.acme.users.v1+json") json.NewEncoder(w).Encode(map[string]any{"id": id, "name": "Alice"}) }) http.ListenAndServe(":8080", r)
	}
