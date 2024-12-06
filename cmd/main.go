/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */

package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/andrey4d/mavenimport/internal/artifacts"
	"github.com/andrey4d/mavenimport/internal/config"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Error:", slog.Any("error", err))
	}

	logHandlerOptions := slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	logHandler := slog.Default().Handler()
	if cfg.LogFormat == "json" {
		logHandler = slog.NewJSONHandler(os.Stdout, &logHandlerOptions)
	} else {
		logHandler = slog.NewTextHandler(os.Stdout, &logHandlerOptions)
	}

	logger := slog.New(logHandler)
	slog.SetDefault(logger)
	slog.Debug("Config", slog.Any("config", cfg))

	slog.Info("Run import...")

	a, err := artifacts.GetArtifacts(cfg.M2Path, cfg.ArtifactsPath, "jar")
	if err != nil {
		slog.Error("get artifacts", slog.Any("error", err))
	}

	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Println(cfg.GetToken())
}
