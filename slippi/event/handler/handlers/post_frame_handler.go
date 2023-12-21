package handlers

import (
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/event"
	"github.com/PMcca/go-slippi/slippi/melee"
)

// PostFrameHandler handles the parsing of PostFrame events.
type PostFrameHandler struct{}

// Parse implements the handler.EventHandler interface. It parses a PostFrame event and puts its output into the
// given slippi.Data struct.
func (h PostFrameHandler) Parse(dec *event.Decoder, data *slippi.Data) error {
	if data.Frames == nil {
		data.Frames = map[int]slippi.Frame{}
	}

	frameNumber := dec.ReadInt32(0x1)
	playerIndex := dec.Read(0x5)
	isFollower := dec.ReadBool(0x6)
	internalCharacterID := melee.InternalCharacterID(dec.Read(0x7))
	selfInducedSpeeds := slippi.SelfInducedSpeeds{
		AirX:    dec.ReadFloat32(0x35),
		AirY:    dec.ReadFloat32(0x39),
		AttackX: dec.ReadFloat32(0x3d),
		AttackY: dec.ReadFloat32(0x41),
		GroundX: dec.ReadFloat32(0x45),
	}

	postFrame := slippi.PostFrameUpdate{
		FrameNumber:             frameNumber,
		PlayerIndex:             playerIndex,
		IsFollower:              isFollower,
		CharacterID:             internalCharacterID,
		ActionStateID:           dec.ReadUint16(0x8),
		XPos:                    dec.ReadFloat32(0xa),
		YPos:                    dec.ReadFloat32(0xe),
		FacingDirection:         dec.ReadFloat32(0x12),
		Percent:                 dec.ReadFloat32(0x16),
		ShieldSize:              dec.ReadFloat32(0x1a),
		LastHittingAttackID:     dec.Read(0x1e),
		CurrentComboCount:       dec.Read(0x1f),
		LastHitBy:               dec.Read(0x20),
		StocksRemaining:         dec.Read(0x21),
		ActionStateFrameCounter: dec.ReadFloat32(0x22),
		MiscActionState:         dec.ReadFloat32(0x2b),
		IsAirborne:              dec.ReadBool(0x2f),
		LastGroundID:            dec.ReadUint16(0x30),
		JumpsRemaining:          dec.Read(0x32),
		LCancelStatus:           dec.Read(0x33),
		HurtboxCollisionState:   slippi.HurtboxCollisionState(dec.Read(0x34)),
		SelfInducedSpeeds:       selfInducedSpeeds,
		HitlagRemaining:         dec.ReadFloat32(0x49),
		AnimationIndex:          dec.ReadUint32(0x4d),
		InstanceHitBy:           dec.ReadUint16(0x51),
		InstanceID:              dec.ReadUint16(0x53),
	}

	frame := fetchFrame(frameNumber, data)
	if isFollower {
		f := frame.Followers[playerIndex]
		f.Post = postFrame
		frame.Followers[playerIndex] = f
	} else {
		pl := frame.Players[playerIndex]
		pl.Post = postFrame
		frame.Players[playerIndex] = pl
	}

	// Pre-1.6.0, GameStart's CharacterID will be Zelda even if player started as Sheik. This check fixes this.
	if frameNumber <= slippi.FirstFrame {
		p := data.GameStart.Players[playerIndex]
		switch internalCharacterID {
		case melee.Int_Sheik:
			p.CharacterID = melee.Ext_Sheik
		case melee.Int_Zelda:
			p.CharacterID = melee.Ext_Zelda
		}
		data.GameStart.Players[playerIndex] = p
	}

	data.Frames[frameNumber] = frame
	return nil
}
