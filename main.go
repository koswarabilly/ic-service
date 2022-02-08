package main

import (
	"context"
	"fmt"
	"koswarabilly/ic-service/internal/router"
	"koswarabilly/ic-service/internal/setup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	config := setup.Configuration{}
	err := envconfig.Process("", &config)
	if err != nil {
		panic(fmt.Errorf("fail to initialize configuration: %w", err).Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	logger := setup.SetupLogger(config)
	redisClient := setup.SetupRedis(config)

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		panic(fmt.Errorf("fail to initialize redis").Error())
	}
	routerPayload := router.RouterPayload{
		Config:      config,
		Logger:      logger,
		RedisClient: redisClient,
	}
	router := router.SetupRouter(routerPayload)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.ServicePort),
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Msgf("listen: %s\n", err)
		}
	}()
	logger.Info().Msgf("server started on port %d...", config.ServicePort)

	quit := make(chan os.Signal, 2)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)
	<-quit

	defer cancel()

	logger.Debug().Msg("shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal().AnErr("graceful server shutdown timed out", err).Msg("")
	}
	logger.Info().Msg("server exited")
}
