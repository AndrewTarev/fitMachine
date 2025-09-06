package logger

import (
	"context"
	"log/slog"
	"os"

	constConfig "fitMachine/internal/config"
	"fitMachine/pkg/config"
)

var _ ILogger = (*Logger)(nil)

type Logger struct {
	*slog.Logger
}

func New(cfg config.IConfig) *Logger {
	// Определяем уровень логирования
	var level slog.Level
	switch cfg.GetString(constConfig.LogLevel) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	var handler slog.Handler
	if cfg.GetString(constConfig.LogFormat) == "json" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
	}

	return &Logger{
		Logger: slog.New(handler),
	}
}

// Debug логирует сообщение уровня debug с контекстом
func (l *Logger) Debug(ctx context.Context, msg string, args ...any) {
	l.Logger.DebugContext(ctx, msg, args...)
}

// Info логирует сообщение уровня info с контекстом
func (l *Logger) Info(ctx context.Context, msg string, args ...any) {
	l.Logger.InfoContext(ctx, msg, args...)
}

// Warn логирует сообщение уровня warning с контекстом
func (l *Logger) Warn(ctx context.Context, msg string, args ...any) {
	l.Logger.WarnContext(ctx, msg, args...)
}

// Error логирует сообщение уровня error с контекстом
func (l *Logger) Error(ctx context.Context, msg string, args ...any) {
	l.Logger.ErrorContext(ctx, msg, args...)
}

// DebugAttrs логирует сообщение уровня debug с атрибутами
func (l *Logger) DebugAttrs(ctx context.Context, msg string, args ...slog.Attr) {
	l.Logger.LogAttrs(ctx, slog.LevelDebug, msg, args...)
}

// InfoAttrs логирует сообщение уровня info с атрибутами
func (l *Logger) InfoAttrs(ctx context.Context, msg string, args ...slog.Attr) {
	l.Logger.LogAttrs(ctx, slog.LevelInfo, msg, args...)
}

// WarnAttrs логирует сообщение уровня warning с атрибутами
func (l *Logger) WarnAttrs(ctx context.Context, msg string, args ...slog.Attr) {
	l.Logger.LogAttrs(ctx, slog.LevelWarn, msg, args...)
}

// ErrorAttrs логирует сообщение уровня error с атрибутами
func (l *Logger) ErrorAttrs(ctx context.Context, msg string, args ...slog.Attr) {
	l.Logger.LogAttrs(ctx, slog.LevelError, msg, args...)
}
