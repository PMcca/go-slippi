package slippi

import "github.com/PMcca/go-slippi/slippi/melee"

// TODO add comments
type TimerType uint8

const (
	NoTimer         TimerType = 0
	TimerDecreasing TimerType = 2
	TimerIncreasing TimerType = 3
)

type InGameMode uint8

const (
	InGameModeTime  InGameMode = iota // 0
	InGameModeStock                   // 1
	InGameModeCoin                    // 2
	InGameModeBonus                   // 3
)

type ItemSpawnBehaviour int8

const (
	ItemSpawnOff      ItemSpawnBehaviour = iota - 1 // -1
	ItemSpawnVeryLow                                // 0
	ItemSpawnLow                                    // 1
	ItemSpawnMed                                    // 2
	ItemSpawnHigh                                   // 3
	ItemSpawnVeryHigh                               // 4
)

type PlayerType uint8

const (
	PlayerTypeHuman PlayerType = iota // 0
	PlayerTypeCPU                     // 1
	PlayerTypeDemo                    // 2
	PlayerTypeEmpty                   // 3
)

type TeamShade uint8

const (
	TeamShadeNormal TeamShade = iota // 0
	TeamShadeLight                   // 1
	TeamShadeDark                    // 2
)

// TeamColour is the colour of the team the player is in. Value is TeamID.
type TeamColour uint8

const (
	TeamColourRed   TeamColour = iota // 0
	TeamColourBlue                    // 1
	TeamColourGreen                   // 2
)

type GameMode uint8

const (
	GameModeVS     GameMode = 0x2
	GameModeOnline GameMode = 0x8
)

type Language uint8

const (
	LanguageJapanese Language = iota // 0
	LanguageEnglish                  // 1
)

type Player struct {
	Index                  int // Port = Index + 1
	Port                   int
	CharacterID            melee.ExternalCharacterID
	PlayerType             PlayerType
	StartStocks            uint8
	CostumeIndex           uint8
	TeamShade              TeamShade
	Handicap               uint8
	TeamColour             TeamColour
	IsStamina              bool
	IsSilent               bool
	IsLowGravity           bool
	IsInvisible            bool
	IsBlackStockIcon       bool
	IsMetal                bool
	IsStartOnAngelPlatform bool
	IsRumbleEnabled        bool
	CPULevel               uint8
	OffenseRatio           float32
	DefenseRation          float32
	ModelScale             float32 // TODO or float?
	ControllerFix          string  // TODO What is this?
	Nametag                string
	DisplayName            string
	ConnectCode            string
	UserID                 string
}

type GameStart struct {
	SlippiVersion      string
	TimerType          TimerType
	InGameMode         InGameMode
	IsFriendlyFire     bool
	IsTeams            bool
	ItemSpawnBehaviour ItemSpawnBehaviour
	Stage              melee.Stage
	TimerStartSeconds  int
	EnabledItems       []melee.Item
	Players            []Player
	Scene              uint8 // Minor scene, should always be 0x2
	GameMode           GameMode
	Language           Language
	RandomSeed         uint32
	IsPAL              bool
	IsFrozenPS         bool
	MatchID            string
	GameNumber         int
	TiebreakerNumber   int
}
