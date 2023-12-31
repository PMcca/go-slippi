package test

import (
	"github.com/pmcca/go-slippi/slippi"
	"github.com/pmcca/go-slippi/slippi/melee"
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

		actual := mustParseSlippiMeta(t, "replays/metadata.slp")
		require.Equal(t, expected, actual)
	})
	t.Run("ParsesLocalReplay", func(t *testing.T) {
		t.Parallel()
		expected := slippi.Metadata{
			StartAt:   "2023-12-24T09:33:18Z",
			LastFrame: 7075,
			Players: slippi.PlayersMeta{
				Port1: slippi.PlayerMeta{
					Characters: slippi.Characters{
						{
							CharacterID:  melee.Int_Fox,
							FramesPlayed: 7199,
						},
					},
				},
				Port2: slippi.PlayerMeta{
					Characters: slippi.Characters{
						{
							CharacterID:  melee.Int_Marth,
							FramesPlayed: 6324,
						},
					},
				},
				Port3: slippi.PlayerMeta{
					Characters: slippi.Characters{
						{
							CharacterID:  melee.Int_Samus,
							FramesPlayed: 7199,
						},
					},
				},
				Port4: slippi.PlayerMeta{
					Characters: slippi.Characters{
						{
							CharacterID:  melee.Int_Roy,
							FramesPlayed: 2182,
						},
					},
				},
			},
			PlayedOn: "dolphin",
		}

		actual := mustParseSlippiMeta(t, "replays/4-player-offline.slp")
		require.Equal(t, expected, actual)
	})
}

func TestNetplayNamesCodes(t *testing.T) {
	t.Parallel()
	t.Run("ReadsNetplayNamesAndCodes", func(t *testing.T) {
		t.Parallel()

		actual := mustParseSlippiMeta(t, "replays/finalized-frame.slp")
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

		actual := mustParseSlippiMeta(t, "replays/realtimeTest.slp")
		require.Equal(t, "Day 1", actual.ConsoleNick)
	})
}
