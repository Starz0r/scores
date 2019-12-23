package routing

import (
	"net/http"
	"strconv"

	"github.com/orchestrafm/scores/src/database"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func deleteScore(c echo.Context) error {
	// auth check
	admin, auth := AuthorizationCheck(c)
	if auth != true {
		logger.Info().
			Msg("user intent to create a delete a score, but was unauthorized.")

		return c.JSON(http.StatusUnauthorized, &struct {
			Message string
		}{
			Message: "Insufficient Permissions."})
	}

	// search for score
	i, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Invalid Parameters for deleting a score.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "."})
	}

	// remove score
	if !admin {
		return c.JSON(http.StatusUnauthorized, &struct {
			Message string
		}{
			Message: "Not an Administrator."})
	}
	err = database.Remove(i)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Invalid Parameters for deleting a score.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "."})
	}
	return c.JSON(http.StatusOK, &struct {
		Message string
	}{
		Message: "Score deleted successfully."})
}
