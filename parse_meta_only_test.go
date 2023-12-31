package goslippi

import (
	"github.com/pmcca/go-slippi/internal/testutil"
	"github.com/pmcca/go-slippi/slippi"
	"github.com/pmcca/go-slippi/slippi/melee"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseMeta(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		filePath     string
		expected     slippi.Metadata
		errAssertion require.ErrorAssertionFunc
	}{
		"EmptyFilePathReturnsError": {
			filePath:     "",
			expected:     slippi.Metadata{},
			errAssertion: testutil.IsError(ErrEmptyFilePath),
		},
		"ErrorReadingFileReturnsError": {
			filePath:     "some-non-existent-file.slp",
			expected:     slippi.Metadata{},
			errAssertion: testutil.IsError(ErrReadingFile),
		},
		"ErrorParsingMetaReturnsError": {
			filePath:     "test/replays/invalid-ubjson.ubj",
			expected:     slippi.Metadata{},
			errAssertion: testutil.IsError(ErrParsingMeta),
		},
		"ParsesAndReturnsMeta": {
			filePath: "test/replays/metadata.slp",
			expected: slippi.Metadata{
				StartAt:   "2022-12-02T18:09:00Z",
				LastFrame: 2011,
				Players: slippi.PlayersMeta{
					Port1: slippi.PlayerMeta{
						Names: slippi.Names{
							Name:       "Smasher",
							SlippiCode: "SMSH#123",
						},
						Characters: slippi.Characters{
							{
								CharacterID:  melee.Int_Zelda,
								FramesPlayed: 532,
							},
							{
								CharacterID:  melee.Int_Sheik,
								FramesPlayed: 1603,
							},
						},
					},
					Port2: slippi.PlayerMeta{
						Names: slippi.Names{
							Name:       "I Love Slippi!",
							SlippiCode: "SLIP#987",
						},
						Characters: slippi.Characters{
							{
								CharacterID:  melee.Int_Mario,
								FramesPlayed: 2135,
							},
						},
					},
				},
				PlayedOn: "dolphin",
			},
			errAssertion: require.NoError,
		},
	}

	for name, testCase := range testCases {
		tc := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := ParseMeta(tc.filePath)
			tc.errAssertion(t, err)
			require.Equal(t, tc.expected, actual)
		})
	}
}
