package router

import (
	"koswarabilly/ic-service/internal/setup"

	"github.com/gin-gonic/gin"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
)

func buildPostPrompt(config setup.Configuration, logger zerolog.Logger, redisClient *goredislib.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
