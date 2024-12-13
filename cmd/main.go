/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */

package main

import (
	"github.com/andrey4d/mavenimport/internal/application"
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
	log.Debug("Config", logger.Any("config", cfg))

	for _, dir := range cfg.ArtifactsPath {

		log.Info("main()", logger.String("Run import", cfg.M2Path+"/"+dir), logger.String("target", cfg.Url+"/service/rest/repository/browse/"+cfg.Repository))

		arts := artifacts.NewArtifacts(log, cfg.M2Path, dir)
		a, err := arts.GetArtifacts()
		if err != nil {
			log.Error("main() get artifacts", logger.Any("error", err))
		}

		client := upload.NewClient(log, cfg.Url, cfg.Repository, cfg.Token)
		application := application.NewApplication(log, *client, a)
		application.Run()
	}

}
