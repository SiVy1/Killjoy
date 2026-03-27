package player

import "time"

type Status string

const (
	StatusStopped Status = "stopped"
	StatusPlaying Status = "playing"
	StatusPaused  Status = "paused"
)

type State struct {
	TrackID  string
	Status   Status
	Position time.Duration
	Duration time.Duration
	Volume   int
}
