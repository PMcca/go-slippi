package testutil

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"path/filepath"
	"testing"
)

// ReadUBJSONFile takes a filename and returns the bytes read from the given text fixture.
func ReadUBJSONFile(name string) ([]byte, error) {
	f, err := os.Open(filepath.Join("testdata", name))
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// IsError returns an ErrorAssertionFunc checking if the returned error is type-equal to the given error.
func IsError(e error) require.ErrorAssertionFunc {
	return func(t require.TestingT, err error, i ...interface{}) {
		require.Error(t, err)

		if !assert.True(t, errors.Is(err, e)) {
			if tt, ok := t.(*testing.T); ok {
				tt.Logf("Incorrect error type for error '%s'. Expected %T, got %T", err, e, err)
			}
		}
	}
}
