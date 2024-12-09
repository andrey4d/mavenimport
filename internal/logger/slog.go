/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package logger

import (
	"log/slog"
	"os"
)

type SLogger struct {
	Logger *slog.Logger
}

// type Handler slog.Handler

const (
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelDebug = slog.LevelDebug
	LevelError = slog.LevelError
)

func (l *SLogger) Info(msg string, args ...any) {
	l.Logger.Info(msg, args...)
}

func (l *SLogger) Debug(msg string, args ...any) {
	l.Logger.Debug(msg, args...)
}

func (l *SLogger) Warn(msg string, args ...any) {
	l.Logger.Warn(msg, args...)
}

func (l *SLogger) Error(msg string, args ...any) {
	l.Logger.Error(msg, args...)
}

func NewLogger(lvl slog.Level) *SLogger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: lvl,
	})
	return &SLogger{
		Logger: slog.New(handler),
	}
}
