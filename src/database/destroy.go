package database

import (
	"github.com/spidernest-go/logger"
)

func Remove(id, tid, bid, pid uint64) error {
	err := db.Collection("boards").
		Find(id).
		Where("track_id = ", tid).
		And("board_id = ", bid).
		And("profile_id = ", pid).
		Delete()
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Entry did not exist or could not be deleted.")
	}
	return err
}
