package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura todas as rotas da aplicação
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Adicione mais rotas aqui

	return r
}

func ListRoutes(r *gin.Engine) {
	for _, route := range r.Routes() {
		log.Printf("Method: %s, Path: %s", route.Method, route.Path)
	}
}
