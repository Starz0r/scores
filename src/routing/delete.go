package routing

import (
	"net/http"
	"strconv"

	"github.com/orchestrafm/scores/src/database"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func deleteScore(c echo.Context) error {
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
