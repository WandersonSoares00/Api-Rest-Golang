package configs

import (
	"encoding/json"
	"os"
)

var (
	cfg    *config
	logger *Logger
)

type config struct {
	ApiPort int    `json:"APIPort"`
	DbHost  string `json:"DBHost"`
	DbPort  int    `json:"DBPort"`
	DbUser  string `json:"DBUser"`
	DbPass  string `json:"DBPass"`
	Dbname  string `json:"DBName"`
	SSL     string `json:"SSLMode"`
}

func init() {
	logger = NewLogger("config")

	cfgFile, err := os.Open("configs/config.json")

	if err != nil {
		logger.Errorf("error while opening config.json file: %s", err.Error())
		return
	}

	defer cfgFile.Close()

	cfg = new(config)

	dec := json.NewDecoder(cfgFile)
	err = dec.Decode(cfg)

	if err != nil {
		logger.Errorf("error while decoding config.json file: %s", err.Error())
		return
	}

	logger.Infof("settings read successfully\n")
}

func GetConfs() *config {
	return cfg
}

func GetServerPort() int {
	if cfg == nil {
		return 0
	}
	return cfg.ApiPort
}

func GetLogger(msg string) *Logger {
	logger = NewLogger(msg)
	return logger
}
