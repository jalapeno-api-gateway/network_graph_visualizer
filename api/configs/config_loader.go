package configs

import (
	"log"
	"network_graph_api/models"
	"os"

	"gopkg.in/yaml.v3"
)

func readConfig() (models.Config, error) {
	var config models.Config

	yamlBytes, err := os.ReadFile("configs/config.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
		return models.Config{}, err
	}

	err = yaml.Unmarshal(yamlBytes, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
		return models.Config{}, err
	}
	return config, nil
}

func LoadConfig() models.Config {
	config, _ := readConfig()
	return config
}
