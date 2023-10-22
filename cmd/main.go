package main

import (
	"context"
	"effective-test/config"
	"effective-test/internal/app"
	"effective-test/pkg/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
)

var loggerSpace = "dev"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	zapLogger := logger.NewZapLogger(loggerSpace)
	logger.SetLogger(zapLogger)

	var cfg config.Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("envconfig.Process err %v", err)
	}

	go app.Run(ctx, &cfg)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	done := make(chan int, 1)

	go func() {
		<-sig
		cancel()
		done <- 1
	}()

	<-done
	logger.Debugf("Graceful shutdown %v", time.Now().UTC())
}
