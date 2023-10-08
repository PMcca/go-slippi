package slippi

import (
	"github.com/PMcca/go-slippi/internal/sentinel"
	"github.com/toitware/ubjson"
	"os"
)

// ParseMeta reads the .slp file given by filePath and returns the decoded metadata fields.
func ParseMeta(filePath string) (Metadata, error) {
	if filePath == "" {
		return Metadata{}, ErrEmptyFilePath
	}

	m := metaOnlyGame{}

	b, err := os.ReadFile(filePath)
	if err != nil {
		return Metadata{}, sentinel.WithMessagef(err, ErrReadingFile, "filePath: %s", filePath)
	}

	if err := ubjson.Unmarshal(b, &m); err != nil {
		return Metadata{}, sentinel.WithMessagef(err, ErrParsingMeta, "filePath: %s", filePath)
	}

	return m.Meta, nil
}
