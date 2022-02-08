package router

import (
	"koswarabilly/ic-service/internal/setup"
	"net/http"

	"github.com/gin-gonic/gin"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
)

type RouterPayload struct {
	Config      setup.Configuration
	Logger      zerolog.Logger
	RedisClient *goredislib.Client
}

func SetupRouter(rp RouterPayload) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	router.POST("/prompt", buildPostPrompt(rp.Config, rp.Logger, rp.RedisClient))
	return router
}
