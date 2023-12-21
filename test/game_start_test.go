package test

import (
	goslippi "github.com/PMcca/go-slippi"
	"github.com/PMcca/go-slippi/slippi/melee"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGameStart(t *testing.T) {
	t.Parallel()
	t.Run("Parses0.1.0GameStartFromSlippiJS", func(t *testing.T) {
		filePath := "replays/sheik_vs_ics_yoshis.slp"
		actual, err := goslippi.ParseGame(filePath)
		require.NoError(t, err)

		gameStart := actual.Data.GameStart
		require.Len(t, gameStart.Players, 4)

		assert.Equal(t, melee.StageYoshisStory, gameStart.Stage)
		assert.Equal(t, melee.Ext_Sheik, gameStart.Players[0].CharacterID)
		assert.Equal(t, melee.Ext_IceClimbers, gameStart.Players[1].CharacterID)
		assert.Equal(t, "0.1.0", gameStart.SlippiVersion)
	})
}

func TestNametags(t *testing.T) {
	t.Parallel()
	t.Run("ParsesNametag1", func(t *testing.T) {
		t.Parallel()

		filePath := "replays/nametags.slp"
		actual, err := goslippi.ParseGame(filePath)
		require.NoError(t, err)

		require.Len(t, actual.Data.GameStart.Players, 4)
		assert.Equal(t, "AMNイ", actual.Data.GameStart.Players[0].Nametag)
	})
	t.Run("ParsesNametag2", func(t *testing.T) {
		t.Parallel()

		filePath := "replays/nametags2.slp"
		actual, err := goslippi.ParseGame(filePath)
		require.NoError(t, err)

		require.Len(t, actual.Data.GameStart.Players, 4)
		assert.Equal(t, "A1=$", actual.Data.GameStart.Players[0].Nametag)
		assert.Equal(t, "か、9@", actual.Data.GameStart.Players[1].Nametag)
	})
	t.Run("ParsesNametag3", func(t *testing.T) {
		t.Parallel()

		filePath := "replays/nametags3.slp"
		actual, err := goslippi.ParseGame(filePath)
		require.NoError(t, err)

		require.Len(t, actual.Data.GameStart.Players, 4)
		assert.Equal(t, "B  R", actual.Data.GameStart.Players[0].Nametag)
		assert.Equal(t, ".  。", actual.Data.GameStart.Players[1].Nametag)
	})
}

func TestPAL(t *testing.T) {
	t.Parallel()
	t.Run("ReadsPALGame", func(t *testing.T) {
		t.Parallel()
		filePath := "replays/pal.slp"
		actual, err := goslippi.ParseGame(filePath)
		require.NoError(t, err)

		assert.Equal(t, true, actual.Data.GameStart.IsPAL)
	})
	t.Run("PALIsFalseWithNTSCGame", func(t *testing.T) {
		t.Parallel()

		filePath := "replays/ntsc.slp"
		actual, err := goslippi.ParseGame(filePath)
		require.NoError(t, err)

		assert.Equal(t, false, actual.Data.GameStart.IsPAL)
	})
}

func TestControllerFix(t *testing.T) {
	t.Parallel()
	t.Run("ReadsDifferentControllerFixes", func(t *testing.T) {
		t.Parallel()

		filePath := "replays/controllerFixes.slp"
		actual, err := goslippi.ParseGame(filePath)
		require.NoError(t, err)

		assert.Equal(t, "Dween", actual.Data.GameStart.Players[0].ControllerFix)
		assert.Equal(t, "UCF", actual.Data.GameStart.Players[1].ControllerFix)
		assert.Equal(t, "None", actual.Data.GameStart.Players[2].ControllerFix)
	})
}

func TestMatchInfo(t *testing.T) {
	t.Parallel()
	t.Run("ReadsRankedMatchInfo", func(t *testing.T) {
		t.Parallel()

		actual, err := goslippi.ParseGame("replays/ranked_game1_tiebreak.slp")
		require.NoError(t, err)

		gameStart := actual.Data.GameStart
		assert.Equal(t, 1, gameStart.GameNumber, "GameNumber not equal")
		assert.Equal(t, 1, gameStart.TiebreakerNumber, "TieBreakerNumber not equal")
		assert.Equal(t, "mode.ranked-2022-12-20T05:36:47.50-0", gameStart.MatchID)
	})

	t.Run("ReadsUnrankedGame1", func(t *testing.T) {
		t.Parallel()

		actual, err := goslippi.ParseGame("replays/unranked_game1.slp")
		require.NoError(t, err)

		gameStart := actual.Data.GameStart
		assert.Equal(t, 1, gameStart.GameNumber, "GameNumber not equal")
		assert.Equal(t, 0, gameStart.TiebreakerNumber, "TieBreakerNumber not equal")
		assert.Equal(t, "mode.unranked-2022-12-21T02:26:27.50-0", gameStart.MatchID)
	})

	t.Run("ReadsUnrankedGame2", func(t *testing.T) {
		t.Parallel()

		actual, err := goslippi.ParseGame("replays/unranked_game2.slp")
		require.NoError(t, err)

		gameStart := actual.Data.GameStart
		assert.Equal(t, 2, gameStart.GameNumber, "GameNumber not equal")
		assert.Equal(t, 0, gameStart.TiebreakerNumber, "TieBreakerNumber not equal")
		assert.Equal(t, "mode.unranked-2022-12-21T02:26:27.50-0", gameStart.MatchID)
	})

	t.Run("EmptyValuesForOldReplay", func(t *testing.T) {
		t.Parallel()

		actual, err := goslippi.ParseGame("replays/old-test.slp")
		require.NoError(t, err)

		gameStart := actual.Data.GameStart
		assert.Equal(t, 0, gameStart.GameNumber, "GameNumber not equal")
		assert.Equal(t, 0, gameStart.TiebreakerNumber, "TieBreakerNumber not equal")
		assert.Equal(t, "", gameStart.MatchID)
	})
}
