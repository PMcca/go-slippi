package handlers_test

import (
	"github.com/pmcca/go-slippi/internal/testutil"
	"github.com/pmcca/go-slippi/slippi/event"
	"github.com/pmcca/go-slippi/slippi/event/handler/handlers"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseEventPayloads(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		input        []byte
		expected     map[event.Code]int
		errAssertion require.ErrorAssertionFunc
	}{
		"FirstByteNotEventPayloadsCommandReturnsError": {
			input:        []byte("this will break"),
			expected:     nil,
			errAssertion: testutil.IsError(handlers.ErrEventPayloadsNotFound),
		},
		"InvalidNumberOfEventsReturnsError": {
			input:        []byte{byte(event.EventPayloadsEvent), 6}, // Payload size = 6 (-1) which is not divisible by 3
			expected:     nil,
			errAssertion: testutil.IsError(handlers.ErrInvalidNumberOfCommands),
		},
		"ReturnsSizesOfEvents": {
			input: []byte{
				byte(event.EventPayloadsEvent),
				7, // Size = this + the following 6 bytes = 7
				byte(event.EventGameStart),
				2,
				3, // size is uint16 so 2 bytes
				byte(event.EventFrameStart),
				4, 5,
			},
			expected: map[event.Code]int{
				event.EventPayloadsEvent: 7,
				event.EventGameStart:     515,
				event.EventFrameStart:    1029,
			},
			errAssertion: require.NoError,
		},
	}

	for name, testCase := range testCases {
		tc := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			dec := event.Decoder{
				Data: tc.input,
				Size: len(tc.input),
			}
			actual, err := handlers.ParseEventPayloads(&dec)

			tc.errAssertion(t, err)
			require.Equal(t, tc.expected, actual)
		})
	}
}
