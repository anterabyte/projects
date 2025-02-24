package config

import "os"

type Config struct {

	Openaikey string

}

func LoadConfig() *Config {

	return &Config{

		Openaikey: os.Getenv("openai_api_key"),
	}

}
