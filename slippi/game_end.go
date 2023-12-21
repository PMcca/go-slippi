package slippi

type GameEndMethod uint8

const (
	GameEndUnresolved GameEndMethod = iota // 0
	GameEndTime                            // 1
	GameEndGameSet                         // 2
	GameEndResolved                        // 3
	GameEndNoContest  GameEndMethod = 7
)

type PlayerPlacement struct {
	PlayerIndex int8
	Placement   int8
}

type GameEnd struct {
	GameEndMethod    GameEndMethod
	LRASInitiatior   int8
	PlayerPlacements []PlayerPlacement
}
