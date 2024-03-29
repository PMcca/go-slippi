package handlers

import (
	"fmt"
	"github.com/pmcca/go-slippi/slippi"
	"github.com/pmcca/go-slippi/slippi/event"
	"github.com/pmcca/go-slippi/slippi/melee"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"strings"
)

// GameStartHandler handles the parsing of GameStart events.
type GameStartHandler struct{}

// Parse implements the handler.EventHandler interface. It parses a GameStart event and puts its output into the
// given slippi.Data struct.
func (h GameStartHandler) Parse(dec *event.Decoder, data *slippi.Data) error {
	slippiVersion := fmt.Sprintf("%d.%d.%d",
		dec.Read(0x1),
		dec.Read(0x2),
		dec.Read(0x3),
	)

	var isFriendlyFire bool
	if dec.ReadWithBitmask(0x6, 0x01) > 0 {
		isFriendlyFire = true
	}

	enabledItems := melee.GetEnabledItems(
		dec.Read(0x28),
		dec.Read(0x29),
		dec.Read(0x2A),
		dec.Read(0x2B),
		dec.Read(0x2C))

	players := make([]slippi.Player, 4)
	for i := 0; i < 4; i++ {
		p, err := parsePlayer(i, dec)
		if err != nil {
			return fmt.Errorf("%w:failed to parse player in port %d", err, i)
		}
		players[i] = p
	}

	matchID, err := parseGameStartString(0x2be, 51, dec, unicode.UTF8.NewDecoder(), false)
	if err != nil {
		return fmt.Errorf("%w:failed to parse matchID", err)
	}

	data.GameStart = slippi.GameStart{
		SlippiVersion:      slippiVersion,
		TimerType:          slippi.TimerType(dec.ReadWithBitmask(0x5, 0x03)),
		InGameMode:         slippi.InGameMode(dec.ReadWithBitmask(0x5, 0xe0)) >> 5,
		IsFriendlyFire:     isFriendlyFire,
		IsTeams:            dec.ReadBool(0xd),
		ItemSpawnBehaviour: slippi.ItemSpawnBehaviour(dec.Read(0x10)),
		Stage:              melee.Stage(dec.ReadUint16(0x13)),
		TimerStartSeconds:  dec.ReadInt32(0x15),
		EnabledItems:       enabledItems,
		Players:            players,
		Scene:              dec.Read(0x1a3),
		GameMode:           slippi.GameMode(dec.Read(0x1a4)),
		Language:           slippi.Language(dec.Read(0x2bd)),
		RandomSeed:         dec.ReadUint32(0x13d),
		IsPAL:              dec.ReadBool(0x1a1),
		IsFrozenPS:         dec.ReadBool(0x1a2),
		MatchID:            matchID,
		GameNumber:         dec.ReadInt32(0x2f1),
		TiebreakerNumber:   dec.ReadInt32(0x2f5),
	}
	return nil
}

// parsePlayer takes a controller port and uses it as the offset for parsing & returning the corresponding Player.
func parsePlayer(playerIndex int, dec *event.Decoder) (slippi.Player, error) {
	offset := playerIndex * 0x8

	dashBack := dec.ReadUint32(0x141 + offset)
	shieldDrop := dec.ReadUint32(0x145 + offset)

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
		return slippi.Player{}, fmt.Errorf("%w:failed to parse name tag", err)
	}
	displayName, err := parseGameStartString((0x1f*playerIndex)+0x1a5, 0x1f, dec, jisDecoder, true)
	if err != nil {
		return slippi.Player{}, fmt.Errorf("%w:failed to parse display name", err)
	}
	connectCode, err := parseGameStartString((0xa*playerIndex)+0x221, 0xa, dec, jisDecoder, true)
	if err != nil {
		return slippi.Player{}, fmt.Errorf("%w:failed to parse connect code", err)
	}
	userID, err := parseGameStartString((0x1d*playerIndex)+0x249, 0x1d, dec, unicode.UTF8.NewDecoder(), false)
	if err != nil {
		return slippi.Player{}, fmt.Errorf("%w:failed to parse userID", err)
	}

	// Update offset and fetch remaining fields..
	offset = playerIndex * 0x24
	playerBitfield := dec.Read(0x6c + playerIndex*0x24)

	return slippi.Player{
		Index:                  uint8(playerIndex),
		Port:                   int8(playerIndex + 1),
		CharacterID:            melee.ExternalCharacterID(dec.Read(0x65 + offset)),
		PlayerType:             slippi.PlayerType(dec.Read(0x66 + offset)),
		StartStocks:            dec.Read(0x67 + offset),
		CostumeIndex:           dec.Read(0x68 + offset),
		TeamShade:              slippi.TeamShade(dec.Read(0x6c + offset)),
		Handicap:               dec.Read(0x6d + offset),
		TeamColour:             slippi.TeamColour(dec.Read(0x6e + offset)),
		IsStamina:              playerBitfield&0x01 > 0,
		IsSilent:               playerBitfield&0x02 > 0,
		IsLowGravity:           playerBitfield&0x04 > 0,
		IsInvisible:            playerBitfield&0x08 > 0,
		IsBlackStockIcon:       playerBitfield&0x10 > 0,
		IsMetal:                playerBitfield&0x20 > 0,
		IsStartOnAngelPlatform: playerBitfield&0x40 > 0,
		IsRumbleEnabled:        playerBitfield&0x80 > 0,
		CPULevel:               dec.Read(0x74 + offset),
		OffenseRatio:           dec.ReadFloat32(0x7d + offset),
		DefenseRatio:           dec.ReadFloat32(0x81 + offset),
		ModelScale:             dec.ReadFloat32(0x85 + offset),
		ControllerFix:          controllerFix,
		NameTag:                nameTag,
		DisplayName:            displayName,
		ConnectCode:            connectCode,
		UserID:                 userID,
	}, nil
}

// parseGameStartString parses a string, such as a displayName or connect code, by decoding the respective bytes using the given
// transformer, and optionally (if in Shift JIS), halving the width of the resulting bytes.
func parseGameStartString(stringStart, stringLength int, dec *event.Decoder, transformer transform.Transformer, toHalfWidth bool) (string, error) {
	stringBuf := dec.ReadN(stringStart, stringStart+stringLength)
	t, _, err := transform.Bytes(transformer, stringBuf)
	if err != nil {
		return "", err
	}

	result := strings.Split(string(t), "\x00")[0] // Strip any nil's
	if result != "" && toHalfWidth {
		result = toHalfWidthChars(result)
	}

	return result, nil
}

// toHalfWidthChars takes a Shift JIS-decoded string and converts specific bytes to half-width.
func toHalfWidthChars(s string) string {
	var ret []rune
	for _, c := range s {
		switch {
		case c > 0xff00 && c < 0xff5f:
			ret = append(ret, 0x0020+(c-0xff00))
		case c == 0x3000: // Space
			ret = append(ret, 0x0020)
		case c == 0x2019: //Single quote (')
			ret = append(ret, 0x0027)
		case c == 0x201d: // Double quote (")
			ret = append(ret, 0x0022)
		default:
			ret = append(ret, c)
		}
	}

	return string(ret)
}
