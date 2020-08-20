package helpers

import (
    "log"
	"os"
	"encoding/json"
)

type Config struct {
    Endpoint string `json:"endpoint"`
    OraclePK [5]string `json:"oraclepk"`
}

type NebulaData struct {
    Abi string
    Address string
    SubscriptionId string
    MockAddress string
}

func LoadConfiguration() (Config, NebulaData) {
	var config Config
	var nebulaData NebulaData

	configFile, err := os.Open("config.json")
    defer configFile.Close()
    if err != nil {
        log.Fatal(err)
    }
    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)

    nebulaDataFile, err := os.Open("nebula.json")
    defer nebulaDataFile.Close()
    if err != nil {
        log.Fatal(err)
    }
    jsonParser2 := json.NewDecoder(nebulaDataFile)
	jsonParser2.Decode(&nebulaData)
	return config, nebulaData
}

