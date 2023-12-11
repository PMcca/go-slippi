package slippi_test

import (
	"encoding/binary"
	"github.com/PMcca/go-slippi/internal/testutil"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/stretchr/testify/require"
	"testing"
)

var rawStart = []byte("$U#l")

func TestData_UnmarshalUBJSON(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input        []byte
		expected     slippi.Data
		errAssertion require.ErrorAssertionFunc
	}{
		"InvalidStartOfRawReturnsError": {
			input:        []byte("this is invalid"),
			expected:     slippi.Data{},
			errAssertion: testutil.IsError(slippi.ErrInvalidRawStart),
		},
		"UnknownEventInPayloadSizesReturnsError": {
			input: append(buildInput(6), []byte{
				0x35, // EventPayloads event
				4,    // Size of eventPayloads
				0x36, // GameStart event
				2, 3, // Size of GameStartEvent
				0x99, // Unknown byte, should break parsing
			}...),
			expected:     slippi.Data{},
			errAssertion: testutil.IsError(slippi.ErrUnknownEventInEventSizes),
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			d := slippi.Data{}
			err := d.UnmarshalUBJSON(tc.input)

			tc.errAssertion(t, err)
			require.Equal(t, tc.expected, d)
		})
	}
}

func buildInput(totalSize int) []byte {
	r := make([]byte, 4)
	copy(r, rawStart)

	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(totalSize))

	return append(r, b...)
}
