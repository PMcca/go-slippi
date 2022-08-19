package melee_test

import (
	"github.com/PMcca/go-slippi/slippi/melee"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInternalCharacterID_String(t *testing.T) {

	internalCharacterIDs := map[melee.InternalCharacterID]string{
		0:  "Mario",
		1:  "Fox",
		2:  "Captain Falcon",
		3:  "Donkey Kong",
		4:  "Kirby",
		5:  "Bowser",
		6:  "Link",
		7:  "Sheik",
		8:  "Ness",
		9:  "Peach",
		10: "Popo",
		11: "Nana",
		12: "Pikachu",
		13: "Samus",
		14: "Yoshi",
		15: "Jigglypuff",
		16: "Mewtwo",
		17: "Luigi",
		18: "Marth",
		19: "Zelda",
		20: "Young Link",
		21: "Dr. Mario",
		22: "Falco",
		23: "Pichu",
		24: "Game & Watch",
		25: "Ganondorf",
		26: "Roy",
		27: "Master Hand",
		28: "Crazy Hand",
		29: "WireFrame Male",
		30: "WireFrame Female",
		31: "Giga Bowser",
		32: "Sandbag",
	}

	for i := 0; i < 33; i++ {
		c := melee.InternalCharacterID(i)
		expected := internalCharacterIDs[c]
		require.Equal(t, expected, c.String())
	}
}
