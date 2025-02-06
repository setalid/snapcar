package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/setalid/snapcar/api/pkg/api"
	"github.com/setalid/snapcar/api/pkg/config"
	"github.com/setalid/snapcar/api/pkg/core"
	"github.com/setalid/snapcar/api/pkg/storage/memory"
	"go.uber.org/zap"
)

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	defer logger.Sync()

	var cfg config.Config
	err = envconfig.Process("", &cfg)
	if err != nil {
		return err
	}

	db := memory.New()

	// Repositories
	rentalRepo := memory.NewRentalRepo(db)
	carCategoryRepo := memory.NewCarCategoryRepo(db)
	rentalRateRepo := memory.NewRentalRateRepo(db)

	// Services
	rentalSvc := core.NewRentalService(
		rentalRepo,
		carCategoryRepo,
		rentalRateRepo,
	)

	rootHandler := api.NewRootHandler(
		logger,
		rentalSvc,
	)

	httpServer := &http.Server{
		Addr:    cfg.HTTPListenAddr,
		Handler: rootHandler,
	}

	go func() {
		logger.Sugar().Infof("listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()

	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
