package database

import (
	"database/sql"

	"github.com/spidernest-go/logger"
)

func SelectID(id uint64) (*Score, error) {
	scores := db.Collection("scores")
	rs := scores.Find(id)
	s := *new(Score)
	err := rs.One(&s)
	if err != nil && err != sql.ErrNoRows {
		logger.Error().
			Err(err).
			Msg("Bad parameters or database error.")
	}

	if err == sql.ErrNoRows {
		return nil, err
	} else {
		return &s, nil
	}
}
