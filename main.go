package main

import (
	"test-fase-3/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.CombineRouter(r)
	r.Run("localhost:8080")
}
