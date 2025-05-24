package main

import (
	"flag"
	"log"
	"polling/api/handler/http"
	"polling/app"
	"polling/config"
)

func main() {
	var path string
	flag.StringVar(&path, "config_path", "./cmd/config.yaml", "path to config file")
	flag.Parse()

	cfg := config.MustReadConfig[config.PollingConfig](path)
	log.Printf("LoadedConfig: %v\n", cfg)

	appContainer, err := app.NewApp(cfg)
	if err != nil {
		log.Fatalf("can not create polling app: %v", err)
	}

	log.Fatal(http.Run(appContainer, cfg.Polling))
}
