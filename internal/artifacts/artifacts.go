/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package artifacts

import (
	"errors"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

type Artifact struct {
	Pom     string
	Package string
}

type Artifacts struct {
	log        slog.Logger
	arts       []Artifact
	repository string
	dir        string
}

const PACKAGING = "jar"

func NewArtifacts(logger slog.Logger, repository, dir string) *Artifacts {
	return &Artifacts{
		log:        logger,
		repository: repository,
		dir:        dir,
	}
}

func (a *Artifacts) GetArtifacts() ([]Artifact, error) {

	os.Chdir(a.repository)
	err := filepath.Walk(a.dir, a.walk())
	if err != nil {
		a.log.Error("GetArtifacts()", slog.Any("error", err))
	}
	return a.arts, nil
}

func (a *Artifacts) walk() filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {

		if err != nil {
			a.log.WithGroup("walk()").Error("GetArtifacts()", slog.Any("error", err))
			return err
		}

		if !info.IsDir() {
			if filepath.Ext(path) == "."+PACKAGING {
				art, err := a.constructArtifact(path)
				if err != nil {
					a.log.WithGroup("constructArtifact()").Error("GetArtifacts()", slog.Any("skip", err))
				} else {
					a.arts = append(a.arts, *art)
				}

			}

		}
		return nil
	}
}

func (a *Artifacts) constructArtifact(path string) (*Artifact, error) {
	art := Artifact{
		Package: a.repository + "/" + path,
	}
	pom := strings.TrimSuffix(art.Package, filepath.Ext(art.Package)) + ".pom"

	if _, err := os.Stat(pom); errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	art.Pom = pom
	return &art, nil
}
