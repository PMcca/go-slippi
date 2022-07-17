package melee_test

import (
	"github.com/PMcca/go-slippi/slippi/melee"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInternalCharacterID_String(t *testing.T) {

	internalCharacterIDs := map[melee.InternalCharacterID]string{
		melee.Mario:           "Mario",
		melee.Fox:             "Fox",
		melee.CaptainFalcon:   "Captain Falcon",
		melee.DonkeyKong:      "Donkey Kong",
		melee.Kirby:           "Kirby",
		melee.Bowser:          "Bowser",
		melee.Link:            "Link",
		melee.Sheik:           "Sheik",
		melee.Ness:            "Ness",
		melee.Peach:           "Peach",
		melee.Popo:            "Popo",
		melee.Nana:            "Nana",
		melee.Pikachu:         "Pikachu",
		melee.Samus:           "Samus",
		melee.Yoshi:           "Yoshi",
		melee.Jigglypuff:      "Jigglypuff",
		melee.Mewtwo:          "Mewtwo",
		melee.Luigi:           "Luigi",
		melee.Marth:           "Marth",
		melee.Zelda:           "Zelda",
		melee.YoungLink:       "Young Link",
		melee.DrMario:         "Dr. Mario",
		melee.Falco:           "Falco",
		melee.Pichu:           "Pichu",
		melee.GameAndWatch:    "Game & Watch",
		melee.Ganondorf:       "Ganondorf",
		melee.Roy:             "Roy",
		melee.MasterHand:      "Master Hand",
		melee.CrazyHand:       "Crazy Hand",
		melee.WireFrameMale:   "WireFrame Male",
		melee.WireFrameFemale: "WireFrame Female",
		melee.GigaBowser:      "Giga Bowser",
		melee.Sandbag:         "Sandbag",
	}

	for i := 0; i < 33; i++ {
		c := melee.InternalCharacterID(i)
		expected := internalCharacterIDs[c]
		require.Equal(t, expected, c.String())
	}
}
