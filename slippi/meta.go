package slippi

// Metadata represents the parsed metadata element from a .slp file.
type Metadata struct {
	StartAt   string                 `ubjson:"startAt"`
	LastFrame int                    `ubjson:"lastFrame"`
	Players   map[string]interface{} `ubjson:"players"`
	PlayedOn  string                 `ubjson:"playedOn"`
}

type Player struct {
}
