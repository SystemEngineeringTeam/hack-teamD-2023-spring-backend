package web

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
}

func Router(server Server) *gin.Engine {
	setMode()
	router := gin.Default()
	/*
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		}))
	*/
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})
	router.GET("/guest", GuestAccess)
	router.NoRoute(NotImplemented)
	return router
}
