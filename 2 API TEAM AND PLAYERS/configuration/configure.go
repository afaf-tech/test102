package configuration

import (
	"encoding/json"
	"os"
)

type Config struct {
	Database struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}

func LoadConfiguration() (Config, error) {
	var config Config
	basePath, err := os.Getwd()
	configFile, err := os.Open(basePath + "/configuration/config.json")
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err

}
