package setup

import (
	"fmt"

	goredislib "github.com/go-redis/redis/v8"
)

func SetupRedis(config Configuration) *goredislib.Client {
	return goredislib.NewClient(&goredislib.Options{
		Addr:     fmt.Sprintf("%s:%d", config.RedisHost, config.RedisPort),
		Password: config.RedisPassword,
		DB:       config.RedisDatabaseNumber,
	})
}
