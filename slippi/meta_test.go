package slippi_test

import (
	"bytes"
	"encoding/hex"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/jmank88/ubjson"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

// This is hacky, but represents valid UBJSON.
const metadataUBJSON = "7B690773746172744174536915323032322D30382D32385431353A35313A31335A5569096C6173744672616D656C000100016907706C61796572737B6901307B69056E616D65737B69076E6574706C61795369074D616E79756C616904636F64655369084D414E59233434347D690A636861726163746572737B6901326C000100017D7D6901317B69056E616D65737B69076E6574706C61795369044A6F686E6904636F64655369064A41592334347D690A636861726163746572737B6901336C000100017D7D7D7D"

var as = []byte{123, 105, 7, 115, 116, 97, 114, 116, 65, 116, 83, 105, 21, 50, 48, 50, 50, 45, 48, 56, 45, 50, 56, 84, 49, 53, 58, 53, 49, 58, 49, 51, 90, 85, 105, 9, 108, 97, 115, 116, 70, 114, 97, 109, 101, 108, 0, 1, 0, 1, 105, 7, 112, 108, 97, 121, 101, 114, 115, 123, 105, 1, 48, 123, 105, 5, 110, 97, 109, 101, 115, 123, 105, 7, 110, 101, 116, 112, 108, 97, 121, 83, 105, 7, 77, 97, 110, 121, 117, 108, 97, 105, 4, 99, 111, 100, 101, 83, 105, 8, 77, 65, 78, 89, 35, 52, 52, 52, 125, 105, 10, 99, 104, 97, 114, 97, 99, 116, 101, 114, 115, 123, 105, 1, 50, 108, 0, 1, 0, 1, 125, 125, 105, 1, 49, 123, 105, 5, 110, 97, 109, 101, 115, 123, 105, 7, 110, 101, 116, 112, 108, 97, 121, 83, 105, 4, 74, 111, 104, 110, 105, 4, 99, 111, 100, 101, 83, 105, 6, 74, 65, 89, 35, 52, 52, 125, 105, 10, 99, 104, 97, 114, 97, 99, 116, 101, 114, 115, 123, 105, 1, 51, 108, 0, 1, 0, 1, 125, 125, 125, 125}

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

			b, err := hex.DecodeString(metadataUBJSON)
			if err != nil {
				log.Fatal(err)
			}

			d := ubjson.NewDecoder(bytes.NewReader(as))
			meta := slippi.Metadata{}
			if er := ubjson.Unmarshal(b, &meta); err != nil {
				log.Fatal(er)
			}
			err = meta.UnmarshalUBJSON(d)

			tc.ErrAssertion(t, err)

		})
	}
}
