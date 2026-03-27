package library

import (
	"path/filepath"
	"strings"
	"time"
)

type Metadata struct {
	Title    string
	Artist   string
	Album    string
	Duration time.Duration
}

func readMetadata(path string) (Metadata, error) {
	metadata := fallbackMetadata(path)
	return metadata, nil
}

func fallbackMetadata(path string) Metadata {
	t := titleFromFilename(path)
	return Metadata{
		Title:    t,
		Artist:   "Unknown Artist",
		Duration: 0 * time.Second,
	}
}

func titleFromFilename(path string) string {
	f := filepath.Base(path)
	ext := filepath.Ext(path)
	title := strings.TrimSuffix(f, ext)
	return title
}

func normalizeText(value string, fallback string) string {
	n := strings.TrimSpace(value)

	if n == "" {
		return fallback
	}
	return n
}
