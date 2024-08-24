package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	register(router)
	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
