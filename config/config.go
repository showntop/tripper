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
	Env           string                 `json:"-"`
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
	confEnv := flag.String("e", "develop", "Env mode")
	flag.Parse()
	//load conf
	Config.Env = *confEnv
	confPath := flag.String("c", "./config/conf."+*confEnv+".json", "Config file")

	configFile, err := os.Open(*confPath)
	if err != nil {
		fmt.Println("opening config file", err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&Config); err != nil {
		fmt.Println("parsing config file", err.Error())
	}
}
