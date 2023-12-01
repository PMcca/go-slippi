package slippi

import (
	"github.com/PMcca/go-slippi/internal/sentinel"
	"github.com/toitware/ubjson"
	"os"
)

// ParseGame reads the .slp file given by filePath and returns the decoded game.
func ParseGame(filePath string) (Game, error) {
	if filePath == "" {
		return Game{}, ErrEmptyFilePath
	}

	b, err := os.ReadFile(filePath)
	if err != nil {
		return Game{}, sentinel.WithMessagef(err, ErrReadingFile, "filePath: %s", filePath)
	}

	g := Game{}
	if err := ubjson.Unmarshal(b, &g); err != nil {
		return Game{}, sentinel.WithMessagef(err, ErrParsingMeta, "filePath: %s", filePath)
	}

	return g, nil
}

// ParseMeta reads the .slp file given by filePath and returns the decoded metadata fields.
func ParseMeta(filePath string) (Metadata, error) {
	if filePath == "" {
		return Metadata{}, ErrEmptyFilePath
	}

	b, err := os.ReadFile(filePath)
	if err != nil {
		return Metadata{}, sentinel.WithMessagef(err, ErrReadingFile, "filePath: %s", filePath)
	}

	m := metaOnlyGame{}
	if err := ubjson.Unmarshal(b, &m); err != nil {
		return Metadata{}, sentinel.WithMessagef(err, ErrParsingMeta, "filePath: %s", filePath)
	}

	return m.Meta, nil
}
