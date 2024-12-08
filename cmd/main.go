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
	"github.com/andrey4d/mavenimport/internal/upload"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	log := InitLog(cfg.LogLevel)
	log.Debug("Config", slog.Any("config", cfg))

	log.Info("Run import...")

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
	var lvl slog.Level
	switch loglvl {
	case "info":
		lvl = slog.LevelInfo
	case "debug":
		lvl = slog.LevelDebug
	case "worn":
		lvl = slog.LevelWarn
	}

	logHandlerOptions := slog.HandlerOptions{
		Level: lvl,
	}

	logHandler := slog.NewJSONHandler(os.Stdout, &logHandlerOptions)
	return slog.New(logHandler)
}
