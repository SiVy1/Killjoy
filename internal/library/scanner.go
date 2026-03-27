package library

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/SiVy1/Killjoy/internal/core"
)

var supportedExt = map[string]struct{}{
	".mp3":  {},
	".flac": {},
	".wav":  {},
	".ogg":  {},
	".m4a":  {},
}

func scanFolders(folders []string) ([]core.Track, error) {
	var out []core.Track

	for _, root := range folders {
		err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}

			ext := strings.ToLower(filepath.Ext(path))
			if _, ok := supportedExt[ext]; !ok {
				return nil
			}

			title := strings.TrimSuffix(filepath.Base(path), ext)
			id := core.TrackID(path)

			out = append(out, core.Track{
				ID:     id,
				Source: core.SourceLocal,
				Path:   path,
				Title:  title,
				Artist: "Unknown",
			})

			return nil
		})

		if err != nil {
			return nil, err
		}
	}
	return out, nil
}
