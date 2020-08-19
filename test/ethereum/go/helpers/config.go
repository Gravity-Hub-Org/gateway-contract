package helpers

import (
	"encoding/json"
	"os"
)

type Config struct {
	Endpoint string    `json:"endpoint"`
	OraclePK [5]string `json:"oraclepk"`
}

type NebulaData struct {
	Abi     string
	Address string
}

func LoadConfiguration() (Config, NebulaData, error) {
	var config Config
	var nebulaData NebulaData

	configFile, err := os.Open("config.json")
	defer configFile.Close()

	if err != nil {
		return Config{}, NebulaData{}, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return Config{}, NebulaData{}, err
	}

	nebulaDataFile, err := os.Open("nebula.json")
	defer nebulaDataFile.Close()

	if err != nil {
		return Config{}, NebulaData{}, err
	}
	jsonParser2 := json.NewDecoder(nebulaDataFile)
	err = jsonParser2.Decode(&nebulaData)
	if err != nil {
		return Config{}, NebulaData{}, err
	}
	return config, nebulaData, nil
}
