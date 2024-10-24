package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/jonathanschmittcs/go-intrumentation-echo-zerolog/internal/mylogger"
	"github.com/jonathanschmittcs/go-intrumentation-echo-zerolog/internal/myserver"
	"github.com/jonathanschmittcs/go-intrumentation-echo-zerolog/internal/mytracer"
	"github.com/rs/zerolog/log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Error loading .env file")
		os.Exit(1)
	}

	mylogger.Init()

	err = mytracer.Init()
	if err != nil {
		log.Fatal().Timestamp().Err(err).Msg("Error initializing APM Tracer")
		os.Exit(1)
	}

	myserver.Start(":8080")
}
