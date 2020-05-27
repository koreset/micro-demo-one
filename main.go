package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/envy"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/health", getHealth)

	port := envy.Get("API_PORT", "9090")
	log.Fatal(router.Run(fmt.Sprintf(":%s", port)))
}

func getHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "api-1 is ok"})
}
