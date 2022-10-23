package slippi

const (
	ErrDecodingStartAt Error = iota
	ErrDecodingLastFrame
	ErrDecodingPlayers
)

type Error uint

// Error returns the string message for the given error.
func (e Error) Error() string {
	switch e {
	case ErrDecodingStartAt:
		return "failed to decode startAt"
	case ErrDecodingLastFrame:
		return "failed to decode lastFrame"
	case ErrDecodingPlayers:
		return "failed to decode players"
	default:
		return "unknown error"
	}
}
