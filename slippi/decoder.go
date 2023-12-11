package slippi

import "encoding/binary"

// decoder wraps the raw data of a .slp file. It serves as a window into a specific part of the array, which is then
// acted upon by the corresponding offsets defined in the .slp spec (https://github.com/PMcca/slippi-wiki/blob/master/SPEC.md)
type decoder struct {
	data []byte
}

// read returns a byte at the given offset.
func (d *decoder) read(offset int) byte {
	return d.data[offset]
}

// readWithBitmask returns a byte bitwise-ANDed against the given bitmask.
func (d *decoder) readWithBitmask(offset int, bitmask byte) byte {
	return d.data[offset] & bitmask
}

// readInt16 returns an int from the 2 bytes from the offset the decoder assumes represents a uint16.
func (d *decoder) readInt16(offset int) int {
	return int(int16(binary.BigEndian.Uint16(d.data[offset:2])))
}

// readInt32 returns an int from the 4 bytes from the offset the decoder assumes represents a uint32.
func (d *decoder) readInt32(offset int) int {
	return int(binary.BigEndian.Uint32(d.data[offset:4]))
}

func (d *decoder) readBool(offset int) bool {
	if d.data[offset] == 'T' {
		return true
	}
	return false
}
