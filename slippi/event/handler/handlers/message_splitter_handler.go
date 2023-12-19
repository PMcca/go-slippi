package handlers

import (
	"fmt"
	"github.com/PMcca/go-slippi/internal/logging"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/event"
)

var log = logging.NewLogger()

const (
	// Size of MessageSplitter events is always 516 bytes (as of 3.3.0)
	messageSplitterSize = 516
)

type MessageSplitterHandler struct{}

// Parse implements the handler.EventHandler interface. It reads all the payloads of each MessageSplitter event
// encountered, then passes these bytes in to the relevant handler.
func (m MessageSplitterHandler) Parse(dec *event.Decoder, data *slippi.Data) error {
	// Read all contiguous MessageSplitter payloads into a buffer
	eventCode := event.Code(dec.Read(0x203))
	var buffer []byte
	for {
		// Assume all MessageSplitter events are contiguous. If we encounter a non-message splitter event, error.
		if event.Code(dec.Read(0x0)) != event.EventMessageSplitter {
			return fmt.Errorf("%w:eventCode %0x", ErrNoMessageSplitterCode, dec.Read(0x0))
		}
		eventSize := dec.ReadInt16(0x201)
		buffer = append(buffer, dec.ReadN(0x1, eventSize)...)

		isLastMessage := dec.ReadBool(0x204)
		if isLastMessage {
			break
		}

		dec.Data = dec.Data[messageSplitterSize+1:]
	}

	// Parse the corresponding event with the accumulated bytes. As of 3.3.0, MessageSplitter events are only used for
	// GeckoCode events.
	switch eventCode {
	case event.EventGeckoList:
		return GeckoCodeHandler{}.Parse(
			&event.Decoder{
				Data: buffer,
				Size: len(buffer),
			}, data)
	default:
		return ErrUnknownMessageSplitEvent
	}
}
