package handlers

import (
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/event"
)

type PreFrameHandler struct{}

func (p PreFrameHandler) Parse(dec *event.Decoder, data *slippi.Data) error {
	if data.Frames == nil {
		data.Frames = map[int]slippi.Frame{}
	}

	frameNumber := dec.ReadInt32(0x1)
	playerIndex := dec.Read(0x5)
	isFollower := dec.ReadBool(0x6)

	preFrame := slippi.PreFrameUpdate{
		FrameNumber:      frameNumber,
		PlayerIndex:      playerIndex,
		IsFollower:       isFollower,
		RandomSeed:       dec.ReadUint32(0x7),
		ActionStateID:    dec.ReadUint16(0xb),
		XPos:             dec.ReadFloat32(0xd),
		YPos:             dec.ReadFloat32(0x11),
		FacingDirection:  dec.ReadFloat32(0x15),
		JoyStickX:        dec.ReadFloat32(0x19),
		JoyStickY:        dec.ReadFloat32(0x1d),
		CStickX:          dec.ReadFloat32(0x21),
		CStickY:          dec.ReadFloat32(0x25),
		Trigger:          dec.ReadFloat32(0x29),
		ProcessedButtons: dec.ReadUint32(0x2d),
		PhysicalButtons:  dec.ReadUint16(0x31),
		PhysicalTriggerL: dec.ReadFloat32(0x33),
		PhysicalTriggerR: dec.ReadFloat32(0x37),
		XAnalogUCF:       dec.ReadInt8(0x3b),
		Percent:          dec.ReadFloat32(0x3c),
		YAnalogUCF:       dec.ReadInt8(0x40),
	}

	frame := fetchFrame(frameNumber, data)
	if isFollower {
		f := frame.Followers[playerIndex]
		f.Pre = preFrame
		frame.Followers[playerIndex] = f
	} else {
		pl := frame.Players[playerIndex]
		pl.Pre = preFrame
		frame.Players[playerIndex] = pl
	}

	data.Frames[frameNumber] = frame
	return nil
}

// fetchFrame returns the frame for the given frameNumber. It also ensures the Players and Followers maps are not nil.
func fetchFrame(frameNumber int, data *slippi.Data) slippi.Frame {
	frame := data.Frames[frameNumber]
	if frame.Players == nil {
		frame.Players = map[uint8]slippi.PlayerFrameUpdate{}
	}
	if frame.Followers == nil {
		frame.Followers = map[uint8]slippi.PlayerFrameUpdate{}
	}
	return frame
}
