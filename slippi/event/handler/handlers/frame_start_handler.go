package handlers

import (
	"github.com/pmcca/go-slippi/slippi"
	"github.com/pmcca/go-slippi/slippi/event"
)

// FrameStartHandler handles the parsing FrameStart of events.
type FrameStartHandler struct{}

// Parse implements the handler.EventHandler interface. It parses a FrameStart event and puts its output into the
// given slippi.Data struct.
func (h FrameStartHandler) Parse(dec *event.Decoder, data *slippi.Data) error {
	if data.Frames == nil {
		data.Frames = map[int]slippi.Frame{}
	}

	frameNumber := dec.ReadInt32(0x1)
	frameStart := slippi.FrameStart{
		FrameNumber:       frameNumber,
		Seed:              dec.ReadUint32(0x5),
		SceneFrameCounter: dec.ReadUint32(0x9),
	}

	frame := data.Frames[frameNumber]
	frame.FrameStart = frameStart
	frame.FrameNumber = frameNumber
	data.Frames[frameNumber] = frame

	return nil
}
