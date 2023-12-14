package slippi

import (
	"encoding/binary"
	"math"
)

// decoder wraps the raw data of a .slp file. It serves as a window into a specific part of the array, which is then
// acted upon by the corresponding offsets defined in the .slp spec (https://github.com/PMcca/slippi-wiki/blob/master/SPEC.md)
type decoder struct {
	data []byte
	size int // size is the size of the event being parsed, used for bounds checking.
}

// read returns a byte at the given offset.
func (d *decoder) read(offset int) uint8 {
	if offset > d.size {
		return 0
	}
	return d.data[offset]
}

// readN returns a slice of bytes between the given offset and the upperBound.
func (d *decoder) readN(offset, upperBound int) []byte {
	if upperBound > d.size {
		return nil
	}
	return d.data[offset:upperBound]
}

// readWithBitmask returns a byte bitwise-ANDed against the given bitmask.
func (d *decoder) readWithBitmask(offset int, bitmask byte) byte {
	if offset > d.size {
		return 0
	}
	return d.data[offset] & bitmask
}

// readInt16 returns an int from the 2 bytes from the offset the decoder assumes represents a uint16.
func (d *decoder) readInt16(offset int) int {
	if offset+2 > d.size {
		return 0
	}
	return int(int16(binary.BigEndian.Uint16(d.data[offset : offset+2])))
}

// readInt32 returns an int from the 4 bytes from the offset the decoder assumes represents a uint32.
func (d *decoder) readInt32(offset int) int {
	if offset+4 > d.size {
		return 0
	}
	return int(binary.BigEndian.Uint32(d.data[offset : offset+4]))
}

func (d *decoder) readBool(offset int) bool {
	if offset > d.size {
		return false
	}
	return d.data[offset] > 0

}

// readFloat32 returns an IEEE-754 32 bit floating point number from the given offset.
func (d *decoder) readFloat32(offset int) float32 {
	if offset+4 > d.size {
		return 0
	}
	nums := d.data[offset : offset+4]
	// Combine the 4 bytes into a single uint32 number
	combined := uint32(nums[0])
	combined <<= 8
	combined += uint32(nums[1])
	combined <<= 8
	combined += uint32(nums[2])
	combined <<= 8
	combined += uint32(nums[3])

	return math.Float32frombits(combined)
}
