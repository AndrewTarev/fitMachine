package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fitMachine/internal/config"
	appconfig "fitMachine/pkg/config"
	"fitMachine/pkg/logger"
	"fitMachine/pkg/sorre"
)

// Server представляет HTTP сервер
type Server struct {
	httpServer *http.Server
	config     appconfig.IConfig
	logger     logger.ILogger
	ctx        context.Context
}

func New(ctx context.Context, cfg appconfig.IConfig, logger logger.ILogger) *Server {
	return &Server{
		ctx:    ctx,
		config: cfg,
		logger: logger,
	}
}

// SetupRoutes настраивает маршруты API
func (s *Server) SetupRoutes() {
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	// API routes будут добавлены здесь
	mux.HandleFunc("/api/v1/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	s.httpServer = &http.Server{
		Addr:         s.getServerAddress(),
		ReadTimeout:  s.config.GetDuration(config.ServerReadTimeout),
		WriteTimeout: s.config.GetDuration(config.ServerWriteTimeout),
		Handler:      mux,
	}
}

// getServerAddress возвращает адрес сервера
func (s *Server) getServerAddress() string {
	host := s.config.GetString(config.ServerHost)
	port := s.config.GetString(config.ServerPort)
	return fmt.Sprintf("%s:%s", host, port)
}

func (s *Server) start() error {
	if s.httpServer == nil {
		return sorre.Unwrap(fmt.Errorf("server not configured, call SetupRoutes first"))
	}

	go func() {
		s.logger.Info(s.ctx, "Server starting", "address", s.getServerAddress())
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error(s.ctx, "Server failed to start", "error", err)
		}
	}()

	return nil
}

func (s *Server) waitForShutdown() {
	// Ожидаем сигнал для graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	s.logger.Info(s.ctx, "Shutting down server...")

	// Graceful shutdown с таймаутом
	shutdownCtx, cancel := context.WithTimeout(s.ctx, 30*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
		s.logger.Error(s.ctx, "Server forced to shutdown", "error", err)
	}

	s.logger.Info(s.ctx, "Server exited")
}

// Run запускает сервер и ожидает shutdown
func (s *Server) Run() error {
	if err := s.start(); err != nil {
		return err
	}

	s.waitForShutdown()
	return nil
}
