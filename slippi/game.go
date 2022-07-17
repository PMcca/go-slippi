package slippi

// Game represents a parsed .slp game.
type Game struct {
	Raw  []byte   `ubjson:"raw"`
	Meta Metadata `ubjson:"metadata"`
	//Meta map[string]interface{} `ubjson:"metadata"`
}

type Game2 struct {
	FieldA string `ubjson:"fielda"`
	FieldB string
}

type Game3 struct {
	FieldC string `ubjson:"fielda"`
}
