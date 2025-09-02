package main

import (
	"Todo_List/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)

	fmt.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}
