package routing

import (
	"net/http"
	"strconv"

	"github.com/orchestrafm/scores/src/database"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func getScore(c echo.Context) error {
	i, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error().
			Err(err).
			Msgf("Passed id parameter (%s) was not a valid number", c.Param("id"))

		return c.JSON(http.StatusBadRequest, nil)
	}

	s, err := database.SelectID(i)
	if err != nil {
		logger.Error().
			Err(err).
			Msg(".")

		c.JSON(http.StatusNotFound, ErrGeneric)
	}

	return c.JSON(http.StatusOK, s)
}