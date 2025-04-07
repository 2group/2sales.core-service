package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env  string     `yaml:"env" env-default:"local"`
	REST RestConfig `yaml:"rest"`
	GRPC GrpcConfig `yaml:"grpc"`
	Psql PsqlConfig `yaml:"postgres"`
}

type PsqlConfig struct {
	Url string `yaml:"url"`
}

type RestConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type GrpcConfig struct {
	User          string `yaml:"user"`
	Organization  string `yaml:"organization"`
	Product       string `yaml:"product"`
	CRM           string `yaml:"crm"`
	Warehouse     string `yaml:"warehouse"`
	Order         string `yaml:"order"`
	Advertisement string `yaml:"advertisement"`
	Customer      string `yaml:"customer" env-default:"localhost:50058"`
	Service       string `yaml:"service" env-default:"localhost:50059"`
}

func MustLoad() *Config {
	path := fetchConfigPath()

	fmt.Println("Using config path:", path)
	if path == "" {
		panic("config path is empty")
	}

	return MustLoadByPath(path)
}

func MustLoadByPath(path string) *Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exists: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
