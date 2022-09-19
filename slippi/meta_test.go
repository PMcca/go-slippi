package slippi_test

import (
	"bytes"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/jmank88/ubjson"
	"github.com/stretchr/testify/require"
	"testing"
)

const metaUBJSON = `[{][i][i:7][S:startAt][S][i][i:21][S:2022-08-28T15:51:13ZU][i][i:9][S:lastFrame][I][I:2011][i][i:7][S:players][{][i][i:1][S:0][{][i][i:10][S:characters][{][i][i:1][S:4][I][I:567][i][i:1][S:5][i][i:77][}][i][i:5][S:names][{][i][i:7][S:netplay][S][i][i:12][S:netplay-name][i][i:4][S:code][S][i][i:8][S:TEST#001][}][}][i][i:1][S:1][{][i][i:10][S:characters][{][i][i:1][S:1][i][i:123][}][}][}][}]`

func TestUnmarshalUBJSON(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		X            string
		ErrAssertion require.ErrorAssertionFunc
	}{
		"": {},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			d := ubjson.NewDecoder(bytes.NewReader([]byte(metaUBJSON)))
			meta := slippi.Metadata{}
			err := meta.UnmarshalUBJSON(d)

			tc.ErrAssertion(t, err)

		})
	}
}
