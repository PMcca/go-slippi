package goslippi

import (
	"github.com/PMcca/go-slippi/internal/errutil"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/toitware/ubjson"
)

// ignoreRaw is used with metaOnly game. It implements ubjson.UnmarshalUBJSON to simply return nil and skip over the
// reading of the raw element.
type ignoreRaw struct{}

// metaOnlyGame is an internal type to omit parsing the 'raw' element and only parse the 'metadata' fields.
type metaOnlyGame struct {
	IgnoreRaw ignoreRaw       `ubjson:"raw"`
	Meta      slippi.Metadata `ubjson:"metadata"`
}

// ParseMeta reads the .slp file given by filePath and returns the decoded metadata fields.
func ParseMeta(filePath string) (slippi.Metadata, error) {
	b, err := readFile(filePath)
	if err != nil {
		return slippi.Metadata{}, err
	}

	m := metaOnlyGame{}
	if err := ubjson.Unmarshal(b, &m); err != nil {
		return slippi.Metadata{}, errutil.WithMessagef(err, ErrParsingMeta, "filePath: %s", filePath)
	}

	return m.Meta, nil
}

// UnmarshalUBJSON returns immediately from reading raw, for performance benefits.
func (i ignoreRaw) UnmarshalUBJSON(_ []byte) error {
	return nil
}
