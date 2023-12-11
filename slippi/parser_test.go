package slippi_test

import (
	"fmt"
	"github.com/PMcca/go-slippi/internal/testutil"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/melee"
	"github.com/stretchr/testify/require"
	"testing"
)

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
			errAssertion: testutil.IsError(slippi.ErrEmptyFilePath),
		},
		"ErrorReadingFileReturnsError": {
			filePath:     "some-non-existent-file.slp",
			expected:     slippi.Game{},
			errAssertion: testutil.IsError(slippi.ErrReadingFile),
		},
		"ErrorParsingMetaReturnsError": {
			filePath:     "test/replays/invalid-ubjson.ubj",
			expected:     slippi.Game{},
			errAssertion: testutil.IsError(slippi.ErrParsingGame),
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := slippi.ParseGame(tc.filePath)

			tc.errAssertion(t, err)
			require.Equal(t, tc.expected, actual)

		})
	}
}

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
			errAssertion: testutil.IsError(slippi.ErrEmptyFilePath),
		},
		"ErrorReadingFileReturnsError": {
			filePath:     "some-non-existent-file.slp",
			expected:     slippi.Metadata{},
			errAssertion: testutil.IsError(slippi.ErrReadingFile),
		},
		"ErrorParsingMetaReturnsError": {
			filePath:     "test/replays/invalid-ubjson.ubj",
			expected:     slippi.Metadata{},
			errAssertion: testutil.IsError(slippi.ErrParsingMeta),
		},
		"ParsesAndReturnsMeta": {
			filePath: "test/replays/metadata.slp",
			//filePath: "testdata/valid-meta-1-player.ubj",
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
								CharacterID:  melee.Zelda,
								FramesPlayed: 532,
							},
							{
								CharacterID:  melee.Sheik,
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
								CharacterID:  melee.Mario,
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

			actual, err := slippi.ParseMeta(tc.filePath)

			tc.errAssertion(t, err)
			require.Equal(t, tc.expected, actual)

		})
	}
}

func TestParse(t *testing.T) {
	//t.SkipNow()
	g, err := slippi.ParseGame("test/replays/20221202T180900.slp")
	//g, err := slippi.ParseGame("test/replays/metadata.slp")
	require.NoError(t, err)
	fmt.Println(g)
}
