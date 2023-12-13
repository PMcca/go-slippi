package util

// ToHalfWidthChars takes a Shift JIS-decoded string and converts specific bytes to half-width.
func ToHalfWidthChars(s string) string {
	var ret []rune
	for _, c := range s {
		switch {
		case c > 0xff00 && c < 0xff5f:
			ret = append(ret, 0x0020+(c-0xff00))
		case c == 0x3000: // Space
			ret = append(ret, 0x0020)
		case c == 0x2019: //Single quote (')
			ret = append(ret, 0x0027)
		case c == 0x201d: // Double quote (")
			ret = append(ret, 0x0022)
		default:
			ret = append(ret, c)
		}
	}

	return string(ret)
}
