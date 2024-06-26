package router

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/", HomeHandler)

	r.GET("/top", TopHandler)
	r.GET("/world", WorldHandler)
	r.GET("/india", IndiaHandler)
}
