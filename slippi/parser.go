package slippi

import (
	"github.com/PMcca/go-slippi/internal/sentinel"
	"github.com/toitware/ubjson"
	"os"
)

// ParseGame reads the .slp file given by filePath and returns the decoded game.
func ParseGame(filePath string) (Game, error) {
	b, err := readFile(filePath)
	if err != nil {
		return Game{}, err
	}

	g := Game{}
	if err := ubjson.Unmarshal(b, &g); err != nil {
		return Game{}, sentinel.WithMessagef(err, ErrParsingGame, "filePath: %s", filePath)
	}

	return g, nil
}

// ParseMeta reads the .slp file given by filePath and returns the decoded metadata fields.
func ParseMeta(filePath string) (Metadata, error) {
	b, err := readFile(filePath)
	if err != nil {
		return Metadata{}, err
	}

	m := metaOnlyGame{}
	if err := ubjson.Unmarshal(b, &m); err != nil {
		return Metadata{}, sentinel.WithMessagef(err, ErrParsingMeta, "filePath: %s", filePath)
	}

	return m.Meta, nil
}

// readFile reads & returns the bytes of the given .slp file.
func readFile(filePath string) ([]byte, error) {
	if filePath == "" {
		return nil, ErrEmptyFilePath
	}

	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, sentinel.WithMessagef(err, ErrReadingFile, "filePath: %s", filePath)
	}

	return b, nil
}
