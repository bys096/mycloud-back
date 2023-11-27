package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Db struct {
	DbDriver string `json:"dbDriver" yaml:"driver"`
	Dsn      string `json:"dsn" yaml:"dsn"`
}

type Config struct {
	Db Db `json:"db" yaml:"db"`
}

func LoadConfigForYaml() (*Config, error) {
	f, err := os.Open("./config/config.yml")
	if err != nil {
		log.Fatal("loadConfigForYaml os.Open err:", err)
		return nil, err
	}
	defer f.Close()

	cfg := &Config{}
	//err = yaml.NewDecoder(f).Decode(&cfg)

	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}

	return cfg, err
}
