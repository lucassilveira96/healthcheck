package api

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lucassilveira96/healthcheck/api/controllers"
)

var server = controllers.Server{}

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	}
	server.Initialize()
	server.Run(":8295")
}
