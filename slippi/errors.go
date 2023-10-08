package slippi

const (
	ErrEmptyFilePath Error = iota
	ErrReadingFile
	ErrParsingMeta
)

type Error uint

// Error returns the string message for the given error.
func (e Error) Error() string {
	switch e {
	case ErrEmptyFilePath:
		return "file path is empty"
	case ErrReadingFile:
		return "failed to read file"
	case ErrParsingMeta:
		return "failed to parse metadata"
	default:
		return "unknown error"
	}
}
