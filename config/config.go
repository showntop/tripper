package config

import (
	"encoding/json"
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
)

type configure struct {
	Host          map[string]string      `json:"host"`
	DefaultLocale string                 `json:"default_locale"`
	ServerPort    string                 `json:"server_port"`
	I18nPath      string                 `json:"i18n_path"`
	Env           string                 `json:"env"`
	TmplPath      string                 `json:"tmpl_path"`
	LogPath       string                 `json:"log_path"`
	LogLevel      log.Level              `json:"log_level"`
	Database      map[string]interface{} `json:"database"`
	Redis         map[string]interface{} `json:"redis"`
}

var (
	Config configure
)

func init() {
	confPath := flag.String("c", "./config/conf.json", "Config file")
	flag.Parse()
	//load conf
	log.Debug("log file:" + *confPath)
	configFile, err := os.Open(*confPath)
	if err != nil {
		fmt.Println("opening config file", err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&Config); err != nil {
		fmt.Println("parsing config file", err.Error())
	}
}
