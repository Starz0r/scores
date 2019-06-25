package database

import (
	"time"

	"github.com/spidernest-go/logger"
)

func (b *Score) New() error {
	b.DateCreated = time.Now()
	// TODO: Make sure something doesn't already exist in the spot [id, track_id, board_id, profile_id]
	_, err := db.InsertInto("scores").
		Values(b).
		Exec()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Score could not be inserted into the table.")
	}

	return err
}
