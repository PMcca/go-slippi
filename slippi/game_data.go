package slippi

import (
	"bytes"
	"fmt"
	"github.com/PMcca/go-slippi/slippi/melee"
)

type TimerType uint8

const (
	NoTimer         TimerType = 0
	TimerDecreasing TimerType = 2
	TimerIncreasing TimerType = 3
)

type InGameMode uint8

const (
	InGameModeTime  InGameMode = iota // 0
	InGameModeStock                   // 1
	InGameModeCoin                    // 2
	InGameModeBonus                   // 3
)

type ItemSpawnBehaviour int8

const (
	ItemSpawnOff      ItemSpawnBehaviour = iota - 1 // -1
	ItemSpawnVeryLow                                // 0
	ItemSpawnLow                                    // 1
	ItemSpawnMed                                    // 2
	ItemSpawnHigh                                   // 3
	ItemSpawnVeryHigh                               // 4
)

type PlayerType uint8

const (
	PlayerTypeHuman PlayerType = iota // 0
	PlayerTypeCPU                     // 1
	PlayerTypeDemo                    // 2
	PlayerTypeEmpty                   // 3
)

type Scene uint8

const (
	SceneVS     Scene = 0x2
	SceneOnline Scene = 0x8
)

type Language uint8

const (
	LanguageJapanese Language = iota // 0
	LanguageEnglish                  // 1
)

type Player struct {
	Index                  int // Port = Index + 1
	Port                   int
	CharacterID            melee.InternalCharacterID
	PlayerType             PlayerType
	StartStocks            int
	CostumeIndex           int // TODO Same as characterColor? Doesn't make sense to be enum?
	IsInvisible            bool
	IsLowGravity           bool
	IsBlackStockIcon       bool
	IsMetal                bool
	IsStartOnAngelPlatform bool
	IsRumble               bool
	CPULevel               int
	OffenseRatio           float32
	DefenseRation          float32
	ModelScale             int    // TODO or float?
	ControllerFix          string // TODO What is this?
	Nametag                string
	DisplayName            string
	ConnectCode            string
	UserID                 string
}

type GameStart struct {
	SlippiVersion      string
	TimerType          TimerType
	InGameMode         InGameMode
	IsFriendlyFire     bool
	IsTeams            bool
	ItemSpawnBehaviour ItemSpawnBehaviour
	Stage              melee.StageID
	TimerStartSeconds  int
	EnabledItems       []melee.Item
	Players            []Player
	Scene              Scene
	GameMode           int // TODO figure this out
	Language           Language
	RandomSeed         int
	IsPAL              bool
	IsFrozenPS         bool
	MatchID            string
	GameNumber         int
	TiebreakerNumber   int
}

// Data holds the parsed game data of the parsed .slp file.
type Data struct {
	GameStart GameStart
}

// UnmarshalUBJSON takes the 'raw' array from the .slp file and parses it into a Game.Data struct.
func (d *Data) UnmarshalUBJSON(b []byte) error {
	// Beginning of raw array should always be '$U#l'.
	if !bytes.Equal(b[0:4], []byte("$U#l")) {
		return fmt.Errorf("%w:expected '$U#l', found %s", ErrInvalidRawStart, b[0:4])
	}

	dec := decoder{
		data: b[8:], // Skip $U#l and 4 bytes for length; size of array is used for bounds checking.
	}

	eventSizes, err := parseEventPayloads(b[8:]) // Skip $U#l and 4 bytes for length.
	if err != nil {
		return err
	}

	// Main event parsing loop
	for i := 0; i < len(dec.data); i++ {
		eventCode := eventType(dec.read(0x0))
		eventSize, ok := eventSizes[eventCode]
		if !ok {
			return fmt.Errorf("%w:eventCode %d", ErrUnknownEventInEventSizes, eventCode)
		}

		var err error
		switch eventCode {
		case eventPayloadsEvent:
			break // Already parsed, so skip.
		case eventGameStart:
			err = parseGameStart(eventSize, &dec, d)
		}
		if err != nil {
			return fmt.Errorf("%w:failed to parse event %d", err, eventCode)
		}

		dec.data = dec.data[eventSize+1:] // Update the window of data, skipping the # of bytes read + the command byte.
	}
	fmt.Println(eventSizes)

	return nil
}
