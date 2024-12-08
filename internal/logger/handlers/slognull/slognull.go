package slognull

import (
	"context"
	"log/slog"
)

type NullHandler struct{}

func New() *slog.Logger {
	return slog.New(NewNullHandler())
}

func NewNullHandler() *NullHandler {
	return &NullHandler{}
}

func (h *NullHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

func (h *NullHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *NullHandler) WithGroup(_ string) slog.Handler {
	return h
}
func (h *NullHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}
