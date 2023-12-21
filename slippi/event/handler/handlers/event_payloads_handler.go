package handlers

import (
	"github.com/PMcca/go-slippi/slippi/event"
)

// ParseEventPayloads parses the EventPayloads command from the byte array, by returning a map of Code -> Event Size.
// It does not implement the handler.EventHandler interface, unlike the other handlers.
func ParseEventPayloads(d *event.Decoder) (map[event.Code]int, error) {
	if event.Code(d.Read(0)) != event.EventPayloadsEvent {
		return nil, ErrEventPayloadsNotFound
	}
	payloadSize := d.Read(1)

	// Command byte + payload size = 3 bytes, so divide by 3 to get # of commands. payloadSize includes itself, so -1.
	if (payloadSize-1)%3 != 0 {
		return nil, ErrInvalidNumberOfCommands
	}
	commands := int((payloadSize - 1) / 3)

	offset := 2
	eventSizes := map[event.Code]int{event.EventPayloadsEvent: int(payloadSize)}
	for i := 0; i < commands; i++ {
		command := event.Code(d.Read(offset)) // Read command byte.
		size := d.ReadInt16(offset + 1)
		eventSizes[command] = size

		offset += 3
	}

	return eventSizes, nil
}
