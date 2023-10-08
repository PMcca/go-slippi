package slippi

// Game represents a parsed .slp game.
type Game struct {
	Raw  []byte   `ubjson:"raw"`
	Meta Metadata `ubjson:"metadata"`
}

// ignoreRaw is used with metaOnly game. It implements ubjson.UnmarshalUBJSON to simply return nil and skip over the
// reading of the raw element.
type ignoreRaw struct{}

// metaOnlyGame is an internal type to omit parsing the 'raw' element and only parse the 'metadata' fields.
type metaOnlyGame struct {
	IgnoreRaw ignoreRaw `ubjson:"raw"`
	Meta      Metadata  `ubjson:"metadata"`
}

// UnmarshalUBJSON returns immediately from reading raw, for performance benefits.
func (i ignoreRaw) UnmarshalUBJSON(_ []byte) error {
	return nil
}
