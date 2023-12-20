package handlers_test

import (
	"github.com/PMcca/go-slippi/internal/testutil"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/event"
	"github.com/PMcca/go-slippi/slippi/event/handler/handlers"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	preFrameSeed     = uint32(123456)
	actionStateID    = uint16(345)
	XPos             = float32(12.4)
	YPos             = float32(17.12)
	facingDirection  = float32(2.0)
	joyStickX        = float32(0.88)
	joyStickY        = float32(0.1)
	cStickX          = float32(0.22)
	cStickY          = float32(0.5)
	trigger          = float32(0.01)
	processedButtons = uint32(0b01010101010101010101010101010101)
	physicalButtons  = uint16(0b1100110011001100)
	physicalLTrigger = float32(0.234)
	physicalRTrigger = float32(0.98)
	analogUCFX       = int8(111)
	percent          = float32(84.1)
	analogUCFY       = int8(12)
)

func TestParsePreFrame(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		frameNumber  int
		playerIndex  uint8
		isFollower   bool
		expected     slippi.Frame
		errAssertion require.ErrorAssertionFunc
	}{
		"CreatesPreFrameForPlayer": {
			frameNumber: -123,
			playerIndex: 2,
			isFollower:  false,
			expected: slippi.Frame{
				Players: map[uint8]slippi.PlayerFrameUpdate{
					2: slippi.PlayerFrameUpdate{
						Pre: slippi.PreFrameUpdate{
							FrameNumber:      -123,
							PlayerIndex:      2,
							IsFollower:       false,
							RandomSeed:       preFrameSeed,
							ActionStateID:    actionStateID,
							XPos:             XPos,
							YPos:             YPos,
							FacingDirection:  facingDirection,
							JoyStickX:        joyStickX,
							JoyStickY:        joyStickY,
							CStickX:          cStickX,
							CStickY:          cStickY,
							Trigger:          trigger,
							ProcessedButtons: processedButtons,
							PhysicalButtons:  physicalButtons,
							PhysicalTriggerL: physicalLTrigger,
							PhysicalTriggerR: physicalRTrigger,
							XAnalogUCF:       analogUCFX,
							Percent:          percent,
							YAnalogUCF:       analogUCFY,
						},
					},
				},
				Followers: map[uint8]slippi.PlayerFrameUpdate{},
			},

			errAssertion: require.NoError,
		},
		"CreatesPreFrameForFollower": {
			frameNumber: -21,
			playerIndex: 1,
			isFollower:  true,
			expected: slippi.Frame{
				Followers: map[uint8]slippi.PlayerFrameUpdate{
					1: slippi.PlayerFrameUpdate{
						Pre: slippi.PreFrameUpdate{
							FrameNumber:      -21,
							PlayerIndex:      1,
							IsFollower:       true,
							RandomSeed:       preFrameSeed,
							ActionStateID:    actionStateID,
							XPos:             XPos,
							YPos:             YPos,
							FacingDirection:  facingDirection,
							JoyStickX:        joyStickX,
							JoyStickY:        joyStickY,
							CStickX:          cStickX,
							CStickY:          cStickY,
							Trigger:          trigger,
							ProcessedButtons: processedButtons,
							PhysicalButtons:  physicalButtons,
							PhysicalTriggerL: physicalLTrigger,
							PhysicalTriggerR: physicalRTrigger,
							XAnalogUCF:       analogUCFX,
							Percent:          percent,
							YAnalogUCF:       analogUCFY,
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

			input := buildPreFrameInput(int32(tc.frameNumber), tc.playerIndex, tc.isFollower)
			dec := event.Decoder{
				Data: input,
				Size: len(input),
			}

			d := slippi.Data{}
			err := handlers.PreFrameHandler{}.Parse(&dec, &d)
			tc.errAssertion(t, err)

			actual := d.Frames[tc.frameNumber]
			require.Equal(t, tc.expected, actual)
		})
	}
}

// buildPreFrameInput creates a valid PreFrame event used for parsing. The ordering of the fields is important here.
func buildPreFrameInput(frameNumber int32, playerIndex uint8, isFollower bool) []byte {
	out := []byte{byte(event.EventPreFrame)}
	follower := byte(0)
	if isFollower {
		follower = 1
	}

	testutil.PutInt32(&out, frameNumber)
	out = append(out, playerIndex)
	out = append(out, follower)
	testutil.PutUint32(&out, preFrameSeed)
	testutil.PutUint16(&out, actionStateID)
	testutil.PutFloat32(&out, XPos)
	testutil.PutFloat32(&out, YPos)
	testutil.PutFloat32(&out, facingDirection)
	testutil.PutFloat32(&out, joyStickX)
	testutil.PutFloat32(&out, joyStickY)
	testutil.PutFloat32(&out, cStickX)
	testutil.PutFloat32(&out, cStickY)
	testutil.PutFloat32(&out, trigger)
	testutil.PutUint32(&out, processedButtons)
	testutil.PutUint16(&out, physicalButtons)
	testutil.PutFloat32(&out, physicalLTrigger)
	testutil.PutFloat32(&out, physicalRTrigger)
	out = append(out, byte(analogUCFX))
	testutil.PutFloat32(&out, percent)
	out = append(out, byte(analogUCFY))

	return out
}
