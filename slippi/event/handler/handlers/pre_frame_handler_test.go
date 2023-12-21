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
	preFrameSeed        = uint32(123456)
	preActionStateID    = uint16(345)
	preXpos             = float32(12.4)
	preYPos             = float32(17.12)
	preFacingDirection  = float32(2.0)
	preJoyStickX        = float32(0.88)
	preJoyStickY        = float32(0.1)
	preCStickX          = float32(0.22)
	preCStickY          = float32(0.5)
	preTrigger          = float32(0.01)
	preProcessedButtons = uint32(0b01010101010101010101010101010101)
	prePhysicalButtons  = uint16(0b1100110011001100)
	prePhysicalLTrigger = float32(0.234)
	prePhysicalRTrigger = float32(0.98)
	preAnalogUCFX       = int8(111)
	prePercent          = float32(84.1)
	preAnalogUCFY       = int8(12)
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
						PlayerIndex: 2,
						Pre: slippi.PreFrameUpdate{
							FrameNumber:      -123,
							PlayerIndex:      2,
							IsFollower:       false,
							RandomSeed:       preFrameSeed,
							ActionStateID:    preActionStateID,
							XPos:             preXpos,
							YPos:             preYPos,
							FacingDirection:  preFacingDirection,
							JoyStickX:        preJoyStickX,
							JoyStickY:        preJoyStickY,
							CStickX:          preCStickX,
							CStickY:          preCStickY,
							Trigger:          preTrigger,
							ProcessedButtons: preProcessedButtons,
							PhysicalButtons:  prePhysicalButtons,
							PhysicalTriggerL: prePhysicalLTrigger,
							PhysicalTriggerR: prePhysicalRTrigger,
							XAnalogUCF:       preAnalogUCFX,
							Percent:          prePercent,
							YAnalogUCF:       preAnalogUCFY,
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
						PlayerIndex: 1,
						Pre: slippi.PreFrameUpdate{
							FrameNumber:      -21,
							PlayerIndex:      1,
							IsFollower:       true,
							RandomSeed:       preFrameSeed,
							ActionStateID:    preActionStateID,
							XPos:             preXpos,
							YPos:             preYPos,
							FacingDirection:  preFacingDirection,
							JoyStickX:        preJoyStickX,
							JoyStickY:        preJoyStickY,
							CStickX:          preCStickX,
							CStickY:          preCStickY,
							Trigger:          preTrigger,
							ProcessedButtons: preProcessedButtons,
							PhysicalButtons:  prePhysicalButtons,
							PhysicalTriggerL: prePhysicalLTrigger,
							PhysicalTriggerR: prePhysicalRTrigger,
							XAnalogUCF:       preAnalogUCFX,
							Percent:          prePercent,
							YAnalogUCF:       preAnalogUCFY,
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
	testutil.PutUint16(&out, preActionStateID)
	testutil.PutFloat32(&out, preXpos)
	testutil.PutFloat32(&out, preYPos)
	testutil.PutFloat32(&out, preFacingDirection)
	testutil.PutFloat32(&out, preJoyStickX)
	testutil.PutFloat32(&out, preJoyStickY)
	testutil.PutFloat32(&out, preCStickX)
	testutil.PutFloat32(&out, preCStickY)
	testutil.PutFloat32(&out, preTrigger)
	testutil.PutUint32(&out, preProcessedButtons)
	testutil.PutUint16(&out, prePhysicalButtons)
	testutil.PutFloat32(&out, prePhysicalLTrigger)
	testutil.PutFloat32(&out, prePhysicalRTrigger)
	out = append(out, byte(preAnalogUCFX))
	testutil.PutFloat32(&out, prePercent)
	out = append(out, byte(preAnalogUCFY))

	return out
}
