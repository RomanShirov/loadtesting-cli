package main

import (
	"github.com/RomanShirov/loadtesting-cli/internal/cli"
	"github.com/RomanShirov/loadtesting-cli/internal/config"
	"github.com/RomanShirov/loadtesting-cli/internal/loadtest"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	loadConfigMode, err := cli.SelectLoadConfigMode()
	if err != nil {
		log.Fatalf("CLI Reading error: %v", err)
	}

	switch loadConfigMode {
	case "ENV":
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
		testConfig := config.GetConfigFromEnvironment()
		loadtest.RunLoadTest(testConfig)
	case "MANUAL":
		testConfig := cli.GatherConfigData()
		loadtest.RunLoadTest(testConfig)
	}
}
