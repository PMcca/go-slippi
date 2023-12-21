package handlers

import (
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/event"
)

// FrameBookendHandler handles the parsing FrameBookend of events.
type FrameBookendHandler struct{}

// Parse implements the handler.EventHandler interface. It parses a FrameBookend event and puts its output into the
// given slippi.Data struct.
func (h FrameBookendHandler) Parse(dec *event.Decoder, data *slippi.Data) error {
	if data.Frames == nil {
		data.Frames = map[int]slippi.Frame{}
	}

	frameNumber := dec.ReadInt32(0x1)
	frameBookend := slippi.FrameBookend{
		FrameNumber:          frameNumber,
		LatestFinalisedFrame: dec.ReadInt32(0x5),
	}

	frame := data.Frames[frameNumber]
	frame.FrameBookend = frameBookend
	data.Frames[frameNumber] = frame
	return nil
}
