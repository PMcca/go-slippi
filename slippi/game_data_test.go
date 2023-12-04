package slippi_test

import (
	"github.com/PMcca/go-slippi/internal/testutil"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/stretchr/testify/require"
	"testing"
)

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
