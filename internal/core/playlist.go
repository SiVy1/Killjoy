package core

import "time"

type Playlist struct {
	ID        string
	Name      string
	TrackIDs  []TrackID
	CreatedAt time.Time
	UpdatedAt time.Time
}
