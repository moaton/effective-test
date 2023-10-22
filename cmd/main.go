package main

import (
	"context"
	"effective-test/config"
	"effective-test/internal/app"
	"effective-test/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cfg := config.GetConfig()

	logger.SetLogger(cfg.IsDebug)

	go app.Run(ctx, cfg)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	done := make(chan int, 1)

	go func() {
		<-sig
		cancel()
		done <- 1
	}()
	logger.Infof("Running...")
	<-done
	logger.Debugf("Graceful shutdown %v", time.Now().UTC())
}
