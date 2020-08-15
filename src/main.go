package main

import (
	"github.com/orchestrafm/scores/src/database"
	"github.com/orchestrafm/scores/src/routers"
	"github.com/spidernest-go/logger"
)

func main() {
	err := database.Connect()
	logger.Error().
		Err(err).
		Msg("MySQL Database could not be attached to.")
	database.Synchronize()

	routers.ListenAndServe()
}
