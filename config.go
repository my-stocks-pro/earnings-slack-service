package main

import (
	"log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type TypeConfig struct {
	Token    string
	Channel  string
	Host     string
	Hostprod string
	Port     string
}

func (s *TypeSlackService) LoadConfig() *TypeConfig {

	conf := &TypeConfig{}

	data, errReadFile := ioutil.ReadFile("config/approved-service.yaml")
	if errReadFile != nil {
		log.Fatalf("error: %v", errReadFile)
	}

	errYaml := yaml.Unmarshal(data, &conf)
	if errYaml != nil {
		log.Fatalf("error: %v", errYaml)
	}

	return conf

}

func (s *TypeSlackService) GetSevicePath() (string, string) {
	prod := os.Getenv("PROD")
	if prod == "1" {
		return s.Config.Hostprod, s.Config.Port
	}
	return "127.0.0.1", s.Config.Port
}
