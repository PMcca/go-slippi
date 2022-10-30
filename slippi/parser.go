package slippi

import (
	"github.com/PMcca/go-slippi/internal/sentinel"
	"github.com/jmank88/ubjson"
	"os"
)

// ParseMeta reads the .slp file given by filePath and returns the decoded metadata fields.
func ParseMeta(filePath string) (*Metadata, error) {
	if filePath == "" {
		return nil, ErrEmptyFilePath
	}

	m := metaOnlyGame{}

	f, err := os.Open(filePath)
	if err != nil {
		return nil, sentinel.WithMessagef(err, ErrOpeningFile, "filePath: %s", filePath)
	}
	defer f.Close()

	dec := ubjson.NewDecoder(f)
	if err := dec.Decode(&m); err != nil {
		return nil, sentinel.WithMessagef(err, ErrParsingMeta, "filePath: %s", filePath)
	}

	return &m.Meta, nil
}
