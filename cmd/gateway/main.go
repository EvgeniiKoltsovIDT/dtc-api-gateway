package main

import (
	"context"
	"github.com/coretech/dtc-api-gateway/internal/app/gateway/transport/graphql/server"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	maxShutdownWait = 5 * time.Second
)

func main() {
	ctx := context.Background()

	srv := server.New()

	wg := &sync.WaitGroup{}

	wg.Add(1)

	go srv.Start(ctx)

	go handleGracefulShutdown(ctx, srv, wg)

	wg.Wait()

}

func handleGracefulShutdown(ctx context.Context, srv *server.Server, wg *sync.WaitGroup) {
	defer wg.Done()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-c

	timeoutCtx, cancel := context.WithTimeout(ctx, maxShutdownWait)
	defer cancel()

	slog.Debug("shutting down server", slog.Any("max_wait", maxShutdownWait))

	srv.Stop(timeoutCtx)
}
