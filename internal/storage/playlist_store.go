package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/SiVy1/Killjoy/internal/core"
)

func loadPlaylists(dir string) ([]core.Playlist, error) {
	entries, err := os.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	playlists := make([]core.Playlist, 0, len(entries))

	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		p := filepath.Join(dir, e.Name())
		f, err := os.Open(p)
		if err != nil {
			return nil, err
		}

		var pl core.Playlist

		decErr := json.NewDecoder(f).Decode(&pl)
		closeErr := f.Close()

		if decErr != nil {
			return nil, decErr
		}

		if closeErr != nil {
			return nil, closeErr
		}

		playlists = append(playlists, pl)
	}

	return playlists, nil
}

func savePlaylist(dir string, playlists []core.Playlist) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	for _, pl := range playlists {
		if pl.ID == "" {
			return os.ErrInvalid
		}

		p := filepath.Join(dir, pl.ID+".json")
		f, err := os.Create(p)
		if err != nil {
			return err
		}

		enc := json.NewEncoder(f)
		enc.SetIndent("", "  ")

		encErr := enc.Encode(pl)
		closeErr := f.Close()

		if encErr != nil {
			return encErr
		}
		if closeErr != nil {
			return closeErr
		}
	}

	return nil
}
