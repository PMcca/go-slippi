package slippi

import "github.com/pmcca/go-slippi/slippi/melee"

const (
	FirstFrame         = -123
	FirstPlayableFrame = -39
)

// HurtboxCollisionState represents the possible HurtboxCollision states a character can be in.
type HurtboxCollisionState uint8

const (
	HurtboxStateVulnerable   HurtboxCollisionState = iota // 0
	HurtboxStateInvulnerable                              // 1
	HurtboxStateIntangible                                // 2

)

// MissileType represents the two kinds of missiles (Side-B) Samus can fire.
type MissileType uint8

const (
	MissileHoming MissileType = iota // 0
	MissileSuper                     // 1
)

// TurnipFace represents the different turnips (Down-B) Peach can pull
type TurnipFace uint8

const (
	TurnipSmile       TurnipFace = iota // 0
	TurnipTEyes                         // 1
	TurnipLineEyes                      // 2
	TurnipCircleEyes                    // 3
	TurnipUpwardCurve                   // 4
	TurnipWink                          // 5
	TurnipDotEyes                       // 6
	TurnipStitchFace                    // 7
)

// SelfInducedSpeeds holds the speeds that are induced by the player.
type SelfInducedSpeeds struct {
	AirX    float32
	AirY    float32
	AttackX float32
	AttackY float32
	GroundX float32
}

// PreFrameUpdate represents a parsed event.EventPreFrame event.
type PreFrameUpdate struct {
	FrameNumber      int
	PlayerIndex      uint8
	IsFollower       bool
	RandomSeed       uint32
	ActionStateID    uint16
	XPos             float32
	YPos             float32
	FacingDirection  float32
	JoyStickX        float32
	JoyStickY        float32
	CStickX          float32
	CStickY          float32
	Trigger          float32
	ProcessedButtons uint32
	PhysicalButtons  uint16
	PhysicalTriggerL float32
	PhysicalTriggerR float32
	XAnalogUCF       int8
	Percent          float32
	YAnalogUCF       int8
}

// PostFrameUpdate represents a parsed event.EventPostFrame event.
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
	PlayerIndex uint8
	Pre         PreFrameUpdate
	Post        PostFrameUpdate
}

// FrameStart represents a parsed event.EventFrameStart event.
type FrameStart struct {
	FrameNumber       int
	Seed              uint32
	SceneFrameCounter uint32
}

// ItemUpdate represents a parsed event.EventItemUpdate event.
type ItemUpdate struct {
	FrameNumber          int
	ItemTypeID           melee.Item
	State                uint8
	FacingDirection      float32
	XVelocity            float32
	YVelocity            float32
	XPos                 float32
	YPos                 float32
	DamageTaken          uint16
	ExpirationTimer      float32
	SpawnID              uint32
	MissileType          MissileType
	TurnipFace           TurnipFace
	ChargeShotIsLaunched bool
	ChargeShotPower      uint8
	Owner                int8
	InstanceID           uint16
}

// FrameBookend represents a parsed event.EventFrameBookend event.
type FrameBookend struct {
	FrameNumber          int
	LatestFinalisedFrame int
}

// Frame represents a single, complete frame in-game, including the updates for all the characters for said frame.
type Frame struct {
	FrameNumber  int
	FrameStart   FrameStart
	Players      map[uint8]PlayerFrameUpdate // Map of PlayerIndex -> PlayerFrameUpdate
	Followers    map[uint8]PlayerFrameUpdate
	ItemUpdates  []ItemUpdate
	FrameBookend FrameBookend
}
