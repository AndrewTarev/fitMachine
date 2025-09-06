package logger

import (
	"context"
	"log/slog"
)

// ILogger определяет интерфейс логгера
type ILogger interface {
	Debug(ctx context.Context, msg string, args ...any)
	Info(ctx context.Context, msg string, args ...any)
	Warn(ctx context.Context, msg string, args ...any)
	Error(ctx context.Context, msg string, args ...any)
	DebugAttrs(ctx context.Context, msg string, args ...slog.Attr)
	InfoAttrs(ctx context.Context, msg string, args ...slog.Attr)
	WarnAttrs(ctx context.Context, msg string, args ...slog.Attr)
	ErrorAttrs(ctx context.Context, msg string, args ...slog.Attr)
}
