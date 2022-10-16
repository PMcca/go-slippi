package slippi_test

import (
	"github.com/PMcca/go-slippi/internal/testutil"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/melee"
	"github.com/jmank88/ubjson"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnmarshalUBJSON(t *testing.T) {

	// See testdata/README for description of test data used.
	parsedMeta := slippi.Metadata{
		StartAt:   "2022-08-28T15:51:13ZU",
		LastFrame: 3000,
		Players: []*slippi.Player{
			{
				Name: slippi.Names{
					Name:       "Netplay Name",
					SlippiCode: "TEST#001",
				},
				Port: 0,
				Characters: []slippi.Character{
					{
						CharacterID:  melee.Sheik,
						FramesPlayed: 2800,
					},
					{
						CharacterID:  melee.Zelda,
						FramesPlayed: 200,
					},
				},
			},
			{
				Name: slippi.Names{
					Name:       "Netplay Name 2",
					SlippiCode: "TEST#002",
				},
				Port: 1,
				Characters: []slippi.Character{
					{
						CharacterID:  melee.DonkeyKong,
						FramesPlayed: 3000,
					},
				},
			},
		},
	}

	t.Parallel()
	testCases := map[string]struct {
		fixtureFilename string
		expected        slippi.Metadata
		errAssertion    require.ErrorAssertionFunc
	}{
		"UnmarshalsValidMetadataUBJSONIntoStruct": {
			fixtureFilename: "valid-meta.ubj",
			expected:        parsedMeta,
			errAssertion:    require.NoError,
		},
		"InvalidFieldTypeReturnsError": {
			fixtureFilename: "invalid-lastFrame.ubj",
			expected:        slippi.Metadata{},
			errAssertion:    testutil.IsError(slippi.ErrDecodingField),
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			tcBytes, err := testutil.ReadUBJSONFile(tc.fixtureFilename)
			require.NoError(t, err)

			actual := slippi.Metadata{}
			err = ubjson.Unmarshal(tcBytes, &actual)

			tc.errAssertion(t, err)

			require.Equal(t, tc.expected, actual)

		})
	}
}
