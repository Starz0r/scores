package database

import (
	"time"
)

type Score struct {
	ID                uint64    `db:"id" json:"id,omitempty"`
	TrackID           uint64    `db:"track_id" json:"track,omitempty"`
	BoardID           uint64    `db:"board_id" json:"board,omitempty"`
	ProfileID         uint64    `db:"profile_id" json:"profile,omitempty"`
	DateCreated       time.Time `db:"date_created" json:"date_created,omitempty"`
	GradeLetter       uint8     `db:"grade_letter" json:"grade,omitempty"`
	PerformanceRating uint32    `db:"performance_rating" json:"performance,omitempty"`
	ScoreAmount       uint64    `db:"score_amount" json:"score,omitempty"`
	MaxCombo          uint64    `db:"max_combo" json:"combo,omitempty"`
	ClearStatus       int8      `db:"clear_status" json:"status,omitempty"`
	EffectiveRate     uint8     `db:"effective_rate" json:"rate,omitempty"`
	Accuracy          float64   `db:"accuracy" json:"accuracy,omitempty"`
	Criticals         uint16    `db:"criticals" json:"criticals,omitempty"`
	Nears             uint16    `db:"nears" json:"nears,omitempty"`
	Errors            uint16    `db:"errors" json:"errors,omitempty"`
	Modifiers         string    `db:"modifiers" json:"mods,omitempty"`
	ReplayData        string    `db:"replaydata" json:"replay,omitempty"`
}
