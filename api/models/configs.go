package models

type Config struct {
	JAGW struct {
		Server             string `yaml:"server"`
		RequestServicePort uint16 `yaml:"requestServicePort"`
	} `yaml:"JAGW"`
}
