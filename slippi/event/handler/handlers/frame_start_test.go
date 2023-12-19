package handlers_test

import (
	"encoding/binary"
	"fmt"
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
		x := int(binary.BigEndian.Uint32(input[0:4]))
		fmt.Println(x)

		expected := slippi.FrameStart{
			FrameNumber:       frameNumber,
			Seed:              seed,
			SceneFrameCounter: sceneFrameCounter,
		}
		dec := event.Decoder{
			Data: input,
			Size: len(input),
		}

		actual := slippi.Data{}
		err := handlers.FrameStartHandler{}.Parse(&dec, &actual)
		require.NoError(t, err)
		require.Equal(t, expected, actual.Frames[-123].FrameStart)
	})
}

func buildFrameStartInput(frameNumber int32, seed, sceneFrameCounter uint32) []byte {
	input := []byte{byte(event.EventFrameStart)}
	buf := make([]byte, 4)

	binary.BigEndian.PutUint32(buf, uint32(frameNumber))
	input = append(input, buf...)

	clear(buf)
	binary.BigEndian.PutUint32(buf, seed)
	input = append(input, buf...)

	clear(buf)
	binary.BigEndian.PutUint32(buf, sceneFrameCounter)
	input = append(input, buf...)

	return input
}
