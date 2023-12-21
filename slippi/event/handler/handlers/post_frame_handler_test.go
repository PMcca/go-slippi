package handlers_test

import (
	"github.com/PMcca/go-slippi/internal/testutil"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/event"
	"github.com/PMcca/go-slippi/slippi/event/handler/handlers"
	"github.com/PMcca/go-slippi/slippi/melee"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	postActionStateID      = uint16(345)
	postXpos               = float32(12.4)
	postYPos               = float32(17.12)
	postFacingDirection    = float32(2.0)
	postPercent            = float32(84.1)
	postShieldSize         = float32(0.22)
	postLastHitAttackID    = uint8(2)
	postComboCount         = uint8(4)
	postLastHitBy          = uint8(1)
	postStocksRemaining    = uint8(3)
	postActionFrameCount   = float32(1.9)
	postMiscAction         = float32(9.1)
	postIsAirborne         = uint8(1) // bool for true
	postLastGroundID       = uint16(444)
	postJumpsRemaining     = uint8(1)
	postLCancelStatus      = uint8(1)
	postHurtboxState       = slippi.HurtboxStateIntangible
	postSelfInducedAirX    = float32(0.11)
	postSelfInducedAirY    = float32(0.22)
	postSelfInducedAttackX = float32(0.33)
	postSelfInducedAttackY = float32(0.44)
	postGroundX            = float32(0.66)
	postHitlagRemaining    = float32(11.22)
	postAnimationIndex     = uint32(987)
	postInstanceHitBy      = uint16(44)
	postInstanceID         = uint16(22)
)

func TestParsePostFrame(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		frameNumber int
		playerIndex uint8
		isFollower  bool
		characterID melee.InternalCharacterID

		expected     slippi.Frame
		errAssertion require.ErrorAssertionFunc
	}{
		"CreatesPostFrameForPlayer": {
			frameNumber: -22,
			playerIndex: 2,
			isFollower:  false,
			characterID: melee.Int_Falco,
			expected: slippi.Frame{
				Players: map[uint8]slippi.PlayerFrameUpdate{
					2: slippi.PlayerFrameUpdate{
						Post: slippi.PostFrameUpdate{
							FrameNumber:             -22,
							PlayerIndex:             2,
							IsFollower:              false,
							CharacterID:             melee.Int_Falco,
							ActionStateID:           postActionStateID,
							XPos:                    postXpos,
							YPos:                    postYPos,
							FacingDirection:         postFacingDirection,
							Percent:                 postPercent,
							ShieldSize:              postShieldSize,
							LastHittingAttackID:     postLastHitAttackID,
							CurrentComboCount:       postComboCount,
							LastHitBy:               postLastHitBy,
							StocksRemaining:         postStocksRemaining,
							ActionStateFrameCounter: postActionFrameCount,
							MiscActionState:         postMiscAction,
							IsAirborne:              true,
							LastGroundID:            postLastGroundID,
							JumpsRemaining:          postJumpsRemaining,
							LCancelStatus:           postLCancelStatus,
							HurtboxCollisionState:   postHurtboxState,
							SelfInducedSpeeds: slippi.SelfInducedSpeeds{
								AirX:    postSelfInducedAirX,
								AirY:    postSelfInducedAirY,
								AttackX: postSelfInducedAttackX,
								AttackY: postSelfInducedAttackY,
								GroundX: postGroundX,
							},
							HitlagRemaining: postHitlagRemaining,
							AnimationIndex:  postAnimationIndex,
							InstanceHitBy:   postInstanceHitBy,
							InstanceID:      postInstanceID,
						},
					},
				},
				Followers: map[uint8]slippi.PlayerFrameUpdate{},
			},

			errAssertion: require.NoError,
		},
		"CreatesPostFrameForFollower": {
			frameNumber: 924,
			playerIndex: 1,
			isFollower:  true,
			characterID: melee.Int_Falco,
			expected: slippi.Frame{
				Followers: map[uint8]slippi.PlayerFrameUpdate{
					1: slippi.PlayerFrameUpdate{
						Post: slippi.PostFrameUpdate{
							FrameNumber:             924,
							PlayerIndex:             1,
							IsFollower:              true,
							CharacterID:             melee.Int_Falco,
							ActionStateID:           postActionStateID,
							XPos:                    postXpos,
							YPos:                    postYPos,
							FacingDirection:         postFacingDirection,
							Percent:                 postPercent,
							ShieldSize:              postShieldSize,
							LastHittingAttackID:     postLastHitAttackID,
							CurrentComboCount:       postComboCount,
							LastHitBy:               postLastHitBy,
							StocksRemaining:         postStocksRemaining,
							ActionStateFrameCounter: postActionFrameCount,
							MiscActionState:         postMiscAction,
							IsAirborne:              true,
							LastGroundID:            postLastGroundID,
							JumpsRemaining:          postJumpsRemaining,
							LCancelStatus:           postLCancelStatus,
							HurtboxCollisionState:   postHurtboxState,
							SelfInducedSpeeds: slippi.SelfInducedSpeeds{
								AirX:    postSelfInducedAirX,
								AirY:    postSelfInducedAirY,
								AttackX: postSelfInducedAttackX,
								AttackY: postSelfInducedAttackY,
								GroundX: postGroundX,
							},
							HitlagRemaining: postHitlagRemaining,
							AnimationIndex:  postAnimationIndex,
							InstanceHitBy:   postInstanceHitBy,
							InstanceID:      postInstanceID,
						},
					},
				},
				Players: map[uint8]slippi.PlayerFrameUpdate{},
			},

			errAssertion: require.NoError,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			input := buildPostFrameInput(int32(tc.frameNumber), tc.playerIndex, tc.isFollower, tc.characterID)
			dec := event.Decoder{
				Data: input,
				Size: len(input),
			}

			d := slippi.Data{}
			err := handlers.PostFrameHandler{}.Parse(&dec, &d)
			tc.errAssertion(t, err)

			actual := d.Frames[tc.frameNumber]
			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestZeldaSheikFix(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		internalCharacterID melee.InternalCharacterID
		expected            slippi.Player
		errAssertion        require.ErrorAssertionFunc
	}{
		"SetsGameStartPlayerCharacterToSheikOnFirstFrame": {
			internalCharacterID: melee.Int_Sheik,
			expected: slippi.Player{
				Index:       0,
				CharacterID: melee.Ext_Sheik,
			},
			errAssertion: require.NoError,
		},
		"SetsGameStartPlayerCharacterToZeldaOnFirstFrame": {
			internalCharacterID: melee.Int_Zelda,
			expected: slippi.Player{
				Index:       0,
				CharacterID: melee.Ext_Zelda,
			},
			errAssertion: require.NoError,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			input := buildPostFrameInput(int32(-123), 0, false, tc.internalCharacterID)
			dec := event.Decoder{
				Data: input,
				Size: len(input),
			}

			d := slippi.Data{
				GameStart: slippi.GameStart{
					Players: []slippi.Player{
						{
							Index: 0,
						},
					},
				},
			}

			err := handlers.PostFrameHandler{}.Parse(&dec, &d)
			tc.errAssertion(t, err)
			require.Truef(t, len(d.GameStart.Players) > 0, "Length of players is empty")

			actual := d.GameStart.Players[0]
			require.Equal(t, tc.expected, actual)
		})
	}
}

func buildPostFrameInput(frameNumber int32, playerIndex uint8, isFollower bool, characterID melee.InternalCharacterID) []byte {
	out := []byte{byte(event.EventPreFrame)}
	follower := byte(0)
	if isFollower {
		follower = 1
	}

	testutil.PutInt32(&out, frameNumber)
	out = append(out, playerIndex)
	out = append(out, follower)
	out = append(out, byte(characterID))
	testutil.PutUint16(&out, postActionStateID)
	testutil.PutFloat32(&out, postXpos)
	testutil.PutFloat32(&out, postYPos)
	testutil.PutFloat32(&out, postFacingDirection)
	testutil.PutFloat32(&out, postPercent)
	testutil.PutFloat32(&out, postShieldSize)
	out = append(out, postLastHitAttackID)
	out = append(out, postComboCount)
	out = append(out, postLastHitBy)
	out = append(out, postStocksRemaining)
	testutil.PutFloat32(&out, postActionFrameCount)
	// Load 5 bytes for the State Bit Flags 1-5
	out = append(out, 0)
	out = append(out, 0)
	out = append(out, 0)
	out = append(out, 0)
	out = append(out, 0)
	testutil.PutFloat32(&out, postMiscAction)
	out = append(out, postIsAirborne)
	testutil.PutUint16(&out, postLastGroundID)
	out = append(out, postJumpsRemaining)
	out = append(out, postLCancelStatus)
	out = append(out, byte(postHurtboxState))
	testutil.PutFloat32(&out, postSelfInducedAirX)
	testutil.PutFloat32(&out, postSelfInducedAirY)
	testutil.PutFloat32(&out, postSelfInducedAttackX)
	testutil.PutFloat32(&out, postSelfInducedAttackY)
	testutil.PutFloat32(&out, postGroundX)
	testutil.PutFloat32(&out, postHitlagRemaining)
	testutil.PutUint32(&out, postAnimationIndex)
	testutil.PutUint16(&out, postInstanceHitBy)
	testutil.PutUint16(&out, postInstanceID)

	return out
}
