package main

import (
	"go-clean-arch/cmd/server/routes"
	"go-clean-arch/config"
	"go-clean-arch/pkg/postgres"
	"log"
	"os"
)

func main() {
	configName := config.GetConfigByName(os.Getenv("config"))

	config, err := config.GetConfig(configName)

	if err != nil {
		log.Fatalf("Config failed %s", err)
	}

	postgresDB, err := postgres.InitPostgres(config)
	if err != nil {
		log.Fatalf("Postgres failed %s", err)
	}
	defer postgresDB.Close()

	log.Printf("Postgres connected: %#v", postgresDB.Stats())

	app := routes.MakeRouters(postgresDB)
	app.Run(config.Server.Port)

	log.Println("Server listening")

}
