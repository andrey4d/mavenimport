/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package logger

import (
	"log/slog"
)

type Logger interface {
	Debug(string, ...any)
	Info(string, ...any)
	Error(string, ...any)
	Fatal(string, ...any)
}

func NewLogger(h slog.Handler) *slog.Logger {
	l := slog.New(h)
	return l
}
