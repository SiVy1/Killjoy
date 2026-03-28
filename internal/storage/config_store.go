package storage

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/SiVy1/Killjoy/internal/core"
)

func LoadConfig(path string) (core.Config, error) {
	cfg := core.DefaultConfig()

	f, err := os.Open(path)

	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return cfg, err
	}

	defer f.Close()

	dec := json.NewDecoder(f)
	if err := dec.Decode(&cfg); err != nil {
		if errors.Is(err, io.EOF) {
			return cfg, nil
		}
		return cfg, err
	}

	return cfg, nil
}

func SaveConfig(path string, config core.Config) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(config)
}
