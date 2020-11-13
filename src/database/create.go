package database

import (
	"database/sql"
	"errors"
	"time"

	orm "github.com/spidernest-go/db"
	"github.com/spidernest-go/logger"
)

var ErrUnderperformed = errors.New("score performed worst than previously submitted")

func (s *Score) New() error {
	// check for a previous score by the same profile
	scores := db.Collection("scores")
	prevscore := *new(Score)
	res := scores.Find().
		Where("track_id = ", s.TrackID).
		And("board_id = ", s.BoardID).
		And("profile_id = ", s.ProfileID)

	err := res.One(&prevscore)

	// update the score if the performance rating is higher
	if err == nil {
		if s.PerformanceRating > prevscore.PerformanceRating {
			s.ID = prevscore.ID
			err = res.Update(s)
			if err != nil {
				logger.Error().
					Err(err).
					Msg("Score entry could not be updated.")
			}

			return err
		} else {
			return ErrUnderperformed
		}
	} else if err != nil && err != sql.ErrNoRows && err != orm.ErrNoMoreRows {
		logger.Error().
			Err(err).
			Msg("Bad parameters or database error.")

		return err
	}

	s.DateCreated = time.Now()
	_, err = db.InsertInto("scores").
		Values(s).
		Exec()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Score could not be inserted into the table.")
	}

	return err
}
