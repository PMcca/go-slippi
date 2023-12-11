package slippi

import (
	"encoding/binary"
	"fmt"
)

// eventType is a type representing the event types in a raw array.
type eventType uint8

// Holds events and each respective event handler.
const (
	eventPayloadsEvent eventType = 0x35
	eventGameStart     eventType = 0x36
	eventPreFrame      eventType = 0x37
	eventPostFrame     eventType = 0x38
	eventGameEnd       eventType = 0x39
	eventFrameStart    eventType = 0x3A
	eventItemUpdate    eventType = 0x3B
	eventFrameBookend  eventType = 0x3C
	eventGeckoList     eventType = 0x3D
)

// parseEventPayloads parses the EventPayloads event and returns a map of event type -> payload size
func parseEventPayloads(d []byte) (map[eventType]int, error) {
	if eventType(d[0]) != eventPayloadsEvent {
		return nil, ErrEventPayloadsNotFound
	}
	payloadSize := d[1]

	// Command byte + payload size = 3 bytes, so divide by 3 to get # of commands. payloadSize includes itself, so -1.
	if (payloadSize-1)%3 != 0 {
		return nil, ErrInvalidNumberOfCommands
	}
	commands := int(payloadSize / 3)

	offset := 2
	eventSizes := map[eventType]int{eventPayloadsEvent: int(payloadSize)}
	for i := 0; i < commands; i++ {
		command := eventType(d[offset])                              // Read command byte.
		size := int(binary.BigEndian.Uint16(d[offset+1 : offset+3])) // Read command size, uint16 = 2 bytes.
		eventSizes[command] = size

		offset += 3
	}

	return eventSizes, nil
}

// parseGameStart parses a GameStart event and populates the given Data struct with its contents.
func parseGameStart(eventSize int, dec *decoder, data *Data) error {
	slippiVersion := fmt.Sprintf("%d.%d.%d",
		dec.read(0x1),
		dec.read(0x2),
		dec.read(0x3),
	)

	//timerType := dec.readWithBitmask(0x5, 0x03)
	//inGameMode := dec.readWithBitmask(0x5, 0xe0)
	//
	//var isFriendlyFire bool
	//if dec.readWithBitmask(0x6, 0x01) == 1 {
	//	isFriendlyFire = true
	//}

	fmt.Println(slippiVersion)
	//timerType := dec.read() & 0x03
	return nil
}
