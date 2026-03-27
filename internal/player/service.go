package player

import (
	"sync"
	"time"

	"github.com/SiVy1/Killjoy/internal/core"
)

type Service struct {
	backend    Backend
	queue      core.Queue
	tracksByID map[core.TrackID]core.Track
	state      State
	current    core.Track
	mu         sync.RWMutex
}

func PlayTrack(track core.Track) error {
	return nil
}

func Pause() error {
	return nil
}

func Resume() error

func Next() error

func Prev() error

func Seek(delta time.Duration) error
