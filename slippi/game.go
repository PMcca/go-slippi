package slippi

// Game represents a parsed .slp game.
type Game struct {
	Raw  []byte   `ubjson:"raw"`
	Meta Metadata `ubjson:"metadata"`
	//Meta map[string]interface{} `ubjson:"metadata"`
}

// metaOnlyGame is an internal type to omit parsing the 'raw' element and only parse the 'metadata' fields.
type metaOnlyGame struct {
	Meta Metadata `ubjson:"metadata"`
}
