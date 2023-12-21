package event

import (
	"encoding/binary"
	"math"
)

// TODO add tests

// Decoder wraps the raw data of a .slp file. It serves as a window into a specific part of the array, which is then
// acted upon by the corresponding offsets defined in the .slp spec (https://github.com/project-slippi/slippi-wiki/blob/master/SPEC.md)
// Decoder.size is the eventSize+1, because the size of the data is a valid index. e.g:
//   - offset = 97
//   - size = 100
//
// When Reading a uint32, we will read bytes [97, 98, 99, 100], so the size must be +1 to allow for this.
type Decoder struct {
	Data []byte
	Size int // size is the size of the event being parsed, used for bounds checking.
}

// Read returns a byte at the given offset.
func (d *Decoder) Read(offset int) uint8 {
	if offset >= d.Size {
		return 0
	}
	return d.Data[offset]
}

// ReadInt8 returns an int8 at the given offset.
func (d *Decoder) ReadInt8(offset int) int8 {
	if offset >= d.Size {
		return 0
	}
	return int8(d.Data[offset])
}

// ReadN returns a slice of bytes between the given offset and the upperBound.
func (d *Decoder) ReadN(offset, upperBound int) []byte {
	if offset >= d.Size || upperBound > d.Size {
		return nil
	}
	return d.Data[offset:upperBound]
}

// ReadWithBitmask returns a byte bitwise-ANDed against the given bitmask.
func (d *Decoder) ReadWithBitmask(offset int, bitmask byte) byte {
	if offset >= d.Size {
		return 0
	}
	return d.Data[offset] & bitmask
}

// ReadUint16 returns an int from the 2 bytes from the offset the Decoder assumes represents a uint16.
func (d *Decoder) ReadUint16(offset int) uint16 {
	if offset >= d.Size || offset+2 > d.Size {
		return 0
	}
	return binary.BigEndian.Uint16(d.Data[offset : offset+2])
}

// ReadInt16 returns an int from the 2 bytes from the offset the Decoder assumes represents an int16.
func (d *Decoder) ReadInt16(offset int) int {
	if offset >= d.Size || offset+2 > d.Size {
		return 0
	}
	return int(int16(binary.BigEndian.Uint16(d.Data[offset : offset+2])))
}

// ReadUint32 returns a uint32 from the 4 bytes from the offset the Decoder assumes represents a uint32.
func (d *Decoder) ReadUint32(offset int) uint32 {
	if offset+4 > d.Size {
		return 0
	}
	return binary.BigEndian.Uint32(d.Data[offset : offset+4])
}

// ReadInt32 returns an int from the 4 bytes from the offset the Decoder assumes represents an int32.
func (d *Decoder) ReadInt32(offset int) int {
	if offset+4 > d.Size {
		return 0
	}
	return int(int32(binary.BigEndian.Uint32(d.Data[offset : offset+4])))
}

// ReadBool returns a bool from the byte from the offset.
func (d *Decoder) ReadBool(offset int) bool {
	if offset > d.Size {
		return false
	}
	return d.Data[offset] > 0

}

// ReadFloat32 returns an IEEE-754 32 bit floating point number from the given offset.
func (d *Decoder) ReadFloat32(offset int) float32 {
	if offset+4 > d.Size {
		return 0
	}
	f := d.ReadUint32(offset)
	return math.Float32frombits(f)
}
