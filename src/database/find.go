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

func SelectOrdinally(track, board uint64, limit int, offset int) ([]*Score, error) {
	if limit > 100 {
		// TODO: return an actual error
		return nil, nil
	}
	var ss []*Score
	scores := db.Collection("scores")
	// TODO: use the pagination api instead of limit and offset
	rs := scores.Find().
		Where("track_id = ", track).
		And("board_id = ", board).
		OrderBy("performance_rating DESC").
		Limit(limit).
		Offset(offset)

	err := rs.All(&ss)
	if err != nil && err != sql.ErrNoRows {
		logger.Error().
			Err(err).
			Msg("Bad parameters or database error.")
	}

	if err == sql.ErrNoRows {
		return nil, err
	} else {
		return ss, nil
	}
}

func SelectProfileByUUID(uuid string) (error, *Profile) {
	pf := *new(Profile)
	err := db.SelectFrom("profiles").
		Where("uuid = " + uuid).
		Limit(1).
		One(&pf)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("SQL Execution had an issue when executing.")
		return err, nil
	}

	// pf.UUID = *new(string) // QUEST: Should I clear this field so it doesn't return?

	return nil, &pf
}
