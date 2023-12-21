package handlers_test

import (
	"encoding/binary"
	"github.com/PMcca/go-slippi/internal/testutil"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/event"
	"github.com/PMcca/go-slippi/slippi/event/handler/handlers"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseMessageSplitter(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		inputConfigureFunc func(b []byte)
		expected           slippi.Data
		errAssertion       require.ErrorAssertionFunc
	}{
		"UnknownInternalEventCodeReturnsError": {
			inputConfigureFunc: func(b []byte) {
				b[515] = 1 // Change internal command code to something unknown
				b[517+515] = 1
			},
			expected:     slippi.Data{},
			errAssertion: testutil.IsError(handlers.ErrUnknownMessageSplitEvent),
		},
		"IncorrectMessageSplitterCodeReturnsError": {
			inputConfigureFunc: func(b []byte) {
				b[517] = 1 // Second MessageSplitter event has incorrect code
			},
			expected:     slippi.Data{},
			errAssertion: testutil.IsError(handlers.ErrNoMessageSplitterCode),
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			input := buildMessageSplitterInput()
			if fn := tc.inputConfigureFunc; fn != nil {
				fn(input)
			}

			dec := event.Decoder{
				Data: input,
				Size: len(input),
			}
			actual := slippi.Data{}

			err := handlers.MessageSplitterHandler{}.Parse(&dec, &actual)

			tc.errAssertion(t, err)
			require.Equal(t, tc.expected, actual)

		})
	}
}

// buildMessageSplitterInput returns two message splitter events, one with 512 bytes and another with 200 bytes in their respective
// payloads, for a GeckoCode event. It can be mutated by individual tests.
func buildMessageSplitterInput() []byte {
	input := make([]byte, 1034)
	// Max, 512 bytes for first payload
	for i := 1; i < 513; i++ {
		input[i] = byte(i + 1) // Generate dummy payload data for first split event
	}
	// 200 bytes for second payload
	for i := 518; i < 718; i++ {
		input[i] = byte(i + 1)
	}

	// Put Command Code as first byte
	input[0] = byte(event.EventMessageSplitter)
	input[517] = byte(event.EventMessageSplitter)

	// Build size of first payload
	sizeBuf := make([]byte, 2)
	binary.BigEndian.PutUint16(sizeBuf, 512)
	input[513] = sizeBuf[0]
	input[514] = sizeBuf[1]

	// Build size of second payload
	sizeBuf = make([]byte, 2)
	binary.BigEndian.PutUint16(sizeBuf, 200)
	input[517+513] = sizeBuf[0]
	input[517+514] = sizeBuf[1]

	// Put command code that these message splitter events are for
	input[515] = byte(event.EventGeckoList)
	input[517+515] = byte(event.EventGeckoList)

	// Finally, set "isFinalMessage" field. First payload is 0, second is 1.
	input[516] = 0
	input[517+516] = 1

	return input
}
