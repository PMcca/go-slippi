package slippi

// GameEndMethod is the way in which the game ended, for example if time ran out, or there was no contest.
type GameEndMethod uint8

const (
	GameEndUnresolved GameEndMethod = iota // 0
	GameEndTime                            // 1
	GameEndGameSet                         // 2
	GameEndResolved                        // 3
	GameEndNoContest  GameEndMethod = 7
)

// PlayerPlacement is the final placement (1st, 2nd, etc.) of a single player.
type PlayerPlacement struct {
	PlayerIndex int8
	Placement   int8
}

// GameEnd represents a parsed event.EventGameEnd event.
type GameEnd struct {
	GameEndMethod    GameEndMethod
	LRASInitiatior   int8
	PlayerPlacements []PlayerPlacement
}
