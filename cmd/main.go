/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */

package main

import (
	"log/slog"
	"os"

	"github.com/andrey4d/mavenimport/internal/artifacts"
	"github.com/andrey4d/mavenimport/internal/config"
	"github.com/andrey4d/mavenimport/internal/logger/handlers/slogpretty"
	"github.com/andrey4d/mavenimport/internal/upload"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	log := InitLog(cfg.LogLevel)
	log.Debug("Config", slog.Any("config", cfg))

	log.Info("Run import", slog.String("source", cfg.M2Path+"/"+cfg.ArtifactsPath), slog.String("target", cfg.Url+"/service/rest/repository/browse/"+cfg.Repository))

	arts := artifacts.NewArtifacts(*log, cfg.M2Path, cfg.ArtifactsPath)

	a, err := arts.GetArtifacts()
	if err != nil {
		slog.Error("get artifacts", slog.Any("error", err))
	}

	client := upload.NewClient(*log, cfg.Url, cfg.Repository, cfg.Token)
	for _, v := range a {
		if err := client.Upload(v); err != nil {
			slog.Error("upload", slog.Any("error", err))
		}
	}
}

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
