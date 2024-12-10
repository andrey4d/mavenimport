/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package logger

import (
	"log/slog"
	"os"

	"github.com/andrey4d/mavenimport/internal/logger/handlers/slogpretty"
)

func InitLog(loglvl string) *slog.Logger {
	var logger *slog.Logger
	switch loglvl {
	case "info":
		logger = stdSLog(slog.LevelInfo)
	case "debug":
		logger = stdSLog(slog.LevelDebug)
	case "worn":
		logger = stdSLog(slog.LevelWarn)
	case "pinfo":
		logger = prettySlog(slog.LevelInfo)
	case "pdebug":
		logger = prettySlog(slog.LevelDebug)
	case "pworn":
		logger = prettySlog(slog.LevelWarn)
	}
	return logger
}

func prettySlog(lvl slog.Level) *slog.Logger {
	logHandlerOptions := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: lvl,
		},
	}
	return slog.New(logHandlerOptions.NewPrettyHandler(os.Stdout))
}

func stdSLog(lvl slog.Level) *slog.Logger {
	logHandlerOptions := slog.HandlerOptions{
		Level: lvl,
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, &logHandlerOptions))
}
