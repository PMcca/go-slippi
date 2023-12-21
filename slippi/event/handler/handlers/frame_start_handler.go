package handlers

import (
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/event"
)

type FrameStartHandler struct{}

func (h FrameStartHandler) Parse(dec *event.Decoder, data *slippi.Data) error {
	frameNumber := dec.ReadInt32(0x1)
	frameStart := slippi.FrameStart{
		FrameNumber:       frameNumber,
		Seed:              dec.ReadUint32(0x5),
		SceneFrameCounter: dec.ReadUint32(0x9),
	}

	if data.Frames == nil {
		data.Frames = map[int]slippi.Frame{}
	}
	frame := data.Frames[frameNumber]
	frame.FrameStart = frameStart
	frame.FrameNumber = frameNumber
	data.Frames[frameNumber] = frame

	return nil
}
