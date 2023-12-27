package handlers

import (
	"github.com/pmcca/go-slippi/slippi"
	"github.com/pmcca/go-slippi/slippi/event"
	"github.com/pmcca/go-slippi/slippi/melee"
)

// ItemUpdateHandler handles the parsing of ItemUpdates events.
type ItemUpdateHandler struct{}

// Parse implements the handler.EventHandler interface. It parses a ItemUpdates event and puts its output into the
// given slippi.Data struct.
func (h ItemUpdateHandler) Parse(dec *event.Decoder, data *slippi.Data) error {
	if data.Frames == nil {
		data.Frames = map[int]slippi.Frame{}
	}
	frameNumber := dec.ReadInt32(0x1)
	itemUpdate := slippi.ItemUpdate{
		FrameNumber:          frameNumber,
		ItemTypeID:           melee.Item(dec.ReadUint16(0x5)),
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
	frame.ItemUpdates = append(frame.ItemUpdates, itemUpdate)
	data.Frames[frameNumber] = frame

	return nil
}
