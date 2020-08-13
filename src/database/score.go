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
	GradeLetter       string    `db:"grade_letter" json:"grade,omitempty"`
	PerformanceRating float32   `db:"performance_rating" json:"performance,omitempty"`
	ScoreAmount       uint32    `db:"score_amount" json:"score,omitempty"`
	MaxCombo          uint64    `db:"max_combo" json:"combo,omitempty"`
	ClearStatus       uint8     `db:"clear_status" json:"status,omitempty"`
	EffectiveRate     float32   `db:"effective_rate" json:"rate,omitempty"`
	Accuracy          float64   `db:"accuracy" json:"accuracy,omitempty"`
	Criticals         uint32    `db:"criticals" json:"criticals,omitempty"`
	Nears             uint32    `db:"nears" json:"nears,omitempty"`
	Errors            uint32    `db:"errors" json:"errors,omitempty"`
	Modifiers         uint32    `db:"modifiers" json:"mods,omitempty"`
	Experience        uint16    `json:"xp,omitempty"`
	ReplayData        string    `db:"replay_data" json:"replay,omitempty"`
}
