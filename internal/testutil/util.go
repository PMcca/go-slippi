package testutil

import (
	"encoding/binary"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

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

// PutInt32 takes an int32 and puts its bytes into the input array.
func PutInt32(input *[]byte, num int32) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(num))
	*input = append(*input, buf...)
}

// PutUint32 takes a uint32 and puts its bytes into the input array.
func PutUint32(input *[]byte, num uint32) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, num)
	*input = append(*input, buf...)
}

// PutInt16 takes an int16 and puts its bytes into the input array.
func PutInt16(input *[]byte, num int16) {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(num))
	*input = append(*input, buf...)
}

// PutUint16 takes a uint16 and puts its bytes into the input array.
func PutUint16(input *[]byte, num uint16) {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, num)
	*input = append(*input, buf...)
}

// PutFloat32 takes a float32 and puts its bytes into the input array.
func PutFloat32(input *[]byte, num float32) {
	buf := make([]byte, 4)
	f := math.Float32bits(num)
	binary.BigEndian.PutUint32(buf, f)
	*input = append(*input, buf...)

}
