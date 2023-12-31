package test

import (
	"github.com/pmcca/go-slippi/slippi"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestItem(t *testing.T) {
	t.Parallel()
	filePath := "replays/item-export.slp"

	t.Run("MonotonicallyIncrementsItemSpawnID", func(t *testing.T) {
		t.Parallel()

		actual := mustParseSlippiGame(t, filePath)
		lastSpawnID := -1
		frames := actual.Data.Frames
		i := slippi.FirstFrame
		for {
			frame, ok := frames[i]
			if !ok {
				break
			}
			for _, item := range frame.ItemUpdates {
				if lastSpawnID < int(item.SpawnID) {
					require.Equal(t, lastSpawnID+1, int(item.SpawnID), "Frame: %d", i)
					lastSpawnID = int(item.SpawnID)
				}
			}

			i++
		}
	})
	t.Run("ItemsHaveValidOwnerIDs", func(t *testing.T) {
		t.Parallel()

		actual := mustParseSlippiGame(t, filePath)
		frames := actual.Data.Frames
		i := slippi.FirstFrame
		for {
			frame, ok := frames[i]
			if !ok {
				break
			}
			for _, item := range frame.ItemUpdates {
				require.LessOrEqual(t, item.Owner, int8(3))
				require.GreaterOrEqual(t, item.Owner, int8(-1))
			}

			i++
		}
	})
}
