package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	EthNodeUrl           string
	WavesNodeUrl         string
	ChainId              string
	ScriptFile           string
	DistributorWavesSeed string
	Admins               []string
	BftCoefficient       int64
	RqTimeout            int64
	EthereumPrivKey      string
}

func Load(filename string) (Config, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}
	config := Config{}
	if err := json.Unmarshal(file, &config); err != nil {
		return Config{}, err
	}
	return config, err
}
