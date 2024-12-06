/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package artifacts

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

type Artifact struct {
	Pom     string
	Package string
	Path    string
}

func GetArtifacts(repository, dir string, packaging string) ([]Artifact, error) {
	out := []Artifact{}
	os.Chdir(repository)
	err := filepath.Walk(dir, visit(&out, packaging))
	if err != nil {
		fmt.Println(err)
	}
	return out, nil
}

func visit(a *[]Artifact, packaging string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {

		if err != nil {
			slog.Error("walk func", slog.Any("error", err))
			return err
		}

		if !info.IsDir() {
			if filepath.Ext(path) == "."+packaging {
				art, err := constructArtifact(path)
				if err != nil {
					return err
				}
				*a = append(*a, *art)
			}

		}
		return nil
	}
}

func constructArtifact(path string) (*Artifact, error) {
	art := Artifact{
		Path:    filepath.Dir(path),
		Package: filepath.Base(path),
	}
	pom := strings.TrimSuffix(art.Package, filepath.Ext(art.Package)) + ".pom"
	if _, err := os.Stat(art.Path + "/" + pom); errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	art.Pom = pom
	return &art, nil
}
