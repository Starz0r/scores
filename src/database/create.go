package database

import (
	"database/sql"
	"errors"
	"time"

	"github.com/spidernest-go/logger"
)

var ErrUnderperformed = errors.New("score performed worst than previously submitted")

func (s *Score) New() error {
	// check for a previous score by the same profile
	scores := db.Collection("scores")
	prevscore := *new(Score)
	err := scores.Find().
		Where("track_id = ", s.TrackID).
		And("board_id = ", s.BoardID).
		And("profile_id = ", s.ProfileID).
		One(&prevscore)

		// update the score if the performance rating is higher
	if err == sql.ErrNoRows {
		if s.PerformanceRating > prevscore.PerformanceRating {
			s.ID = prevscore.ID
			stmt, err := db.Update("scores").
				Set(s).
				Prepare()
			if err != nil {
				logger.Error().
					Err(err).
					Msg("SQL Statement on score update couldn't be prepared.")

				return err
			}

			_, err = stmt.Exec()
			if err != nil {
				logger.Error().
					Err(err).
					Msg("Score entry could not be updated.")
			}

			return err
		} else {
			return ErrUnderperformed
		}
	} else if err != nil && err != sql.ErrNoRows {
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
