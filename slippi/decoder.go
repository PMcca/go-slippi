package slippi

// decoder wraps the raw data of a .slp file and adds QoL functionality to reading it.
type decoder struct {
	data   []byte
	offset int
}

// read returns the next byte read from data.
func (d *decoder) read() byte {
	b := d.data[d.offset]
	d.offset++
	return b
}

// readN returns n-number of bytes from data.
func (d *decoder) readN(n int) []byte {
	e := d.offset + n
	b := d.data[d.offset:e]
	d.offset += n
	return b
}
