package main

import (
	"context"
	"gokit-example/config"
	"gokit-example/pkg/gokit/endpoint"
	"gokit-example/pkg/gokit/service"
	"gokit-example/pkg/gokit/transport"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	conn := config.ConnectToDB()
	if conn == nil {
		logger.Panic("Failed trying to connect Postgres...")
	}

	{
		svc := service.NewUsrService(logger, conn)

		endpoints := endpoint.NewUserSvcSet(svc, logger)
		httpHandler := transport.NewUserHTTPHandler(endpoints, logger)
		http.Handle("/", httpHandler)
	}

	s := &http.Server{
		Addr: ":8080",
		// IdleTimeout:  120 * time.Second,
		// ReadTimeout:  5 * time.Second,
		// WriteTimeout: 5 * time.Second,
	}

	// wrapping ListenAndServe in gofunc so it's not going to block
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			logger.Error("Error:", zap.Any("Server start error", err))
		}
	}()

	// make a new channel to notify on os interrupt of server (ctrl + C)
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// This blocks the code until the channel receives some message
	sig := <-sigChan
	logger.Info("Received terminate, graceful shutdown", zap.Any("channel", sig))

	// Once message is consumed shut everything down
	// Gracefully shuts down all client requests. Makes server more reliable
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)

}
