package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	DbAddr string `yaml:"DbAddr"`
	Port   int    `yaml:"Port"`
}

func newConfig() *Config {
	return &Config{}
}

func Parse(filename string) (*Config, error) {

	yamlBytes, err := os.ReadFile(filename)
	//fmt.Printf(string(yamlBytes))
	if err != nil {
		return nil, err
	}

	config := newConfig()
	err = yaml.Unmarshal(yamlBytes, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
