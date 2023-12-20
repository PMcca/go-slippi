package handlers

import (
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/event"
)

type ItemUpdateHandler struct{}

func (h ItemUpdateHandler) Parse(dec *event.Decoder, data *slippi.Data) error {
	if data.Frames == nil {
		data.Frames = map[int]slippi.Frame{}
	}
	frameNumber := dec.ReadInt32(0x1)
	itemUpdate := slippi.ItemUpdate{
		FrameNumber:          frameNumber,
		ItemTypeID:           dec.ReadUint16(0x5),
		State:                dec.Read(0x7),
		FacingDirection:      dec.ReadFloat32(0x8),
		XVelocity:            dec.ReadFloat32(0xc),
		YVelocity:            dec.ReadFloat32(0x10),
		XPos:                 dec.ReadFloat32(0x14),
		YPos:                 dec.ReadFloat32(0x18),
		DamageTaken:          dec.ReadUint16(0x1c),
		ExpirationTimer:      dec.ReadFloat32(0x1e),
		SpawnID:              dec.ReadUint32(0x22),
		MissileType:          slippi.MissileType(dec.Read(0x26)),
		TurnipFace:           slippi.TurnipFace(dec.Read(0x27)),
		ChargeShotIsLaunched: dec.ReadBool(0x28),
		ChargeShotPower:      dec.Read(0x29),
		Owner:                dec.ReadInt8(0x2a),
		InstanceID:           dec.ReadUint16(0x2b),
	}

	frame := fetchFrame(frameNumber, data)
	frame.ItemUpdate = append(frame.ItemUpdate, itemUpdate)
	data.Frames[frameNumber] = frame

	return nil
}
