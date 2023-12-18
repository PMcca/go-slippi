package slippi

// Game represents a parsed .slp game.
type Game struct {
	Data Data     `ubjson:"raw"`
	Meta Metadata `ubjson:"metadata"`
}
