package slippi

import (
	"fmt"
	"github.com/PMcca/go-slippi/slippi/melee"
)

type TimerType int

const (
	NoTimer         TimerType = 0
	TimerDecreasing           = 2
	TimerIncreasing           = 3
)

type GameMode int

const (
	GameModeTime  = 0
	GameModeStock = 1
	GameModeCoin  = 2
	GameModeBonus = 3
)

type ItemSpawnBehaviour int8

const (
	ItemSpawnOff      = -1
	ItemSpawnVeryLow  = 0
	ItemSpawnLow      = 1
	ItemSpawnMed      = 2
	ItemSpawnHigh     = 3
	ItemSpawnVeryHigh = 4
)

type GameStart struct {
	SlippiVersion      string
	TimerType          TimerType
	GameMode           GameMode
	IsFriendlyFire     bool
	IsTeams            bool
	ItemSpawnBehaviour ItemSpawnBehaviour
	Stage              melee.StageID
	TimerStartSeconds  int
}

type GameEnd struct {
}

type Frame struct {
}

// Data holds the parsed game data of the parsed .slp file.
type Data struct {
}

func (r Data) UnmarshalUBJSON(bytes []byte) error {
	fmt.Println("HELLO I AM IN RAW")
	return nil
}

// Game represents a parsed .slp game.
type Game struct {
	Raw Data `ubjson:"raw"`
	//Data  []byte   `ubjson:"raw"`
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
