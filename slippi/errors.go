package slippi

const (
	ErrDecodingField Error = iota
)

type Error uint

// Error returns the string message for the given error.
func (e Error) Error() string {
	switch e {
	case ErrDecodingField:
		return "failed to decode field %s"
	default:
		return "unknown error"
	}
}
