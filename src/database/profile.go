package database

//HACK: I'd like to get rid of this somehow, and get back to the pureness
// of a microservice application structure
type Profile struct {
	ID                uint64   `db:"id" json:"id"`
	UUID              string   `db:"uuid" json:"uuid,omitempty"`
	Username          string   `json:"name,omitempty"`
	Groups            []string `json:"groups,omitempty"`
	Experience        uint64   `db:"experience" json:"experience"`
	Level             uint64   `db:"level" json:"level"`
	TotalScore        uint64   `db:"total_score" json:"total_score"`
	PlayCount         uint64   `db:"play_count" json:"play_count"`
	Mastery           uint8    `db:"mastery" json:"mastery"`
	PerformanceRating uint64   `db:"performance_rating" json:"performance_rating"`
}