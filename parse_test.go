package goslippi

import (
	"encoding/binary"
	"github.com/PMcca/go-slippi/internal/testutil"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/stretchr/testify/require"
	"testing"
)

var rawStart = []byte("$U#l")

func TestParseGame(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		filePath     string
		expected     slippi.Game
		errAssertion require.ErrorAssertionFunc
	}{
		"EmptyFilePathReturnsError": {
			filePath:     "",
			expected:     slippi.Game{},
			errAssertion: testutil.IsError(ErrEmptyFilePath),
		},
		"ErrorReadingFileReturnsError": {
			filePath:     "some-non-existent-file.slp",
			expected:     slippi.Game{},
			errAssertion: testutil.IsError(ErrReadingFile),
		},
		"ErrorParsingMetaReturnsError": {
			filePath:     "test/replays/invalid-ubjson.ubj",
			expected:     slippi.Game{},
			errAssertion: testutil.IsError(ErrParsingGame),
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := ParseGame(tc.filePath)

			tc.errAssertion(t, err)
			require.Equal(t, tc.expected, actual)

		})
	}
}

func TestData_UnmarshalUBJSON(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input        []byte
		expected     rawParser
		errAssertion require.ErrorAssertionFunc
	}{
		"InvalidStartOfRawReturnsError": {
			input:        []byte("this is invalid"),
			expected:     rawParser{},
			errAssertion: testutil.IsError(ErrInvalidRawStart),
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			d := rawParser{}
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

func TestParse(t *testing.T) {
	t.SkipNow()
	//g, err := slippi.ParseGame("test/replays/nametags.slp")
	//g, err := slippi.ParseGame("test/replays/ranked_game1_tiebreak.slp")
	x, err := ParseGame("test/replays/20221202T180900.slp")
	x = x
	//g, err := slippi.ParseGame("test/replays/metadata.slp")
	require.NoError(t, err)
	//fmt.Println(g)
}
