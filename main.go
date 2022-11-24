package main

import (
	"github.com/RomanShirov/loadtesting-cli/internal/cli"
	"log"
)

func main() {
	_, err := cli.SelectLoadConfigMode()
	if err != nil {
		log.Fatalf("CLI Reading error: %v", err)
	}
}
