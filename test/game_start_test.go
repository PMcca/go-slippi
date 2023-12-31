package test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pmcca/go-slippi/slippi"
	"github.com/pmcca/go-slippi/slippi/melee"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGameStart(t *testing.T) {
	allItems := []melee.EnabledItem{ // All items are enabled
		melee.EnabledMetalBox,
		melee.EnabledCloakingDevice,
		melee.EnabledPokeBall,
		melee.EnabledUnknownItemBit4,
		melee.EnabledUnknownItemBit5,
		melee.EnabledUnknownItemBit6,
		melee.EnabledUnknownItemBit7,
		melee.EnabledUnknownItemBit8,
		melee.EnabledFan,
		melee.EnabledFireFlower,
		melee.EnabledSuperMushroom,
		melee.EnabledPoisonMushroom,
		melee.EnabledHammer,
		melee.EnabledWarpStar,
		melee.EnabledScrewAttack,
		melee.EnabledBunnyHood,
		melee.EnabledRayGun,
		melee.EnabledFreezie,
		melee.EnabledFood,
		melee.EnabledMotionSensorBomb,
		melee.EnabledFlipper,
		melee.EnabledSuperScope,
		melee.EnabledStarRod,
		melee.EnabledLipsStick,
		melee.EnabledHeartContainer,
		melee.EnabledMaximTomato,
		melee.EnabledStarman,
		melee.EnabledHomeRunBat,
		melee.EnabledBeamSword,
		melee.EnabledParasol,
		melee.EnabledGreenShell,
		melee.EnabledRedShell,
		melee.EnabledCapsule,
		melee.EnabledBox,
		melee.EnabledBarrel,
		melee.EnabledEgg,
		melee.EnabledPartyBall,
		melee.EnabledBarrelCannon,
		melee.EnabledBobOmb,
		melee.EnabledMrSaturn,
	}
	t.Parallel()
	t.Run("Parses0.1.0GameStartFromSlippiJS", func(t *testing.T) {
		t.Parallel()

		actual := mustParseSlippiGame(t, "replays/sheik_vs_ics_yoshis.slp")
		gameStart := actual.Data.GameStart
		require.Len(t, gameStart.Players, 4)

		assert.Equal(t, melee.StageYoshisStory, gameStart.Stage)
		assert.Equal(t, melee.Ext_Sheik, gameStart.Players[0].CharacterID)
		assert.Equal(t, melee.Ext_IceClimbers, gameStart.Players[1].CharacterID)
		assert.Equal(t, "0.1.0", gameStart.SlippiVersion)
	})
	t.Run("Parses3.16.0GameStart", func(t *testing.T) {
		t.Parallel()

		expected := slippi.GameStart{
			SlippiVersion:      "3.16.0",
			TimerType:          slippi.TimerDecreasing,
			InGameMode:         slippi.InGameModeStock,
			IsFriendlyFire:     true,
			IsTeams:            false,
			ItemSpawnBehaviour: slippi.ItemSpawnOff,
			Stage:              melee.StageYoshisStory,
			TimerStartSeconds:  480,
			EnabledItems:       allItems,
			Scene:              2,
			GameMode:           slippi.GameModeOnline,
			Language:           slippi.LanguageEnglish,
			RandomSeed:         6309,
			IsPAL:              false,
			IsFrozenPS:         false,
			MatchID:            "mode.unranked-2023-12-24T16:14:03.31-1",
			GameNumber:         2,
			TiebreakerNumber:   0,
		}
		expectedPlayers := []slippi.Player{
			{
				Index:                  0,
				Port:                   1,
				CharacterID:            melee.Ext_Fox,
				PlayerType:             slippi.PlayerTypeHuman,
				StartStocks:            4,
				CostumeIndex:           1,
				TeamShade:              slippi.TeamShadeNormal,
				Handicap:               9,
				TeamColour:             slippi.TeamColourBlue,
				IsStamina:              false,
				IsSilent:               false,
				IsLowGravity:           false,
				IsBlackStockIcon:       false,
				IsMetal:                false,
				IsStartOnAngelPlatform: false,
				CPULevel:               1,
				OffenseRatio:           1,
				DefenseRatio:           1,
				ModelScale:             1,
				ControllerFix:          "UCF",
				NameTag:                "",
				DisplayName:            "Manyula",
				ConnectCode:            "MANY#444",
				UserID:                 "gyhNynd2IzPaJrXttbdAA7LC2tW2",
			},
			{
				Index:                  1,
				Port:                   2,
				CharacterID:            melee.Ext_Falco,
				PlayerType:             slippi.PlayerTypeHuman,
				StartStocks:            4,
				CostumeIndex:           1,
				TeamShade:              slippi.TeamShadeNormal,
				Handicap:               9,
				TeamColour:             slippi.TeamColourRed,
				IsStamina:              false,
				IsSilent:               false,
				IsLowGravity:           false,
				IsBlackStockIcon:       false,
				IsMetal:                false,
				IsStartOnAngelPlatform: false,
				CPULevel:               1,
				OffenseRatio:           1,
				DefenseRatio:           1,
				ModelScale:             1,
				ControllerFix:          "UCF",
				NameTag:                "",
				DisplayName:            "sandwichyo",
				ConnectCode:            "SAND#103",
				UserID:                 "01lb4B1tQLQGxMaR3qOdj07YIUH3",
			},
		}

		actual := mustParseSlippiGame(t, "replays/3-16-0-online.slp").Data.GameStart
		diff := cmp.Diff(
			expected,
			actual,
			cmpopts.IgnoreFields(slippi.GameStart{}, "Players"))
		if diff != "" {
			t.Logf("GameStart not equal. Diff: %s", diff)
			t.Fail()
		}

		require.Len(t, actual.Players, 4)
		require.Equal(t, expectedPlayers[0], actual.Players[0])
		require.Equal(t, expectedPlayers[1], actual.Players[1])
		// YoungLink is the default character if no player is present.
		require.Equal(t, melee.Ext_YoungLink, actual.Players[2].CharacterID)
		require.Equal(t, melee.Ext_YoungLink, actual.Players[3].CharacterID)
	})

	t.Run("Parses4PlayerLocal", func(t *testing.T) {
		t.Parallel()

		expected := slippi.GameStart{
			SlippiVersion:      "3.12.0",
			TimerType:          slippi.TimerDecreasing,
			InGameMode:         slippi.InGameModeStock,
			IsFriendlyFire:     true,
			IsTeams:            false,
			ItemSpawnBehaviour: slippi.ItemSpawnOff,
			Stage:              melee.StageYoshisStory,
			TimerStartSeconds:  480,
			EnabledItems:       allItems,
			Players: []slippi.Player{
				{
					Index:                  0,
					Port:                   1,
					CharacterID:            melee.Ext_Fox,
					PlayerType:             slippi.PlayerTypeHuman,
					StartStocks:            4,
					CostumeIndex:           1,
					TeamShade:              slippi.TeamShadeNormal,
					Handicap:               9,
					TeamColour:             slippi.TeamColourRed,
					IsStamina:              false,
					IsSilent:               false,
					IsLowGravity:           false,
					IsBlackStockIcon:       false,
					IsMetal:                false,
					IsStartOnAngelPlatform: false,
					CPULevel:               1,
					OffenseRatio:           1,
					DefenseRatio:           1,
					ModelScale:             1,
					ControllerFix:          "UCF",
					NameTag:                "ABC4",
					DisplayName:            "",
					ConnectCode:            "",
					UserID:                 "",
				},
				{
					Index:                  1,
					Port:                   2,
					CharacterID:            melee.Ext_Marth,
					PlayerType:             slippi.PlayerTypeCPU,
					StartStocks:            4,
					CostumeIndex:           0,
					TeamShade:              slippi.TeamShadeNormal,
					Handicap:               9,
					TeamColour:             slippi.TeamColourRed,
					IsStamina:              false,
					IsSilent:               false,
					IsLowGravity:           false,
					IsBlackStockIcon:       false,
					IsMetal:                false,
					IsStartOnAngelPlatform: false,
					CPULevel:               1,
					OffenseRatio:           1,
					DefenseRatio:           1,
					ModelScale:             1,
					ControllerFix:          "UCF",
					NameTag:                "",
					DisplayName:            "",
					ConnectCode:            "",
					UserID:                 "",
				},
				{
					Index:                  2,
					Port:                   3,
					CharacterID:            melee.Ext_Samus,
					PlayerType:             slippi.PlayerTypeCPU,
					StartStocks:            4,
					CostumeIndex:           0,
					TeamShade:              slippi.TeamShadeNormal,
					Handicap:               9,
					TeamColour:             slippi.TeamColourRed,
					IsStamina:              false,
					IsSilent:               false,
					IsLowGravity:           false,
					IsBlackStockIcon:       false,
					IsMetal:                false,
					IsStartOnAngelPlatform: false,
					CPULevel:               1,
					OffenseRatio:           1,
					DefenseRatio:           1,
					ModelScale:             1,
					ControllerFix:          "UCF",
					NameTag:                "",
					DisplayName:            "",
					ConnectCode:            "",
					UserID:                 "",
				},
				{
					Index:                  3,
					Port:                   4,
					CharacterID:            melee.Ext_Roy,
					PlayerType:             slippi.PlayerTypeCPU,
					StartStocks:            4,
					CostumeIndex:           0,
					TeamShade:              slippi.TeamShadeNormal,
					Handicap:               9,
					TeamColour:             slippi.TeamColourRed,
					IsStamina:              false,
					IsSilent:               false,
					IsLowGravity:           false,
					IsBlackStockIcon:       false,
					IsMetal:                false,
					IsStartOnAngelPlatform: false,
					CPULevel:               1,
					OffenseRatio:           1,
					DefenseRatio:           1,
					ModelScale:             1,
					ControllerFix:          "UCF",
					NameTag:                "",
					DisplayName:            "",
					ConnectCode:            "",
					UserID:                 "",
				},
			},
			Scene:            2,
			GameMode:         slippi.GameModeVS,
			Language:         slippi.LanguageEnglish,
			RandomSeed:       793180696,
			IsPAL:            false,
			IsFrozenPS:       false,
			MatchID:          "",
			GameNumber:       0,
			TiebreakerNumber: 0,
		}

		actual := mustParseSlippiGame(t, "replays/4-player-offline.slp").Data.GameStart
		require.Equal(t, expected, actual)
	})
}

func TestNameTags(t *testing.T) {
	t.Parallel()
	t.Run("ParsesNameTag1", func(t *testing.T) {
		t.Parallel()

		actual := mustParseSlippiGame(t, "replays/nametags.slp")
		require.Len(t, actual.Data.GameStart.Players, 4)
		assert.Equal(t, "AMNイ", actual.Data.GameStart.Players[0].NameTag)
	})
	t.Run("ParsesNameTag2", func(t *testing.T) {
		t.Parallel()

		actual := mustParseSlippiGame(t, "replays/nametags2.slp")
		require.Len(t, actual.Data.GameStart.Players, 4)
		assert.Equal(t, "A1=$", actual.Data.GameStart.Players[0].NameTag)
		assert.Equal(t, "か、9@", actual.Data.GameStart.Players[1].NameTag)
	})
	t.Run("ParsesNameTag3", func(t *testing.T) {
		t.Parallel()

		actual := mustParseSlippiGame(t, "replays/nametags3.slp")
		require.Len(t, actual.Data.GameStart.Players, 4)
		assert.Equal(t, "B  R", actual.Data.GameStart.Players[0].NameTag)
		assert.Equal(t, ".  。", actual.Data.GameStart.Players[1].NameTag)
	})
}

func TestPAL(t *testing.T) {
	t.Parallel()
	t.Run("ReadsPALGame", func(t *testing.T) {
		t.Parallel()

		actual := mustParseSlippiGame(t, "replays/pal.slp")
		assert.Equal(t, true, actual.Data.GameStart.IsPAL)
	})
	t.Run("PALIsFalseWithNTSCGame", func(t *testing.T) {
		t.Parallel()

		actual := mustParseSlippiGame(t, "replays/ntsc.slp")
		assert.Equal(t, false, actual.Data.GameStart.IsPAL)
	})
}

func TestControllerFix(t *testing.T) {
	t.Parallel()
	t.Run("ReadsDifferentControllerFixes", func(t *testing.T) {
		t.Parallel()

		actual := mustParseSlippiGame(t, "replays/controller-fixes.slp")
		assert.Equal(t, "Dween", actual.Data.GameStart.Players[0].ControllerFix)
		assert.Equal(t, "UCF", actual.Data.GameStart.Players[1].ControllerFix)
		assert.Equal(t, "None", actual.Data.GameStart.Players[2].ControllerFix)
	})
}

func TestMatchInfo(t *testing.T) {
	t.Parallel()
	t.Run("ReadsRankedMatchInfo", func(t *testing.T) {
		t.Parallel()

		actual := mustParseSlippiGame(t, "replays/ranked_game1_tiebreak.slp")
		gameStart := actual.Data.GameStart
		assert.Equal(t, 1, gameStart.GameNumber, "GameNumber not equal")
		assert.Equal(t, 1, gameStart.TiebreakerNumber, "TieBreakerNumber not equal")
		assert.Equal(t, "mode.ranked-2022-12-20T05:36:47.50-0", gameStart.MatchID)
	})

	t.Run("ReadsUnrankedGame1", func(t *testing.T) {
		t.Parallel()

		actual := mustParseSlippiGame(t, "replays/unranked_game1.slp")
		gameStart := actual.Data.GameStart
		assert.Equal(t, 1, gameStart.GameNumber, "GameNumber not equal")
		assert.Equal(t, 0, gameStart.TiebreakerNumber, "TieBreakerNumber not equal")
		assert.Equal(t, "mode.unranked-2022-12-21T02:26:27.50-0", gameStart.MatchID)
	})

	t.Run("ReadsUnrankedGame2", func(t *testing.T) {
		t.Parallel()

		actual := mustParseSlippiGame(t, "replays/unranked_game2.slp")
		gameStart := actual.Data.GameStart
		assert.Equal(t, 2, gameStart.GameNumber, "GameNumber not equal")
		assert.Equal(t, 0, gameStart.TiebreakerNumber, "TieBreakerNumber not equal")
		assert.Equal(t, "mode.unranked-2022-12-21T02:26:27.50-0", gameStart.MatchID)
	})

	t.Run("EmptyValuesForOldReplay", func(t *testing.T) {
		t.Parallel()

		actual := mustParseSlippiGame(t, "replays/old-test.slp")
		gameStart := actual.Data.GameStart
		assert.Equal(t, 0, gameStart.GameNumber, "GameNumber not equal")
		assert.Equal(t, 0, gameStart.TiebreakerNumber, "TieBreakerNumber not equal")
		assert.Equal(t, "", gameStart.MatchID)
	})
}
