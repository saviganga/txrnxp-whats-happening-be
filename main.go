package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"txrnxp-whats-happening/api"
	configs "txrnxp-whats-happening/config"
	media "txrnxp-whats-happening/external/media/files"
	"txrnxp-whats-happening/internal/database"
	services "txrnxp-whats-happening/internal/services"
	mediaService "txrnxp-whats-happening/internal/services/media"

	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()

	env := flag.String("env", "local", "Environment setting")
	flag.Parse()

	if env == nil {
		panic("-env=value flag is required with allowed values (production, local, staging)")
	}

	if err := run(ctx, *env); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "startup error: %s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, envArg string) error {
	// --- Config ---
	cfg := configs.NewConfig()
	if err := cfg.LoadConfig(); err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	fmt.Printf("config loaded for environment: %s\n", envArg)

	// --- DB ---
	db, err := configs.NewDatabase(cfg)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()
	log.Info()

	repo := database.NewGormRepository(db.DB)

	mediaProvider := media.NewImageKit(cfg)
	mediaService := mediaService.NewMediaService(mediaProvider)

	// --- Service Bundle ---
	serviceBundle := services.NewBundle(repo, cfg, mediaProvider, *mediaService)
	log.Info()

	// --- HTTP Server ---
	srv := &http.Server{
		Addr:         net.JoinHostPort("0.0.0.0", "8080"),
		Handler:      api.NewServer(*serviceBundle),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// --- Start server ---
	go func() {
		log.Info()
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error()
		}
	}()

	// --- Graceful shutdown ---
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig

		log.Info()
		ctxShutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctxShutdown); err != nil {
			log.Error()
		}
	}()

	wg.Wait()
	return nil
}
