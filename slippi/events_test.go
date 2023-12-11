package slippi

import (
	"github.com/PMcca/go-slippi/internal/testutil"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_parseEventPayloads(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input        []byte
		expected     map[eventType]int
		errAssertion require.ErrorAssertionFunc
	}{
		"FirstByteNotEventPayloadsCommandReturnsError": {
			input:        []byte("this will break"),
			expected:     nil,
			errAssertion: testutil.IsError(ErrEventPayloadsNotFound),
		},
		"InvalidNumberOfEventsReturnsError": {
			input:        []byte{byte(eventPayloadsEvent), 6}, // Payload size = 6 (-1) which is not divisible by 3
			expected:     nil,
			errAssertion: testutil.IsError(ErrInvalidNumberOfCommands),
		},
		"ReturnsSizesOfEvents": {
			input: []byte{
				byte(eventPayloadsEvent),
				7, // Size of eventPayloads = 4-1 = 3
				byte(eventGameStart),
				2,
				3, // size is uint16 so 2 bytes
				byte(eventFrameStart),
				4, 5,
			},
			expected: map[eventType]int{
				eventGameStart:  515,
				eventFrameStart: 1029,
			},
			errAssertion: require.NoError,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := parseEventPayloads(tc.input)

			tc.errAssertion(t, err)
			require.Equal(t, tc.expected, actual)
		})
	}
}

func Test_parseGameStart(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input        []byte
		expected     *Data
		errAssertion require.ErrorAssertionFunc
	}{
		"ParsesGameStartEvent": {
			input: []byte{
				3, 14, 0, 0, // Slippi semver. Forth byte unused.
			},
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			d := decoder{
				data: tc.input,
			}

			actual := Data{}
			err := parseGameStart(len(tc.input), &d, &actual)

			tc.errAssertion(t, err)
			require.Equal(t, tc.expected.GameStart, actual.GameStart)
		})
	}
}
