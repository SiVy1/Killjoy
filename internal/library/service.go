package library

import (
	"strings"

	"github.com/SiVy1/Killjoy/internal/core"
)

type Service struct {
	tracks []core.Track
	byID   map[core.TrackID]core.Track
}

func NewService() *Service {
	return &Service{
		tracks: []core.Track{},
		byID:   map[core.TrackID]core.Track{},
	}
}

func (s *Service) ScanFolders(folders []string) error {
	tracks, err := scanFolders(folders)

	if err != nil {
		return err
	}

	s.tracks = tracks
	s.byID = make(map[core.TrackID]core.Track, len(tracks))

	for _, t := range tracks {
		s.byID[t.ID] = t
	}
	return nil
}

func (s *Service) AllTracks() []core.Track {
	out := make([]core.Track, len(s.tracks))
	copy(out, s.tracks)
	return out
}

func (s *Service) FindById(id core.TrackID) (core.Track, bool) {
	t, ok := s.byID[id]
	return t, ok
}

func (s *Service) Search(query string) []core.Track {
	q := strings.ToLower(strings.TrimSpace(query))

	if q == "" {
		return s.AllTracks()
	}

	var out []core.Track
	for _, t := range s.tracks {
		hay := strings.ToLower(t.Artist + " " + t.Title + " " + t.Album)
		if strings.Contains(hay, q) {
			out = append(out, t)
		}
	}

	return out
}
