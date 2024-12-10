/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */

package main

import (
	"log/slog"

	"sync"

	"github.com/andrey4d/mavenimport/internal/artifacts"
	"github.com/andrey4d/mavenimport/internal/config"
	"github.com/andrey4d/mavenimport/internal/logger"

	"github.com/andrey4d/mavenimport/internal/upload"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	log := logger.InitLog(cfg.LogLevel)
	log.Debug("Config", slog.Any("config", cfg))

	log.Info("Run import", slog.String("source", cfg.M2Path+"/"+cfg.ArtifactsPath), slog.String("target", cfg.Url+"/service/rest/repository/browse/"+cfg.Repository))

	arts := artifacts.NewArtifacts(*log, cfg.M2Path, cfg.ArtifactsPath)

	a, err := arts.GetArtifacts()
	if err != nil {
		slog.Error("main() get artifacts", slog.Any("error", err))
	}

	client := upload.NewClient(*log, cfg.Url, cfg.Repository, cfg.Token)

	errs := make(chan error)

	go func() {
		for err := range errs {
			slog.Error("main() Upload", slog.Any("error", err))
		}
	}()

	var wg sync.WaitGroup
	for i, v := range a {
		wg.Add(1)
		go client.UploadGoWG(v, &wg, errs, i)
	}
	wg.Wait()
}
