package slippi

type PreFrameUpdate struct {
}

type PostFrameUpdate struct {
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
	Players     map[int]PlayerFrameUpdate // Map of PlayerIndex -> PlayerFrameUpdate
	Followers   map[int]PlayerFrameUpdate
	ItemUpdate  []ItemUpdate
}
