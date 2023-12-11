package slippi

const (
	ErrEmptyFilePath Error = iota
	ErrReadingFile
	ErrParsingGame
	ErrParsingMeta
	ErrInvalidRawStart
	ErrEventPayloadsNotFound
	ErrInvalidNumberOfCommands
	ErrUnknownEventInEventSizes
)

type Error uint

// Error returns the string message for the given error.
func (e Error) Error() string {
	switch e {
	case ErrEmptyFilePath:
		return "file path is empty"
	case ErrReadingFile:
		return "failed to read file"
	case ErrParsingGame:
		return "failed to parse game"
	case ErrParsingMeta:
		return "failed to parse metadata"
	case ErrInvalidRawStart:
		return "unexpected beginning of raw array"
	case ErrEventPayloadsNotFound:
		return "event payloads not found in raw"
	case ErrInvalidNumberOfCommands:
		return "invalid number of commands in event payloads, must be divisible by 3"
	case ErrUnknownEventInEventSizes:
		return "unknown event in event payload sizes"
	default:
		return "unknown error"
	}
}
