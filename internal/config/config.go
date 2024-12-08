/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package config

import (
	"encoding/base64"
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

const CONFIG_PATH = "./configs/config.yaml"

type Config struct {
	LogLevel      string `yaml:"log_level"`
	ArtifactsPath string `yaml:"artifacts_path"`
	M2Path        string `yaml:"m2_path"`
	Token         string `yaml:"token"`
	Url           string `yaml:"repository_url"`
	Repository    string `yaml:"repository_name"`
}

func NewConfig() *Config {
	return &Config{}
}

func LoadConfig() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		slog.Warn("CONFIG_PATH is not set. Use default", slog.String("path", CONFIG_PATH))
		configPath = CONFIG_PATH
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		slog.Error("Config file doesn't exist.", slog.String("path", configPath))
		return nil, err
	}

	cfg_data, err := os.ReadFile(configPath)

	if err != nil {
		slog.Error("Read file", slog.String("path", configPath))
		return nil, err
	}

	config := NewConfig()
	yaml.Unmarshal(cfg_data, config)

	return config, nil
}

func (c *Config) GetConfig() *Config {
	return c
}

func (c *Config) GetToken() (string, error) {
	data, err := base64.StdEncoding.DecodeString(c.Token)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
