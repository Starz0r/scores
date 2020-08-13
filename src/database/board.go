package database

import (
	"database/sql"
	"time"

	"github.com/spidernest-go/logger"
)

// HACK: Including this file pretty much "breaks" the contingious structure
// of microservices I had in mind for this project, but currently this is
// no way to get a board object from another microservice without using
// an http request, which I find wasteful. All changes to the board structure
// from the other microservice must also be reflected in here as well.

type Board struct {
	TrackID          uint64    `db:"track_id" json:"track_id,omitempty"`
	ID               uint64    `db:"id" json:"id,omitempty"`
	DateCreated      time.Time `db:"date_created" json:"date_created,omitempty"`
	DateModified     time.Time `db:"date_modified" json:"date_modified,omitempty"`
	SHA3             string    `db:"sha3" json:"sha3,omitempty"`
	Jacket           []byte    `db:"jacket" json:"jacket,omitempty"`
	Charters         string    `db:"charters" json:"charters,omitempty"`
	DifficultyName   uint64    `db:"difficulty_name" json:"difficulty_name"`
	DifficultyRating uint8     `db:"difficulty_rating" json:"difficulty_rating"`
}

func BoardGetFromIDs(track, board uint64) (*Board, error) {
	b := *new(Board)
	boards := db.Collection("boards")
	// TODO: use the pagination api instead of limit and offset
	rs := boards.Find().
		Where("track_id = ", track).
		And("id = ", board)

	err := rs.One(&b)
	if err != nil && err != sql.ErrNoRows {
		logger.Error().
			Err(err).
			Msg("Bad parameters or database error.")
	}

	if err == sql.ErrNoRows {
		return nil, err
	} else {
		return &b, nil
	}
}
