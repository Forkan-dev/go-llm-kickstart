package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server     ServerConfig     `yaml:"server"`
	Database   DatabaseConfig   `yaml:"database"`
	Validation ValidationConfig `yaml:"validation"`
	Token      TokenConfig      `yaml:"token"`
}

type ValidationConfig struct {
	Password PasswordValidationConfig `yaml:"password"`
}

type PasswordValidationConfig struct {
	MinLength int `yaml:"min_length"`
	MaxLength int `yaml:"max_length"`
}

type ServerConfig struct {
	Port      int    `yaml:"port"`
	JWTSecret string `yaml:"jwt_secret"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type TokenConfig struct {
	Issuer                 string `yaml:"issuer"`
	AccessTokenExpiration  int    `yaml:"access_token_expiration"`
	RefreshTokenExpiration int    `yaml:"refresh_token_expiration"`
}

var cfg *Config

func Load() (*Config, error) {
	data, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		return nil, err
	}

	cfg = &Config{}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func Get() *Config {
	if cfg == nil {
		panic("config not loaded")
	}
	return cfg
}
