package player

import "time"

type Backend interface {
	Load(path string) error
	Play() error
	Pause() error
	Stop() error
	Seek(time.Duration) error
	Position() time.Duration
	Duration() time.Duration
}
