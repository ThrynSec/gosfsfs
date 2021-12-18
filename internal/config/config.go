package config

import (
	"encoding/json"
	"log"
	"os"
)

var (
	Path     string
	FTP      FTPConfig
	Settings BackupConfig
)

type FTPConfig struct {
	Adress   string `json:"ftp_server"`
	Port     int    `json:"ftp_port"`
	User     string `json:"ftp_user"`
	Password string `json:"ftp_password"`
}

type BackupConfig struct {
	Max       int    `json:"max_backup"`
	Unit      string `json:"frequency_unit"`
	Frequency int    `json:"frequency_occurence"`
}

func LoadConfig() {
	var err error

	Path, err = os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	var genConfig struct {
		FTPConf  FTPConfig    `json:"ftp"`
		SaveConf BackupConfig `json:"save"`
	}

	genJSON, err := os.ReadFile(Path + "\\internal\\config\\config.json")
	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal([]byte(genJSON), &genConfig)
	if err != nil {
		log.Panic(err)
	}

	FTP = genConfig.FTPConf
	Settings = genConfig.SaveConf

}
