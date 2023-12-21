package handlers_test

import (
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/event"
	"github.com/PMcca/go-slippi/slippi/event/handler/handlers"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	gameEndMethod     = slippi.GameEndGameSet
	gameEndLRAS       = int8(1)
	gameEndPlacement0 = int8(1)
	gameEndPlacement1 = int8(2)
	gameEndPlacement2 = int8(-1)
	gameEndPlacement3 = int8(3)
)

func TestParseGameEnd(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		expected     slippi.GameEnd
		errAssertion require.ErrorAssertionFunc
	}{
		"CreatesItemUpdateForItem": {
			expected: slippi.GameEnd{
				GameEndMethod:  slippi.GameEndGameSet,
				LRASInitiatior: 1,
				PlayerPlacements: []slippi.PlayerPlacement{
					{
						PlayerIndex: 0,
						Placement:   1,
					},
					{
						PlayerIndex: 1,
						Placement:   2,
					},
					{
						PlayerIndex: 2,
						Placement:   -1,
					},
					{
						PlayerIndex: 3,
						Placement:   3,
					},
				},
			},
			errAssertion: require.NoError,
		},
	}

	for name, testCase := range testCases {
		tc := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			input := buildGameEbdInput()
			dec := event.Decoder{
				Data: input,
				Size: len(input),
			}

			d := slippi.Data{}
			err := handlers.GameEndHandler{}.Parse(&dec, &d)
			tc.errAssertion(t, err)

			actual := d.GameEnd
			require.Equal(t, tc.expected, actual)
		})
	}
}

func buildGameEbdInput() []byte {
	out := []byte{byte(event.EventGameEnd)}

	out = append(out, byte(gameEndMethod))
	out = append(out, byte(gameEndLRAS))
	out = append(out, byte(gameEndPlacement0))
	out = append(out, byte(gameEndPlacement1))
	out = append(out, byte(gameEndPlacement2))
	out = append(out, byte(gameEndPlacement3))

	return out
}
