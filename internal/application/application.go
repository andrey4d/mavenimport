/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package application

import (
	"log/slog"
	"sync"

	"github.com/andrey4d/mavenimport/internal/artifacts"
	"github.com/andrey4d/mavenimport/internal/upload"
)

type Application struct {
	logger    slog.Logger
	client    upload.Client
	artifacts []artifacts.Artifact
}

func NewApplication(logger slog.Logger, client upload.Client, artifacts []artifacts.Artifact) *Application {
	return &Application{
		logger:    logger,
		client:    client,
		artifacts: artifacts,
	}
}

func (a *Application) Run() {
	errs := make(chan error)

	go func() {
		for err := range errs {
			slog.Error("main() Upload", slog.Any("error", err))
		}
	}()

	var wg sync.WaitGroup
	for i, v := range a.artifacts {
		wg.Add(1)
		go a.client.UploadGoWG(v, &wg, errs, i)
	}
	wg.Wait()
}
