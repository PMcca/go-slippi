package test

import (
	"github.com/PMcca/go-slippi/slippi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNametags(t *testing.T) {
	t.Parallel()

	t.Run("ParsesNametag1", func(t *testing.T) {
		t.Parallel()

		filePath := "replays/nametags.slp"
		actual, err := slippi.ParseGame(filePath)
		require.NoError(t, err)

		require.Len(t, actual.Data.GameStart.Players, 4)
		assert.Equal(t, "AMNイ", actual.Data.GameStart.Players[0].Nametag)
	})
	t.Run("ParsesNametag2", func(t *testing.T) {
		t.Parallel()

		filePath := "replays/nametags2.slp"
		actual, err := slippi.ParseGame(filePath)
		require.NoError(t, err)

		require.Len(t, actual.Data.GameStart.Players, 4)
		assert.Equal(t, "A1=$", actual.Data.GameStart.Players[0].Nametag)
		assert.Equal(t, "か、9@", actual.Data.GameStart.Players[1].Nametag)
	})
	t.Run("ParsesNametag3", func(t *testing.T) {
		t.Parallel()

		filePath := "replays/nametags3.slp"
		actual, err := slippi.ParseGame(filePath)
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
		actual, err := slippi.ParseGame(filePath)
		require.NoError(t, err)

		assert.Equal(t, true, actual.Data.GameStart.IsPAL)
	})
	t.Run("PALIsFalseWithNTSCGame", func(t *testing.T) {
		t.Parallel()

		filePath := "replays/ntsc.slp"
		actual, err := slippi.ParseGame(filePath)
		require.NoError(t, err)

		assert.Equal(t, false, actual.Data.GameStart.IsPAL)
	})
}

func TestControllerFix(t *testing.T) {
	t.Parallel()

	t.Run("ReadsDifferentControllerFixes", func(t *testing.T) {
		t.Parallel()

		filePath := "replays/controllerFixes.slp"
		actual, err := slippi.ParseGame(filePath)
		require.NoError(t, err)

		assert.Equal(t, "Dween", actual.Data.GameStart.Players[0].ControllerFix)
		assert.Equal(t, "UCF", actual.Data.GameStart.Players[1].ControllerFix)
		assert.Equal(t, "None", actual.Data.GameStart.Players[2].ControllerFix)
	})
}

func TestMatchInfo(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		filePath                 string
		expectedGameNumber       int
		expectedTiebreakerNumber int
		expectedMatchID          string
	}{
		"ReadsRankedMatchInfo": {
			filePath:                 "replays/ranked_game1_tiebreak.slp",
			expectedGameNumber:       1,
			expectedTiebreakerNumber: 1,
			expectedMatchID:          "mode.ranked-2022-12-20T05:36:47.50-0",
		},
	}

	for name, testCase := range testCases {
		tc := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := slippi.ParseGame(tc.filePath)

			require.NoError(t, err)
			gameStart := actual.Data.GameStart
			require.Equal(t, tc.expectedGameNumber, gameStart.GameNumber)
			require.Equal(t, tc.expectedTiebreakerNumber, gameStart.TiebreakerNumber)
			require.Equal(t, tc.expectedMatchID, gameStart.MatchID)
		})
	}
}
