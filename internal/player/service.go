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

func (s *Service) PlayTrack(track core.Track) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.backend.Load(track.Path); err != nil {
		return err
	}

	if err := s.backend.Play(); err != nil {
		return err
	}

	s.state.Status = StatusPlaying
	s.state.Position = 0
	s.state.Duration = track.DurationSec
	s.current = track
	s.state.TrackID = string(s.current.ID)

	return nil
}

func Pause() error {
	return nil
}

func Resume() error

func Next() error

func Prev() error

func Seek(delta time.Duration) error
