package goslippi

// Error is an error for this package.
type Error string

const (
	ErrEmptyFilePath            Error = "file path is empty"
	ErrReadingFile              Error = "failed to read file"
	ErrParsingGame              Error = "failed to parse .slp file"
	ErrParsingMeta              Error = "failed to parse metadata"
	ErrInvalidRawStart          Error = "unexpected beginning of raw array"
	ErrUnknownEventInEventSizes Error = "unknown event in event payload sizes"
	ErrNoHandlerForEvent        Error = "unable to handle unknown event"
	ErrFailedEventParsing       Error = "failed to parse event"
)

// Error implements the error interface.
func (e Error) Error() string {
	return string(e)
}
