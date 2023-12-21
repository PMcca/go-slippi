package slippi

import (
	"encoding/binary"
	"fmt"
	"github.com/PMcca/go-slippi/slippi/melee"
	"strconv"
)

// Metadata represents the parsed metadata element from a .slp file.
type Metadata struct {
	StartAt     string      `ubjson:"startAt"`
	LastFrame   int         `ubjson:"lastFrame"`
	Players     PlayersMeta `ubjson:"players"`
	PlayedOn    string      `ubjson:"playedOn"`
	ConsoleNick string      `ubjson:"consoleNick"`
}

// PlayerMeta is a single player in the game including their Slippi display name or in-game name (dependent on online/local).
type PlayerMeta struct {
	Names      Names      `ubjson:"names"`
	Characters Characters `ubjson:"characters"`
}

// PlayersMeta holds a PlayerMeta object for each controller port. The corresponding object for a given port will be empty if
// the port was not used during the game.
type PlayersMeta struct {
	Port1 PlayerMeta `ubjson:"0"`
	Port2 PlayerMeta `ubjson:"1"`
	Port3 PlayerMeta `ubjson:"2"`
	Port4 PlayerMeta `ubjson:"3"`
}

// Character is the Melee character that was present in the match.
type Character struct {
	CharacterID  melee.InternalCharacterID
	FramesPlayed int
}

// Characters is the list of characters that a player used during a game.
type Characters []Character

// Names holds the player's name as well as the Slippi code, if any.
type Names struct {
	Name       string `ubjson:"netplay"`
	SlippiCode string `ubjson:"code"`
}

// UnmarshalUBJSON implements the ubjson.Unmarshaler interface. It unmarshals the UBJSON bytes of the characters into
// their Characters type. This makes strong assumptions about the format of the characters UBJSON object; if it changes
// or is updated, this function will break.
func (c *Characters) UnmarshalUBJSON(b []byte) error {
	i := 0
	for {
		// First byte will be 'U', so skip over it.
		i++
		if i >= len(b) {
			break
		}

		// Read the size of the ID in bytes
		idLen := int(b[i])
		i++

		characterID, err := strconv.Atoi(string(b[i : i+idLen]))
		if err != nil {
			return fmt.Errorf("%w:failed to convert metadata CharacterID to int", err)
		}
		i += idLen

		character := Character{}
		character.CharacterID = melee.InternalCharacterID(characterID)

		// Next byte = 'l', read next 4 bytes
		if b[i] != 'l' {
			return fmt.Errorf("invalid UBJSON type for 'framesPlayed'. expected l, got %c", b[i])
		}
		frames := binary.BigEndian.Uint32(b[i+1 : i+5])
		i += 5
		character.FramesPlayed = int(frames)

		*c = append(*c, character)
	}

	return nil
}
