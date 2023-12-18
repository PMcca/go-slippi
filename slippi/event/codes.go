package event

// Code is a type representing the Event codes in a raw array.
type Code uint8

const (
	EventPayloadsEvent   Code = 0x35
	EventGameStart       Code = 0x36
	EventPreFrame        Code = 0x37
	EventPostFrame       Code = 0x38
	EventGameEnd         Code = 0x39
	EventFrameStart      Code = 0x3A
	EventItemUpdate      Code = 0x3B
	EventFrameBookend    Code = 0x3C
	EventGeckoList       Code = 0x3D
	EventMessageSplitter Code = 0x10
)
