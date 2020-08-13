package routing

import (
	"net/http"

	"github.com/orchestrafm/scores/src/database"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func createScore(c echo.Context) error {
	if authorized := HasRole(c, "create-score"); authorized != true {
		logger.Info().
			Msg("user intent to submit a score, but was unauthorized.")

		return c.JSON(http.StatusUnauthorized, &struct {
			Message string
		}{
			Message: ErrPermissions.Error()})
	}
	claims := SelfAuthCheck(c) //TODO: log a wtf is this is nil
	err, pf := database.SelectProfileByUUID(claims.Subject)
	if err != nil {
		logger.Warn().
			Err(err).
			Msg("user profile could not be found")

		return c.JSON(http.StatusUnauthorized, &struct {
			Message string
		}{
			Message: ErrPermissions.Error()})
	}

	// Data Binding

	s := new(database.Score)
	s.ProfileID = pf.ID
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
	// Get Related Board

	b, err := database.BoardGetFromIDs(s.TrackID, s.BoardID)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Score submission request missing track and board IDs")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Score data was invalid or malformed."})
	}

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
