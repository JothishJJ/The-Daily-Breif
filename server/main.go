package main

import (
	"github.com/JothishJJ/thedailybrief/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func makeServer() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}

	r.Use(cors.New(config))

	return r
}

func main() {
	r := makeServer()
	router.Routes(r)
	r.Run()
}
