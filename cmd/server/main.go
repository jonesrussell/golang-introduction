package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jonesrussell/go-fundamentals-best-practices/internal/api"
	"github.com/jonesrussell/go-fundamentals-best-practices/internal/executor"
	"github.com/jonesrussell/go-fundamentals-best-practices/internal/parser"
	"github.com/jonesrussell/go-fundamentals-best-practices/internal/storage"
)

// Server timeout constants.
const (
	readTimeout       = 15 * time.Second
	writeTimeout      = 15 * time.Second
	idleTimeout       = 60 * time.Second
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 30 * time.Second
)

// config holds application configuration.
type config struct {
	port         string
	tutorialsDir string
	dataDir      string
}

func main() {
	// Setup structured logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	if err := run(logger); err != nil {
		logger.Error("application failed", "error", err)
		os.Exit(1)
	}
}

// run contains the main application logic, separated for better testability.
func run(logger *slog.Logger) error {
	// Load configuration
	cfg := loadConfig()

	// Ensure data directory exists
	if err := os.MkdirAll(cfg.dataDir, 0o755); err != nil {
		return fmt.Errorf("create data directory: %w", err)
	}

	// Initialize components with explicit error handling
	tutorialParser := parser.NewTutorialParser(cfg.tutorialsDir)

	codeExecutor, err := executor.NewCodeExecutor()
	if err != nil {
		return fmt.Errorf("create code executor: %w", err)
	}
	defer func() {
		if cleanupErr := codeExecutor.Cleanup(); cleanupErr != nil {
			logger.Error("cleanup code executor failed", "error", cleanupErr)
		}
	}()

	progressStorage, err := storage.NewProgressStorage(cfg.dataDir)
	if err != nil {
		return fmt.Errorf("create progress storage: %w", err)
	}

	handlers, err := api.NewHandlers(tutorialParser, codeExecutor, progressStorage)
	if err != nil {
		return fmt.Errorf("create handlers: %w", err)
	}

	// Setup routes and middleware
	mux := handlers.SetupRoutes()
	handler := api.CORSMiddleware(mux)

	// Create HTTP server with secure defaults
	server := &http.Server{
		Addr:              ":" + cfg.port,
		Handler:           handler,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		ReadHeaderTimeout: readHeaderTimeout, // Protects against Slowloris attacks
	}

	// Setup graceful shutdown
	serverErrors := make(chan error, 1)
	go func() {
		logger.Info("server starting",
			"port", cfg.port,
			"tutorials_dir", cfg.tutorialsDir,
			"data_dir", cfg.dataDir,
		)

		serverErrors <- server.ListenAndServe()
	}()

	// Wait for shutdown signal or server error
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	select {
	case serverErr := <-serverErrors:
		if serverErr != nil && !errors.Is(serverErr, http.ErrServerClosed) {
			return fmt.Errorf("server error: %w", serverErr)
		}

	case sig := <-shutdown:
		logger.Info("shutdown signal received", "signal", sig.String())

		// Create shutdown context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if shutdownErr := server.Shutdown(ctx); shutdownErr != nil {
			// Force close if graceful shutdown times out
			if closeErr := server.Close(); closeErr != nil {
				return fmt.Errorf("server shutdown failed: %w; force close failed: %w", shutdownErr, closeErr)
			}
			return fmt.Errorf("server shutdown failed: %w", shutdownErr)
		}
	}

	logger.Info("server stopped cleanly")
	return nil
}

// loadConfig loads configuration from environment variables with defaults.
func loadConfig() config {
	return config{
		port:         getEnv("PORT", "8080"),
		tutorialsDir: getEnv("TUTORIALS_DIR", "tutorials"),
		dataDir:      getEnv("DATA_DIR", "data"),
	}
}

// getEnv retrieves an environment variable or returns a default value.
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
