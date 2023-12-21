package melee

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_allItemsOrder(t *testing.T) {
	t.Parallel()
	t.Run("allItemsArrayIsInExpectedOrder", func(t *testing.T) {
		expected := []Item{
			MetalBox,
			CloakingDevice,
			PokeBall,
			UnknownItemBit4,
			UnknownItemBit5,
			UnknownItemBit6,
			UnknownItemBit7,
			UnknownItemBit8,
			Fan,
			FireFlower,
			SuperMushroom,
			PoisonMushroom,
			Hammer,
			WarpStar,
			ScrewAttack,
			BunnyHood,
			RayGun,
			Freezie,
			Food,
			MotionSensorBomb,
			Flipper,
			SuperScope,
			StarRod,
			LipsStick,
			HeartContainer,
			MaximTomato,
			Starman,
			HomeRunBat,
			BeamSword,
			Parasol,
			GreenShell,
			RedShell,
			Capsule,
			Box,
			Barrel,
			Egg,
			PartyBall,
			BarrelCannon,
			BobOmb,
			MrSaturn,
		}

		require.Equal(t, expected, allItems)
	})
}

func Test_concatenateBitfields(t *testing.T) {
	t.Parallel()
	t.Run("ConcatenatesItemBitfieldsIntoInt64", func(t *testing.T) {
		bitfield1 := uint8(0b10000001)
		bitfield2 := uint8(0b00110011)
		bitfield3 := uint8(0b10000000)
		bitfield4 := uint8(0b00000001)
		bitfield5 := uint8(0b01010101)

		expected := int64(0b0101010100000001100000000011001110000001)
		actual := concatenateBitfields(bitfield1, bitfield2, bitfield3, bitfield4, bitfield5)

		require.Equal(t, expected, actual)
	})
}

func TestGetEnabledItems(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		itemBitfield1 uint8
		itemBitfield2 uint8
		itemBitfield3 uint8
		itemBitfield4 uint8
		itemBitfield5 uint8
		expected      []Item
	}{
		"ReturnsAllEnabledItems": {
			itemBitfield1: 0b11111111,
			itemBitfield2: 0b11111111,
			itemBitfield3: 0b11111111,
			itemBitfield4: 0b11111111,
			itemBitfield5: 0b11111111,
			expected:      allItems,
		},
		"ReturnsNoItemsIfNoneEnabled": {
			itemBitfield1: 0b00000000,
			itemBitfield2: 0b00000000,
			itemBitfield3: 0b00000000,
			itemBitfield4: 0b00000000,
			itemBitfield5: 0b00000000,
			expected:      nil,
		},
		"ReturnsOnlyEnabledItems": {
			itemBitfield1: 0b00000101,
			itemBitfield2: 0b10000000,
			itemBitfield3: 0b00000110,
			itemBitfield4: 0b00010011,
			itemBitfield5: 0b10000001,
			expected: []Item{
				MetalBox,
				PokeBall,
				BunnyHood,
				Freezie,
				Food,
				HeartContainer,
				MaximTomato,
				BeamSword,
				Capsule,
				MrSaturn,
			},
		},
	}

	for name, testCase := range testCases {
		tc := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := GetEnabledItems(tc.itemBitfield1, tc.itemBitfield2, tc.itemBitfield3, tc.itemBitfield4, tc.itemBitfield5)
			require.Equal(t, tc.expected, actual)
		})
	}
}
