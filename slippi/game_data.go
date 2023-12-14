package slippi

import (
	"bytes"
	"fmt"
	"github.com/PMcca/go-slippi/internal/util"
	"github.com/PMcca/go-slippi/slippi/melee"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"strings"
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

type TeamShade uint8

const (
	TeamShadeNormal TeamShade = iota // 0
	TeamShadeLight                   // 1
	TeamShadeDark                    // 2
)

// TeamColour is the colour of the team the player is in. Value is TeamID.
type TeamColour uint8

const (
	TeamColourRed   TeamColour = iota // 0
	TeamColourBlue                    // 1
	TeamColourGreen                   // 2
)

type GameMode uint8

const (
	GameModeVS     GameMode = 0x2
	GameModeOnline GameMode = 0x8
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
	StartStocks            uint8
	CostumeIndex           uint8
	TeamShade              TeamShade
	Handicap               uint8
	TeamColour             TeamColour
	IsStamina              bool
	IsSilent               bool
	IsLowGravity           bool
	IsInvisible            bool
	IsBlackStockIcon       bool
	IsMetal                bool
	IsStartOnAngelPlatform bool
	IsRumbleEnabled        bool
	CPULevel               uint8
	OffenseRatio           float32
	DefenseRation          float32
	ModelScale             float32 // TODO or float?
	ControllerFix          string  // TODO What is this?
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
	Stage              melee.Stage
	TimerStartSeconds  int
	EnabledItems       []melee.Item
	Players            []Player
	Scene              uint8 // Minor scene, should always be 0x2
	GameMode           GameMode
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
		dec.size = eventSize

		var err error
		switch eventCode {
		case eventPayloadsEvent:
			break // Already parsed, so skip.
		case eventGameStart:
			err = parseGameStart(&dec, d)
		}
		if err != nil {
			return fmt.Errorf("%w:failed to parse event %d", err, eventCode)
		}

		dec.data = dec.data[eventSize+1:] // Update the window of data, skipping the # of bytes read + the command byte.
	}

	return nil
}

// parsePlayer takes a controller port and uses it as the offset for parsing & returning the corresponding Player.
func parsePlayer(playerIndex int, dec *decoder) (Player, error) {
	offset := playerIndex * 0x8

	dashBack := dec.readInt32(0x141 + offset)
	shieldDrop := dec.readInt32(0x145 + offset)

	var controllerFix string
	switch {
	case dashBack != shieldDrop:
		controllerFix = "Mixed"
	case dashBack == 1:
		controllerFix = "UCF"
	case dashBack == 2:
		controllerFix = "Dween"
	default:
		controllerFix = "None"
	}

	jisDecoder := japanese.ShiftJIS.NewDecoder()
	// Start is the length of the string * playerIndex, + the offset.
	nameTag, err := parseGameStartString((0x10*playerIndex)+0x161, 0x10, dec, jisDecoder, true)
	if err != nil {
		return Player{}, fmt.Errorf("%w:failed to parse name tag", err)
	}
	displayName, err := parseGameStartString((0x1f*playerIndex)+0x1a5, 0x1a5, dec, jisDecoder, true)
	if err != nil {
		return Player{}, fmt.Errorf("%w:failed to parse display name", err)
	}
	connectCode, err := parseGameStartString((0xa*playerIndex)+0x221, 0x221, dec, jisDecoder, true)
	if err != nil {
		return Player{}, fmt.Errorf("%w:failed to parse connect code", err)
	}
	userID, err := parseGameStartString((0x1d*playerIndex)+0x249, 0x249, dec, unicode.UTF8.NewDecoder(), false)
	if err != nil {
		return Player{}, fmt.Errorf("%w:failed to parse userID", err)
	}

	// Update offset and fetch remaining fields..
	offset = playerIndex * 0x24
	playerBitfield := dec.read(0x6c + playerIndex*0x24)

	return Player{
		Index:                  playerIndex,
		Port:                   playerIndex + 1,
		CharacterID:            melee.InternalCharacterID(dec.read(0x65 + offset)),
		PlayerType:             PlayerType(dec.read(0x66 + offset)),
		StartStocks:            dec.read(0x67 + offset),
		CostumeIndex:           dec.read(0x68 + offset),
		TeamShade:              TeamShade(dec.read(0x6c + offset)),
		Handicap:               dec.read(0x6d + offset),
		TeamColour:             TeamColour(dec.read(0x6e + offset)),
		IsStamina:              playerBitfield&0x01 > 0,
		IsSilent:               playerBitfield&0x02 > 0,
		IsLowGravity:           playerBitfield&0x04 > 0,
		IsInvisible:            playerBitfield&0x08 > 0,
		IsBlackStockIcon:       playerBitfield&0x10 > 0,
		IsMetal:                playerBitfield&0x20 > 0,
		IsStartOnAngelPlatform: playerBitfield&0x40 > 0,
		IsRumbleEnabled:        playerBitfield&0x80 > 0,
		CPULevel:               dec.read(0x74 + offset),
		OffenseRatio:           dec.readFloat32(0x7d + offset),
		DefenseRation:          dec.readFloat32(0x81 + offset),
		ModelScale:             dec.readFloat32(0x85 + offset),
		ControllerFix:          controllerFix,
		Nametag:                nameTag,
		DisplayName:            displayName,
		ConnectCode:            connectCode,
		UserID:                 userID,
	}, nil
}

// parseGameStartString parses a string, such as a displayName or connect code, by decoding the respective bytes using the given
// transformer, and optionally (if in Shift JIS), halving the width of the resulting bytes.
func parseGameStartString(stringStart, stringLength int, dec *decoder, transformer transform.Transformer, toHalfWidth bool) (string, error) {
	stringBuf := dec.readN(stringStart, stringStart+stringLength)
	t, _, err := transform.Bytes(transformer, stringBuf)
	if err != nil {
		return "", err
	}

	result := strings.Split(string(t), "\x00")[0] // Strip any nil's
	if result != "" && toHalfWidth {
		result = util.ToHalfWidthChars(result)
	}

	return result, nil
}
