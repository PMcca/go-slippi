package slippi

import "fmt"

// Holds events and each respective event handler.
const (
	eventPayloadsEvent = 0x35
	eventGameStart     = 0x36
	eventPreFrame      = 0x37
	eventPostFrame     = 0x38
	eventGameEnd       = 0x39
	eventFrameStart    = 0x3A
	eventItemUpdate    = 0x3B
	eventFrameBookend  = 0x3C
	eventGeckoList     = 0x3D
)

type eventPayloads struct {
}

// parseEventPayloads parses the EventPayloads event and returns the payloads of each event.
func parseEventPayloads(d *decoder) eventPayloads {
	payloadSize := d.read()
	fmt.Println(payloadSize)
	return eventPayloads{}
}
