package main

import (
	"net/http"

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

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	router.Routes(r)

	r.Run()
}
