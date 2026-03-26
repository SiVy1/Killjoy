package core

import "time"

type TrackID string

type Track struct {
	ID          TrackID
	Source      string
	Path        string
	Album       string
	Artist      string
	Title       string
	DurationSec time.Duration
}
