package routing

import (
	"net/http"

	"github.com/orchestrafm/scores/src/database"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func createScore(c echo.Context) error {
	if err := FullAuthCheck(c); err != nil {
		logger.Info().
			Msg("user intent to create a new score, but was unauthorized.")

		return err
	}

	s := new(database.Score)
	if err := c.Bind(s); err != nil {
		logger.Error().
			Err(err).
			Msg("Invalid or malformed score data.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Score data was invalid or malformed."})
	}

	err := s.New()
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Score data did not get submitted to the database."})
	}

	return c.JSON(http.StatusOK, &struct {
		Message string
	}{
		Message: "OK."})
}
