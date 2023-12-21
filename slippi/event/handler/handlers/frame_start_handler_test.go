package handlers_test

import (
	"github.com/PMcca/go-slippi/internal/testutil"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/event"
	"github.com/PMcca/go-slippi/slippi/event/handler/handlers"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseFrameStart(t *testing.T) {
	t.Parallel()
	t.Run("ParsesFrameStartEvent", func(t *testing.T) {
		t.Parallel()
		frameNumber := -123
		seed := uint32(19283746)
		sceneFrameCounter := uint32(123)
		input := buildFrameStartInput(int32(frameNumber), seed, sceneFrameCounter)

		expected := slippi.Frame{
			FrameNumber: frameNumber,
			FrameStart: slippi.FrameStart{
				FrameNumber:       frameNumber,
				Seed:              seed,
				SceneFrameCounter: sceneFrameCounter,
			},
		}
		dec := event.Decoder{
			Data: input,
			Size: len(input),
		}

		d := slippi.Data{}
		err := handlers.FrameStartHandler{}.Parse(&dec, &d)
		require.NoError(t, err)

		actual := d.Frames[frameNumber]
		require.Equal(t, expected, actual)
	})
}

func buildFrameStartInput(frameNumber int32, seed, sceneFrameCounter uint32) []byte {
	input := []byte{byte(event.EventFrameStart)}
	testutil.PutInt32(&input, frameNumber)
	testutil.PutUint32(&input, seed)
	testutil.PutUint32(&input, sceneFrameCounter)
	return input
}
