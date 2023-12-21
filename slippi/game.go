package slippi

// Data holds the parsed game data of the parsed .slp file.
type Data struct {
	GameStart  GameStart
	Frames     map[int]Frame // Map of FrameNumber -> Frame
	GameEnd    GameEnd
	GeckoCodes GeckoCodeList
}

// Game represents a parsed .slp game.
type Game struct {
	Data Data     `ubjson:"raw"`
	Meta Metadata `ubjson:"metadata"`
}
