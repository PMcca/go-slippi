package test

import (
	goslippi "github.com/PMcca/go-slippi"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/melee"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMetadataParse(t *testing.T) {
	t.Parallel()
	t.Run("SuccessfullyParsesMetadata", func(t *testing.T) {
		t.Parallel()

		expected := slippi.Metadata{
			StartAt:   "2022-12-02T18:09:00Z",
			LastFrame: 2011,
			Players: slippi.PlayersMeta{
				Port1: slippi.PlayerMeta{
					Names: slippi.Names{
						Name:       "Smasher",
						SlippiCode: "SMSH#123",
					},
					Characters: slippi.Characters{
						{
							CharacterID:  melee.Int_Zelda,
							FramesPlayed: 532,
						},
						{
							CharacterID:  melee.Int_Sheik,
							FramesPlayed: 1603,
						},
					},
				},
				Port2: slippi.PlayerMeta{
					Names: slippi.Names{
						Name:       "I Love Slippi!",
						SlippiCode: "SLIP#987",
					},
					Characters: slippi.Characters{
						{
							CharacterID:  melee.Int_Mario,
							FramesPlayed: 2135,
						},
					},
				},
			},
			PlayedOn: "dolphin",
		}

		filePath := "replays/metadata.slp"
		actual, err := goslippi.ParseMeta(filePath)
		require.NoError(t, err)

		assert.Equal(t, expected, actual)
	})
}

func TestNetplayNamesCodes(t *testing.T) {
	t.Parallel()
	t.Run("ReadsNetplayNamesAndCodes", func(t *testing.T) {
		t.Parallel()

		filePath := "replays/finalizedFrame.slp"
		actual, err := goslippi.ParseMeta(filePath)
		require.NoError(t, err)

		assert.Equal(t, "V", actual.Players.Port1.Names.Name)
		assert.Equal(t, "VA#0", actual.Players.Port1.Names.SlippiCode)
		assert.Equal(t, "Fizzi", actual.Players.Port2.Names.Name)
		assert.Equal(t, "FIZZI#36", actual.Players.Port2.Names.SlippiCode)
	})
}

func TestConsoleNickname(t *testing.T) {
	t.Parallel()
	t.Run("ReadsConsoleNickname", func(t *testing.T) {
		t.Parallel()

		filePath := "replays/realtimeTest.slp"
		actual, err := goslippi.ParseMeta(filePath)
		require.NoError(t, err)

		assert.Equal(t, "Day 1", actual.ConsoleNick)
	})
}
