package event_test

import (
	"encoding/binary"
	"github.com/pmcca/go-slippi/internal/testutil"
	"github.com/pmcca/go-slippi/slippi/event"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRead(t *testing.T) {
	t.Parallel()
	t.Run("ReadsSingleByte", func(t *testing.T) {
		t.Parallel()
		expected := byte(5)
		data := []byte{expected}
		dec := event.Decoder{
			Data: data,
			Size: 1,
		}

		actual := dec.Read(0)
		require.Equal(t, expected, actual)
	})

	t.Run("OffsetLargerThanSizeReturns0", func(t *testing.T) {
		t.Parallel()
		data := []byte{5}
		dec := event.Decoder{
			Data: data,
			Size: 1,
		}

		expected := byte(0)
		actual := dec.Read(100)
		require.Equal(t, expected, actual)
	})
}

func TestReadInt8(t *testing.T) {
	t.Parallel()
	t.Run("ReadsInt8", func(t *testing.T) {
		t.Parallel()
		data := []byte{5}
		dec := event.Decoder{
			Data: data,
			Size: 1,
		}

		expected := int8(5)
		actual := dec.ReadInt8(0)
		require.Equal(t, expected, actual)
	})

	t.Run("OffsetLargerThanSizeReturns0", func(t *testing.T) {
		t.Parallel()
		data := []byte{5}
		dec := event.Decoder{
			Data: data,
			Size: 1,
		}

		expected := int8(0)
		actual := dec.ReadInt8(100)
		require.Equal(t, expected, actual)
	})
}

func TestReadN(t *testing.T) {
	t.Parallel()
	t.Run("ReadsN-number of bytes", func(t *testing.T) {
		t.Parallel()
		data := []byte{5, 6, 7, 8, 9}
		dec := event.Decoder{
			Data: data,
			Size: 5,
		}

		expected := []byte{5, 6, 7, 8, 9}
		actual := dec.ReadN(0, 5)
		require.Equal(t, expected, actual)
	})

	t.Run("OffsetLargerThanSizeReturnsNil", func(t *testing.T) {
		t.Parallel()
		data := []byte{5, 6, 7, 8, 9}
		dec := event.Decoder{
			Data: data,
			Size: 5,
		}

		var expected []byte = nil
		actual := dec.ReadN(6, 10)
		require.Equal(t, expected, actual)
	})

	t.Run("UpperBoundLargerThanSizeReturnsNil", func(t *testing.T) {
		t.Parallel()
		data := []byte{5, 6, 7, 8, 9}
		dec := event.Decoder{
			Data: data,
			Size: 5,
		}

		var expected []byte = nil
		actual := dec.ReadN(0, 10)
		require.Equal(t, expected, actual)
	})

	t.Run("OffsetLargerThanUpperBoundReturnsNil", func(t *testing.T) {
		t.Parallel()
		data := []byte{5, 6, 7, 8, 9}
		dec := event.Decoder{
			Data: data,
			Size: 5,
		}

		var expected []byte = nil
		actual := dec.ReadN(3, 2)
		require.Equal(t, expected, actual)
	})
}

func TestReadWithBitmask(t *testing.T) {
	t.Parallel()
	t.Run("ReadsWithBitmask", func(t *testing.T) {
		t.Parallel()
		data := []byte{0b11111111}
		bitMask := byte(0b00000011)
		dec := event.Decoder{
			Data: data,
			Size: 1,
		}

		expected := byte(0b00000011)
		actual := dec.ReadWithBitmask(0, bitMask)
		require.Equal(t, expected, actual)
	})

	t.Run("OffsetLargerThanSizeReturns0", func(t *testing.T) {
		t.Parallel()
		data := []byte{0b11111111}
		bitMask := byte(0b00000011)
		dec := event.Decoder{
			Data: data,
			Size: 1,
		}

		expected := byte(0)
		actual := dec.ReadWithBitmask(10, bitMask)
		require.Equal(t, expected, actual)
	})
}

func TestReadUint16(t *testing.T) {
	t.Parallel()
	t.Run("ReadsUint16", func(t *testing.T) {
		t.Parallel()
		data := make([]byte, 2)
		binary.BigEndian.PutUint16(data, 5)
		dec := event.Decoder{
			Data: data,
			Size: 2,
		}

		expected := uint16(5)
		actual := dec.ReadUint16(0)
		require.Equal(t, expected, actual)
	})

	t.Run("OffsetLargerThanSizeReturns0", func(t *testing.T) {
		t.Parallel()
		data := make([]byte, 2)
		binary.BigEndian.PutUint16(data, 5)
		dec := event.Decoder{
			Data: data,
			Size: 2,
		}

		expected := uint16(0)
		actual := dec.ReadUint16(4)
		require.Equal(t, expected, actual)
	})
}

func TestReadInt16(t *testing.T) {
	t.Parallel()
	t.Run("ReadsInt16", func(t *testing.T) {
		t.Parallel()
		data := make([]byte, 2)
		input := int16(-12)
		binary.BigEndian.PutUint16(data, uint16(input))
		dec := event.Decoder{
			Data: data,
			Size: 2,
		}

		expected := -12
		actual := dec.ReadInt16(0)
		require.Equal(t, expected, actual)
	})

	t.Run("OffsetLargerThanSizeReturns0", func(t *testing.T) {
		t.Parallel()
		data := make([]byte, 2)
		input := int16(-12)
		binary.BigEndian.PutUint16(data, uint16(input))
		dec := event.Decoder{
			Data: data,
			Size: 2,
		}

		expected := 0
		actual := dec.ReadInt16(5)
		require.Equal(t, expected, actual)
	})
}

func TestReadUint32(t *testing.T) {
	t.Parallel()
	t.Run("ReadsUint32", func(t *testing.T) {
		t.Parallel()
		data := make([]byte, 4)
		binary.BigEndian.PutUint32(data, 5)
		dec := event.Decoder{
			Data: data,
			Size: 4,
		}

		expected := uint32(5)
		actual := dec.ReadUint32(0)
		require.Equal(t, expected, actual)
	})

	t.Run("OffsetLargerThanSizeReturns0", func(t *testing.T) {
		t.Parallel()
		data := make([]byte, 4)
		binary.BigEndian.PutUint32(data, 5)
		dec := event.Decoder{
			Data: data,
			Size: 4,
		}

		expected := uint32(0)
		actual := dec.ReadUint32(6)
		require.Equal(t, expected, actual)
	})
}

func TestReadInt32(t *testing.T) {
	t.Parallel()
	t.Run("ReadsInt32", func(t *testing.T) {
		t.Parallel()
		data := make([]byte, 4)
		input := int32(-360)
		binary.BigEndian.PutUint32(data, uint32(input))
		dec := event.Decoder{
			Data: data,
			Size: 4,
		}

		expected := -360
		actual := dec.ReadInt32(0)
		require.Equal(t, expected, actual)
	})

	t.Run("OffsetLargerThanSizeReturns0", func(t *testing.T) {
		t.Parallel()
		data := make([]byte, 4)
		input := int32(-360)
		binary.BigEndian.PutUint32(data, uint32(input))
		dec := event.Decoder{
			Data: data,
			Size: 4,
		}

		expected := 0
		actual := dec.ReadInt32(6)
		require.Equal(t, expected, actual)
	})
}

func TestReadBool(t *testing.T) {
	t.Parallel()
	t.Run("ReadsBool", func(t *testing.T) {
		t.Parallel()
		data := []byte{1, 0} // True and False
		dec := event.Decoder{
			Data: data,
			Size: 2,
		}

		actual := dec.ReadBool(0)
		require.True(t, actual)

		actual = dec.ReadBool(1)
		require.False(t, actual)
	})

	t.Run("OffsetLargerThanSizeReturnsFalse", func(t *testing.T) {
		t.Parallel()
		data := []byte{1, 0} // True and False
		dec := event.Decoder{
			Data: data,
			Size: 2,
		}

		actual := dec.ReadBool(5)
		require.False(t, actual)
	})
}

func TestReadFloat32(t *testing.T) {
	t.Parallel()
	t.Run("ReadsFloat32", func(t *testing.T) {
		t.Parallel()
		data := make([]byte, 0)
		testutil.PutFloat32(&data, -64.25)
		dec := event.Decoder{
			Data: data,
			Size: 4,
		}

		expected := float32(-64.25)
		actual := dec.ReadFloat32(0)
		require.Equal(t, expected, actual)
	})

	t.Run("OffsetLargerThanSizeReturns0", func(t *testing.T) {
		t.Parallel()
		data := make([]byte, 0)
		testutil.PutFloat32(&data, -64.25)
		dec := event.Decoder{
			Data: data,
			Size: 4,
		}

		expected := float32(0)
		actual := dec.ReadFloat32(8)
		require.Equal(t, expected, actual)
	})
}
