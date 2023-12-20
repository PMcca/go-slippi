package slippi

import "github.com/PMcca/go-slippi/slippi/melee"

type HurtboxCollisionState uint8

const (
	HurtboxStateVulnerable   HurtboxCollisionState = iota // 0
	HurtboxStateInvulnerable                              // 1
	HurtboxStateIntangible                                // 2

)

type SelfInducedSpeeds struct {
	AirX    float32
	AirY    float32
	AttackX float32
	AttackY float32
	GroundX float32
}

type PreFrameUpdate struct {
	FrameNumber      int
	PlayerIndex      uint8
	IsFollower       bool
	RandomSeed       uint32
	ActionStateID    uint16 //TODO enum?
	XPos             float32
	YPos             float32
	FacingDirection  float32
	JoyStickX        float32
	JoyStickY        float32
	CStickX          float32
	CStickY          float32
	Trigger          float32
	ProcessedButtons uint32 //TODO figure out
	PhysicalButtons  uint16 // TODO enum?
	PhysicalTriggerL float32
	PhysicalTriggerR float32
	XAnalogUCF       int8
	Percent          float32
	YAnalogUCF       int8
}

type PostFrameUpdate struct {
	FrameNumber             int
	PlayerIndex             uint8
	IsFollower              bool
	CharacterID             melee.InternalCharacterID
	ActionStateID           uint16
	XPos                    float32
	YPos                    float32
	FacingDirection         float32
	Percent                 float32
	ShieldSize              float32
	LastHittingAttackID     uint8
	CurrentComboCount       uint8
	LastHitBy               uint8
	StocksRemaining         uint8
	ActionStateFrameCounter float32
	MiscActionState         float32
	IsAirborne              bool
	LastGroundID            uint16
	JumpsRemaining          uint8
	LCancelStatus           uint8
	HurtboxCollisionState   HurtboxCollisionState
	SelfInducedSpeeds       SelfInducedSpeeds
	HitlagRemaining         float32
	AnimationIndex          uint32
	InstanceHitBy           uint16
	InstanceID              uint16
}

// PlayerFrameUpdate holds the Pre/Post-frame updates for a given player/follower.
type PlayerFrameUpdate struct {
	Pre  PreFrameUpdate
	Post PostFrameUpdate
}

type FrameStart struct {
	FrameNumber       int
	Seed              uint32
	SceneFrameCounter uint32
}

type ItemUpdate struct {
}

// Frame represents a single, complete frame in-game, including the updates for all the characters for said frame.
type Frame struct {
	FrameNumber int
	FrameStart  FrameStart
	Players     map[uint8]PlayerFrameUpdate // Map of PlayerIndex -> PlayerFrameUpdate
	Followers   map[uint8]PlayerFrameUpdate
	ItemUpdate  []ItemUpdate
}
