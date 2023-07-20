package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/willian2s/web-service-gin/api/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	routes.Execute()
}
