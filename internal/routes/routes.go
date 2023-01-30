package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hhelix/baulkham/internal/config"
	"github.com/hhelix/baulkham/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func Routes(app *config.Application) *gin.Engine {
	router := gin.Default()
	corsConfig := cors.Config{
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:  []string{"Accept", "Content-Type", "X-CSRF-Token", "Authorization"},
		ExposeHeaders: []string{"*"},
		//TODO: Change this to the actual domain origin
		AllowAllOrigins:  true,
		AllowCredentials: true,
	}

	err := cors.Config.Validate(corsConfig)
	if err != nil {
		log.Error("Failed to validate CORS: Please check getCORS()")
		return nil
	}

	router.Use(cors.New(corsConfig))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/query", handlers.Repo.GraphqlHandler())
	router.GET("/playground", handlers.Repo.PlaygroundHandler()) // !! TODO remove this shit in prod

	return router
}
