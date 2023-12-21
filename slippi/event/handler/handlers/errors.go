package handlers

// Error is an error that can occur in this package.
type Error string

const (
	ErrInvalidNumberOfCommands  Error = "invalid number of commands in event payloads, must be divisible by 3"
	ErrEventPayloadsNotFound    Error = "event payloads not found in raw"
	ErrInvalidRawStart          Error = "unexpected beginning of raw array"
	ErrUnknownEventInEventSizes Error = "unknown event in event payload sizes"
	ErrNoHandlerForEvent        Error = "unable to handle unknown event"
	ErrFailedEventParsing       Error = "failed to parse event"
	ErrNoMessageSplitterCode    Error = "event code for message splitter incorrect, but did not encounter 'lastMessage'"
	ErrUnknownMessageSplitEvent Error = "unknown event in message splitters"
)

// Error implements the error interface.
func (e Error) Error() string {
	return string(e)
}
