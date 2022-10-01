package testutil

import (
	"io"
	"os"
	"path/filepath"
)

// ReadUBJSONFile takes a filename and returns the bytes read from the given text fixture.
func ReadUBJSONFile(name string) ([]byte, error) {
	f, err := os.Open(filepath.Join("testdata", name))
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return b, nil
}
