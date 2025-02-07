package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/setalid/snapcar/api/pkg/config"
	"github.com/setalid/snapcar/api/pkg/core"
	"github.com/setalid/snapcar/api/pkg/storage/memory"
	"go.uber.org/zap"
)

func Run(ctx context.Context) error {
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
	carCategoryRepo := memory.NewCarCategoryRepo(db)
	rentalRepo := memory.NewRentalRepo(db, carCategoryRepo)
	rentalRateRepo := memory.NewRentalRateRepo(db)

	// Services
	rentalSvc := core.NewRentalService(
		rentalRepo,
		carCategoryRepo,
		rentalRateRepo,
	)

	// Initial global settings - ignore errors
	rentalRateRepo.Create(ctx, core.NewRentalRate(1, 1))

	// Initial categories - ignore errors
	carCategoryRepo.Create(ctx, core.SmallCar())
	carCategoryRepo.Create(ctx, core.Combi())
	carCategoryRepo.Create(ctx, core.Truck())

	rootHandler := NewRootHandler(
		logger,
		carCategoryRepo,
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

func encode[T any](w http.ResponseWriter, _ *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}
