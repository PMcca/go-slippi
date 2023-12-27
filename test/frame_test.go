package test

import (
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/melee"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestOffline4PlayerFrames(t *testing.T) {
	t.Parallel()
	parsedGame := mustParseSlippiGame(t, "replays/4-player-offline.slp")
	require.Contains(t, parsedGame.Data.Frames, 715)
	actual := parsedGame.Data.Frames[715]

	t.Run("ParsesFrame", func(t *testing.T) {
		t.Parallel()

		expected := slippi.Frame{
			FrameNumber: 715,
			FrameStart: slippi.FrameStart{
				FrameNumber:       715,
				Seed:              773713751,
				SceneFrameCounter: 838,
			},
			FrameBookend: slippi.FrameBookend{
				FrameNumber:          715,
				LatestFinalisedFrame: 715,
			},
		}
		diff := cmp.Diff(
			expected,
			actual,
			cmpopts.IgnoreFields(slippi.Frame{}, "Players", "Followers", "ItemUpdates"))
		if diff != "" {
			t.Logf("Frames not equal. Diff: %s", diff)
			t.Fail()
		}
	})

	t.Run("ParsesFramePlayers", func(t *testing.T) {
		t.Parallel()

		// MistActionState for player index 1 is a very small number, so create what we expect it to be here.
		m, err := strconv.ParseFloat("1.401298e-45", 32)
		require.NoError(t, err)
		expectedPlayer1MiscState := float32(m)

		expected := map[uint8]slippi.PlayerFrameUpdate{
			0: {
				PlayerIndex: 0,
				Pre: slippi.PreFrameUpdate{
					FrameNumber:      715,
					PlayerIndex:      0,
					IsFollower:       false,
					RandomSeed:       4011085961,
					ActionStateID:    345,
					XPos:             -25.4558,
					YPos:             31.280098,
					FacingDirection:  1,
					JoyStickX:        -0.325,
					JoyStickY:        -0.9125,
					CStickX:          0,
					CStickY:          0,
					Trigger:          0,
					ProcessedButtons: 131072,
					PhysicalButtons:  0,
					PhysicalTriggerL: 0,
					PhysicalTriggerR: 0,
					XAnalogUCF:       -26,
					Percent:          15,
					YAnalogUCF:       0,
				},
				Post: slippi.PostFrameUpdate{
					FrameNumber:             715,
					PlayerIndex:             0,
					IsFollower:              false,
					CharacterID:             melee.Int_Fox,
					ActionStateID:           345,
					XPos:                    -25.131298,
					YPos:                    31.050098,
					FacingDirection:         1,
					Percent:                 15,
					ShieldSize:              60,
					LastHittingAttackID:     21,
					CurrentComboCount:       1,
					LastHitBy:               1,
					StocksRemaining:         4,
					ActionStateFrameCounter: 7,
					MiscActionState:         0,
					IsAirborne:              true,
					LastGroundID:            3,
					JumpsRemaining:          1,
					LCancelStatus:           0,
					HurtboxCollisionState:   slippi.HurtboxStateVulnerable,
					SelfInducedSpeeds: slippi.SelfInducedSpeeds{
						AirX:    0.3245001,
						AirY:    -0.23000021,
						AttackX: 0,
						AttackY: 0,
						GroundX: 0,
					},
					HitlagRemaining: 0,
					AnimationIndex:  299,
					InstanceHitBy:   0,
					InstanceID:      0,
				},
			},
			1: {
				PlayerIndex: 1,
				Pre: slippi.PreFrameUpdate{
					FrameNumber:     715,
					PlayerIndex:     1,
					IsFollower:      false,
					RandomSeed:      4011085961,
					ActionStateID:   14,
					XPos:            4.782982,
					YPos:            0.0001,
					FacingDirection: 1,
					Percent:         44.47,
				},
				Post: slippi.PostFrameUpdate{
					FrameNumber:             715,
					PlayerIndex:             1,
					IsFollower:              false,
					CharacterID:             melee.Int_Marth,
					ActionStateID:           14,
					XPos:                    4.4829817,
					YPos:                    0.0001,
					FacingDirection:         1,
					Percent:                 44.47,
					ShieldSize:              60,
					LastHittingAttackID:     12,
					CurrentComboCount:       1,
					LastHitBy:               3,
					StocksRemaining:         4,
					ActionStateFrameCounter: 27,
					MiscActionState:         expectedPlayer1MiscState,
					IsAirborne:              false,
					LastGroundID:            3,
					JumpsRemaining:          2,
					LCancelStatus:           0,
					HurtboxCollisionState:   slippi.HurtboxStateVulnerable,
					HitlagRemaining:         0,
					AnimationIndex:          2,
					InstanceHitBy:           0,
					InstanceID:              0,
				},
			},
			2: {
				PlayerIndex: 2,
				Pre: slippi.PreFrameUpdate{
					FrameNumber:     715,
					PlayerIndex:     2,
					IsFollower:      false,
					RandomSeed:      4011085961,
					ActionStateID:   14,
					XPos:            11.230772,
					YPos:            0.0001,
					FacingDirection: -1,
					Percent:         42.05,
				},
				Post: slippi.PostFrameUpdate{
					FrameNumber:             715,
					PlayerIndex:             2,
					IsFollower:              false,
					CharacterID:             melee.Int_Samus,
					ActionStateID:           14,
					XPos:                    11.388272,
					YPos:                    0.0001,
					FacingDirection:         -1,
					Percent:                 42.05,
					ShieldSize:              60,
					LastHittingAttackID:     3,
					CurrentComboCount:       0,
					LastHitBy:               6,
					StocksRemaining:         4,
					ActionStateFrameCounter: 6,
					MiscActionState:         -0.73125,
					IsAirborne:              false,
					LastGroundID:            3,
					JumpsRemaining:          2,
					LCancelStatus:           0,
					HurtboxCollisionState:   slippi.HurtboxStateVulnerable,
					SelfInducedSpeeds: slippi.SelfInducedSpeeds{
						AirX:    -0.14249992,
						AirY:    0,
						AttackX: 0,
						AttackY: 0,
						GroundX: -0.14249992,
					},
					AnimationIndex: 2,
				},
			},
			3: {
				PlayerIndex: 3,
				Pre: slippi.PreFrameUpdate{
					FrameNumber:     715,
					PlayerIndex:     3,
					IsFollower:      false,
					RandomSeed:      4011085961,
					ActionStateID:   183,
					XPos:            10.692667,
					YPos:            0.0001,
					FacingDirection: -1,
					Percent:         7.6499996,
				},
				Post: slippi.PostFrameUpdate{
					FrameNumber:             715,
					PlayerIndex:             3,
					IsFollower:              false,
					CharacterID:             melee.Int_Roy,
					ActionStateID:           183,
					XPos:                    11.9542885,
					YPos:                    0.0001,
					FacingDirection:         -1,
					Percent:                 7.6499996,
					ShieldSize:              60,
					LastHittingAttackID:     2,
					CurrentComboCount:       1,
					LastHitBy:               0,
					StocksRemaining:         3,
					ActionStateFrameCounter: 21,
					MiscActionState:         31,
					IsAirborne:              true,
					LastGroundID:            3,
					JumpsRemaining:          1,
					LCancelStatus:           0,
					HurtboxCollisionState:   slippi.HurtboxStateVulnerable,
					SelfInducedSpeeds: slippi.SelfInducedSpeeds{
						AirX:    0,
						AirY:    0,
						AttackX: 1.2616214,
						AttackY: 0,
						GroundX: 0,
					},
					AnimationIndex: 183,
				},
			},
		}

		actualPlayers := actual.Players
		for i := uint8(0); i < 4; i++ {
			require.Contains(t, actualPlayers, i)
			expectedPlayer := expected[i]
			actualPlayer := actualPlayers[i]
			require.Equal(t, expectedPlayer, actualPlayer, "PlayerIndex %d not equal", i)
		}
	})

	t.Run("ParsesItemUpdates", func(t *testing.T) {
		t.Parallel()

		expected := []slippi.ItemUpdate{
			{
				FrameNumber:     715,
				ItemTypeID:      melee.ItemTargetMonster,
				State:           1,
				FacingDirection: 1,
				XVelocity:       0.3,
				YVelocity:       -0.57792664,
				XPos:            -76.59964,
				YPos:            49.37851,
				DamageTaken:     0,
				ExpirationTimer: 1400,
				SpawnID:         0,
				TurnipFace:      255,
				Owner:           -1,
				InstanceID:      0,
			},
			{
				FrameNumber:          715,
				ItemTypeID:           melee.ItemFoxBlaster,
				State:                4,
				FacingDirection:      1,
				XVelocity:            0,
				YVelocity:            0,
				XPos:                 -30.5598,
				YPos:                 18.6301,
				DamageTaken:          0,
				ExpirationTimer:      1400,
				SpawnID:              1,
				MissileType:          slippi.MissileType(4), // Garbage data
				TurnipFace:           slippi.TurnipFace(4),  // Garbage data
				ChargeShotIsLaunched: true,                  // Garbage data
				ChargeShotPower:      0,
				Owner:                0,
				InstanceID:           0,
			},
			{
				FrameNumber:          715,
				ItemTypeID:           melee.ItemFoxLaser,
				State:                0,
				FacingDirection:      1,
				XVelocity:            7,
				YVelocity:            0,
				XPos:                 4.2016563,
				YPos:                 40.63906,
				DamageTaken:          0,
				ExpirationTimer:      32,
				SpawnID:              2,
				MissileType:          slippi.MissileType(240), // Garbage data
				TurnipFace:           slippi.TurnipFace(0),    // Garbage data
				ChargeShotIsLaunched: false,
				ChargeShotPower:      91, // Garbage data
				Owner:                0,
				InstanceID:           0,
			},
		}
		require.Equal(t, expected, actual.ItemUpdates)
	})
}
