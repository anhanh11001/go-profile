package main

import (
	"fmt"

	"github.com/anhanh11001/go-profile/middlewares"
	"github.com/anhanh11001/go-profile/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server at 8080")

	server := gin.Default()

	middlewares.SetUpMongoDb()
	routers.SetUpRouters(server)

	server.Run() // Run at 8080
}
