package routers

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

func getScoresFromBoard(c echo.Context) error {
	t, err := strconv.ParseUint(c.QueryParam("track"), 10, 64)
	if err != nil {
		logger.Error().
			Err(err).
			Msgf("Passed track parameter (%s) was not a valid number", c.QueryParam("track"))

		return c.JSON(http.StatusBadRequest, nil)
	}
	b, err := strconv.ParseUint(c.QueryParam("board"), 10, 64)
	if err != nil {
		logger.Error().
			Err(err).
			Msgf("Passed board parameter (%s) was not a valid number", c.QueryParam("board"))

		return c.JSON(http.StatusBadRequest, nil)
	}
	l, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		logger.Error().
			Err(err).
			Msgf("Passed limit parameter (%s) was not a valid number", c.QueryParam("limit"))

		return c.JSON(http.StatusBadRequest, nil)
	}
	o, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		logger.Error().
			Err(err).
			Msgf("Passed offset parameter (%s) was not a valid number", c.QueryParam("offset"))

		return c.JSON(http.StatusBadRequest, nil)
	}

	ss, err := database.SelectOrdinally(t, b, l, o)
	if err != nil {
		logger.Error().
			Err(err).
			Msg(".")

		c.JSON(http.StatusNotFound, ss)
	}

	return c.JSON(http.StatusOK, ss)
}
