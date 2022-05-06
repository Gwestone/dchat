package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	DbAddr        string `yaml:"DbAddr"`
	Port          int    `yaml:"Port"`
	RedisAddr     string `yaml:"RedisAddr"`
	RedisPassword string `yaml:"RedisPassword"`
	JWTSecret     string `yaml:"JWTSecret"`
	UseSSL        bool   `yaml:"UseSSL"`
}

func NewConfig() *Config {
	return &Config{}
}

func Parse(filename string) (*Config, error) {

	yamlBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := NewConfig()
	err = yaml.Unmarshal(yamlBytes, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
