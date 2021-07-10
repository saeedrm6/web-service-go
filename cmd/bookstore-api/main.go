package main

import (
	"github.com/saeedrm6/web-service-go/internal/application"
	"github.com/saeedrm6/web-service-go/internal/config"
	"log"
)

var cfg = &config.Config{}

func init() {

	if err := config.Parse("./build/config/config.yaml",cfg); err != nil {
		log.Fatalln(err)
	}

	if err := config.ReadEnv(cfg); err != nil {
		log.Fatalln(err)
	}

	config.SetConfig(cfg)
}

func main() {
	if err := application.Run(cfg); err != nil {
		log.Fatalln(err)
	}
}
