package slippi

const (
	ErrDecodingStartAt Error = iota
	ErrDecodingLastFrame
	ErrDecodingPlayers
	ErrEmptyFilePath
	ErrOpeningFile
	ErrParsingMeta
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
	case ErrEmptyFilePath:
		return "file path is empty"
	case ErrOpeningFile:
		return "failed to open file"
	case ErrParsingMeta:
		return "failed to parse metadata"
	default:
		return "unknown error"
	}
}
