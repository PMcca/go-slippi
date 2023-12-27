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
	itemTypeID          = melee.ItemNessBat
	itemState           = uint8(123)
	itemFacingDirection = float32(1.23)
	itemXVel            = float32(12.9)
	itemYVel            = float32(-16.2)
	itemXPos            = float32(33.32)
	itemYPos            = float32(0.02)
	itemDamageTaken     = uint16(708)
	itemExpirationTimer = float32(79.4)
	itemSpawnID         = uint32(999)
	itemMissile         = slippi.MissileSuper
	itemTurnip          = slippi.TurnipStitchFace
	itemChargeLaunched  = uint8(1)
	itemChargePower     = uint8(200)
	itemOwner           = int8(4)
	itemInstanceID      = uint16(44)
)

func TestParseItemUpdate(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		frameNumber  int
		expected     slippi.Frame
		errAssertion require.ErrorAssertionFunc
	}{
		"CreatesItemUpdateForItem": {
			frameNumber: -123,
			expected: slippi.Frame{
				ItemUpdates: []slippi.ItemUpdate{
					{
						FrameNumber:          -123,
						ItemTypeID:           melee.ItemNessBat,
						State:                itemState,
						FacingDirection:      itemFacingDirection,
						XVelocity:            itemXVel,
						YVelocity:            itemYVel,
						XPos:                 itemXPos,
						YPos:                 itemYPos,
						DamageTaken:          itemDamageTaken,
						ExpirationTimer:      itemExpirationTimer,
						SpawnID:              itemSpawnID,
						MissileType:          itemMissile,
						TurnipFace:           itemTurnip,
						ChargeShotIsLaunched: true,
						ChargeShotPower:      itemChargePower,
						Owner:                itemOwner,
						InstanceID:           itemInstanceID,
					},
				},
				Players:   map[uint8]slippi.PlayerFrameUpdate{},
				Followers: map[uint8]slippi.PlayerFrameUpdate{},
			},

			errAssertion: require.NoError,
		},
	}

	for name, testCase := range testCases {
		tc := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			input := buildItemUpdateInput(int32(tc.frameNumber))
			dec := event.Decoder{
				Data: input,
				Size: len(input),
			}

			d := slippi.Data{}
			err := handlers.ItemUpdateHandler{}.Parse(&dec, &d)
			tc.errAssertion(t, err)

			actual := d.Frames[tc.frameNumber]
			require.Equal(t, tc.expected, actual)
		})
	}
}

func buildItemUpdateInput(frameNumber int32) []byte {
	out := []byte{byte(event.EventItemUpdate)}

	testutil.PutInt32(&out, frameNumber)
	testutil.PutUint16(&out, uint16(itemTypeID))
	out = append(out, itemState)
	testutil.PutFloat32(&out, itemFacingDirection)
	testutil.PutFloat32(&out, itemXVel)
	testutil.PutFloat32(&out, itemYVel)
	testutil.PutFloat32(&out, itemXPos)
	testutil.PutFloat32(&out, itemYPos)
	testutil.PutUint16(&out, itemDamageTaken)
	testutil.PutFloat32(&out, itemExpirationTimer)
	testutil.PutUint32(&out, itemSpawnID)
	out = append(out, byte(itemMissile))
	out = append(out, byte(itemTurnip))
	out = append(out, itemChargeLaunched) // Bool 1 = true
	out = append(out, itemChargePower)
	out = append(out, byte(itemOwner))
	testutil.PutUint16(&out, itemInstanceID)

	return out
}
