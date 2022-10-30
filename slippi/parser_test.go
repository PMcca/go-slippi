package slippi_test

import (
	"github.com/PMcca/go-slippi/internal/testutil"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/melee"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseMeta(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		filePath     string
		expected     *slippi.Metadata
		errAssertion require.ErrorAssertionFunc
	}{
		"EmptyFilePathReturnsError": {
			filePath:     "",
			expected:     nil,
			errAssertion: testutil.IsError(slippi.ErrEmptyFilePath),
		},
		"ErrorReadingFileReturnsError": {
			filePath:     "some-non-existent-file-iodsjfoisdnhgs.slp",
			expected:     nil,
			errAssertion: testutil.IsError(slippi.ErrOpeningFile),
		},
		"ErrorParsingMetaReturnsError": {
			filePath:     "testdata/invalid-meta.ubj",
			expected:     nil,
			errAssertion: testutil.IsError(slippi.ErrParsingMeta),
		},
		"ParsesAndReturnsMeta": {
			filePath: "testdata/valid-meta-1-player.ubj",
			expected: &slippi.Metadata{
				StartAt:   "2022-08-28T15:51:13ZU",
				LastFrame: 3000,
				Players: []*slippi.Player{
					{
						Name: slippi.Names{
							Name:       "name",
							SlippiCode: "TEST#001",
						},
						Port: 0,
						Characters: []slippi.Character{
							{
								CharacterID:  melee.Mario,
								FramesPlayed: 3000,
							},
						},
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

			actual, err := slippi.ParseMeta(tc.filePath)

			tc.errAssertion(t, err)
			require.Equal(t, tc.expected, actual)

		})
	}
}
