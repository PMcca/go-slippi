package goslippi

import (
	"github.com/pmcca/go-slippi/internal/testutil"
	"github.com/pmcca/go-slippi/slippi"
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
