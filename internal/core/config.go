package core

type Config struct {
	MusicFolders   []string `json:"musicFolders"`
	LastPlaylistID string   `json:"lastPlaylistId,omitempty"`
	Volume         int      `json:"volume"`
}

func DefaultConfig() Config {
	return Config{
		MusicFolders: []string{},
		Volume:       80,
	}
}
