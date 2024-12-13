/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package logger

import (
	"log/slog"
	"os"
	"time"
)

type Slogger struct {
	Logger *slog.Logger
}

type Handler slog.Handler

const (
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelDebug = slog.LevelDebug
	LevelError = slog.LevelError
)

func NewJSONLogger(lvl slog.Level) *Slogger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: lvl,
	})
	return &Slogger{
		Logger: slog.New(handler),
	}
}

func NewTextLogger(lvl slog.Level) *Slogger {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: lvl,
	})
	return &Slogger{
		Logger: slog.New(handler),
	}
}

func (l *Slogger) Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

func (l *Slogger) Debug(msg string, args ...any) {
	l.Logger.Debug(msg, args...)
}

func (l *Slogger) Warn(msg string, args ...any) {
	l.Logger.Warn(msg, args...)
}

func (l *Slogger) Error(msg string, args ...any) {
	l.Logger.Error(msg, args...)
}

func Any(key string, value any) any {
	return slog.Any(key, value)
}

// Int64 returns an Attr for an int64.
func Int64(key string, value int64) any {
	return slog.Int64(key, value)
}

// Int converts an int to an int64 and returns
// an Attr with that value.
func Int(key string, value int) any {
	return Int64(key, int64(value))
}

// Uint64 returns an Attr for a uint64.
func Uint64(key string, value uint64) any {
	return slog.Uint64(key, value)
}

// Float64 returns an Attr for a floating-point number.
func Float64(key string, value float64) any {
	return slog.Float64(key, value)
}

// Bool returns an Attr for a bool.
func Bool(key string, value bool) any {
	return slog.Bool(key, value)
}

// Time returns an Attr for a [time.Time].
// It discards the monotonic portion.
func Time(key string, value time.Time) any {
	return slog.Time(key, value)
}

// Duration returns an Attr for a [time.Duration].
func Duration(key string, value time.Duration) any {
	return slog.Duration(key, value)
}

// String returns an Attr for a string value.
func String(key, value string) any {
	return slog.String(key, value)
}

// Warn calls [Logger.Warn] on the default logger.
func Warn(msg string, args ...any) {
	slog.Warn(msg, args...)
}

// Error calls [Logger.Error] on the default logger.
func Error(msg string, args ...any) {
	slog.Error(msg, args...)
}

func Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

func Debug(msg string, args ...any) {
	slog.Debug(msg, args...)
}
