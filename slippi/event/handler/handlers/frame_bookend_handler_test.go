package handlers_test

import (
	"github.com/PMcca/go-slippi/internal/testutil"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/event"
	"github.com/PMcca/go-slippi/slippi/event/handler/handlers"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseFrameBookend(t *testing.T) {
	testCases := map[string]struct {
		frameNumber          int
		latestFinalisedFrame int
		expected             slippi.FrameBookend
		errAssertion         require.ErrorAssertionFunc
	}{
		"ParsesFrameBookend": {
			frameNumber:          -123,
			latestFinalisedFrame: 490,
			expected: slippi.FrameBookend{
				FrameNumber:          -123,
				LatestFinalisedFrame: 490,
			},
			errAssertion: require.NoError,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			input := buildFrameBookendInput(int32(tc.frameNumber), int32(tc.latestFinalisedFrame))
			dec := event.Decoder{
				Data: input,
				Size: len(input),
			}

			d := slippi.Data{}
			err := handlers.FrameBookendHandler{}.Parse(&dec, &d)
			tc.errAssertion(t, err)

			actual := d.Frames[tc.frameNumber].FrameBookend
			require.Equal(t, tc.expected, actual)
		})
	}

}

func buildFrameBookendInput(frameNumber, latestFinalFrame int32) []byte {
	out := []byte{byte(event.EventFrameBookend)}
	testutil.PutInt32(&out, frameNumber)
	testutil.PutInt32(&out, latestFinalFrame)

	return out
}
