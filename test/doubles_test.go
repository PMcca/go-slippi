package test

import (
	goslippi "github.com/PMcca/go-slippi"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlayersEliminated(t *testing.T) {
	t.Parallel()

	filePath := "replays/doubles.slp"
	actual, err := goslippi.ParseGame(filePath)
	require.NoError(t, err)
	require.Len(t, actual.Data.GameStart.Players, 4)

	p1ElimFrame := 7754
	p1StockStealFrame := 7783
	p1ElimFrame2 := 8236
	gameEndFrame := actual.Meta.LastFrame
	frames := actual.Data.Frames

	require.Contains(t, frames, p1ElimFrame)
	require.Contains(t, frames, p1StockStealFrame)
	require.Contains(t, frames, p1ElimFrame2)
	require.Contains(t, frames, gameEndFrame)

	// Check that p1 still has a frame when they get eliminated
	require.Len(t, frames[p1ElimFrame].Players, 4)
	require.Contains(t, frames[p1ElimFrame].Players, uint8(0))

	// Ensure all frames between p1ElimFrame+1 and p1StockStealFrame don't have player 0.
	for i := p1ElimFrame + 1; i < p1StockStealFrame; i++ {
		require.Len(t, frames[i].Players, 3,
			"Expected length of players to be 3 after eliminated, was %d on frame %d",
			len(frames[i].Players), i)
		require.NotContains(t, frames[i].Players, uint8(0), "Frame %d should not contain playerIndex 0", i)
	}

	// Ensure player 0 now exists after stock steal, until they get eliminated again.
	for i := p1StockStealFrame; i <= p1ElimFrame2; i++ {
		require.Len(t, frames[i].Players, 4,
			"Expected length of p1ElimFrame+i players to be 4, was %d on frame %d",
			len(frames[i].Players), i)
		require.Contains(t, frames[i].Players, uint8(0),
			"Frame %d does not contain player 0 after stock steal", i)
	}

	for i := p1ElimFrame2 + 1; i <= gameEndFrame; i++ {
		for i := p1ElimFrame + 1; i < p1StockStealFrame; i++ {
			require.Len(t, frames[i].Players, 3,
				"Expected length of players to be 3 after eliminated, was %d on frame %d",
				len(frames[i].Players), i)
			require.NotContains(t, frames[i].Players, uint8(0), "Frame %d should not contain playerIndex 0", i)
		}
	}
}
