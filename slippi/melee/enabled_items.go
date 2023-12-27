package melee

// EnabledItem represents an item which can be enabled, found in the item spawn bitfields.
type EnabledItem int

const (
	EnabledMetalBox         EnabledItem = 1
	EnabledCloakingDevice   EnabledItem = 1 << 1
	EnabledPokeBall         EnabledItem = 1 << 2
	EnabledUnknownItemBit4  EnabledItem = 1 << 3
	EnabledUnknownItemBit5  EnabledItem = 1 << 4
	EnabledUnknownItemBit6  EnabledItem = 1 << 5
	EnabledUnknownItemBit7  EnabledItem = 1 << 6
	EnabledUnknownItemBit8  EnabledItem = 1 << 7
	EnabledFan              EnabledItem = 1 << 8
	EnabledFireFlower       EnabledItem = 1 << 9
	EnabledSuperMushroom    EnabledItem = 1 << 10
	EnabledPoisonMushroom   EnabledItem = 1 << 11
	EnabledHammer           EnabledItem = 1 << 12
	EnabledWarpStar         EnabledItem = 1 << 13
	EnabledScrewAttack      EnabledItem = 1 << 14
	EnabledBunnyHood        EnabledItem = 1 << 15
	EnabledRayGun           EnabledItem = 1 << 16
	EnabledFreezie          EnabledItem = 1 << 17
	EnabledFood             EnabledItem = 1 << 18
	EnabledMotionSensorBomb EnabledItem = 1 << 19
	EnabledFlipper          EnabledItem = 1 << 20
	EnabledSuperScope       EnabledItem = 1 << 21
	EnabledStarRod          EnabledItem = 1 << 22
	EnabledLipsStick        EnabledItem = 1 << 23
	EnabledHeartContainer   EnabledItem = 1 << 24
	EnabledMaximTomato      EnabledItem = 1 << 25
	EnabledStarman          EnabledItem = 1 << 26
	EnabledHomeRunBat       EnabledItem = 1 << 27
	EnabledBeamSword        EnabledItem = 1 << 28
	EnabledParasol          EnabledItem = 1 << 29
	EnabledGreenShell       EnabledItem = 1 << 30
	EnabledRedShell         EnabledItem = 1 << 31
	EnabledCapsule          EnabledItem = 1 << 32
	EnabledBox              EnabledItem = 1 << 33
	EnabledBarrel           EnabledItem = 1 << 34
	EnabledEgg              EnabledItem = 1 << 35
	EnabledPartyBall        EnabledItem = 1 << 36
	EnabledBarrelCannon     EnabledItem = 1 << 37
	EnabledBobOmb           EnabledItem = 1 << 38
	EnabledMrSaturn         EnabledItem = 1 << 39
)

// allItems is an in-order collection of all the items in Melee which are not character-specific (e.g. Link's bombs).
var allItems = []EnabledItem{
	EnabledMetalBox,
	EnabledCloakingDevice,
	EnabledPokeBall,
	EnabledUnknownItemBit4,
	EnabledUnknownItemBit5,
	EnabledUnknownItemBit6,
	EnabledUnknownItemBit7,
	EnabledUnknownItemBit8,
	EnabledFan,
	EnabledFireFlower,
	EnabledSuperMushroom,
	EnabledPoisonMushroom,
	EnabledHammer,
	EnabledWarpStar,
	EnabledScrewAttack,
	EnabledBunnyHood,
	EnabledRayGun,
	EnabledFreezie,
	EnabledFood,
	EnabledMotionSensorBomb,
	EnabledFlipper,
	EnabledSuperScope,
	EnabledStarRod,
	EnabledLipsStick,
	EnabledHeartContainer,
	EnabledMaximTomato,
	EnabledStarman,
	EnabledHomeRunBat,
	EnabledBeamSword,
	EnabledParasol,
	EnabledGreenShell,
	EnabledRedShell,
	EnabledCapsule,
	EnabledBox,
	EnabledBarrel,
	EnabledEgg,
	EnabledPartyBall,
	EnabledBarrelCannon,
	EnabledBobOmb,
	EnabledMrSaturn,
}

// GetEnabledItems takes the 5 ItemSpawnBitfields and returns their corresponding enabled items.
func GetEnabledItems(
	itemBitfield1,
	itemBitfield2,
	itemBitfield3,
	itemBitfield4,
	itemBitfield5 uint8) []EnabledItem {
	allBitfields := concatenateBitfields(itemBitfield1, itemBitfield2, itemBitfield3, itemBitfield4, itemBitfield5)

	var items []EnabledItem
	for _, item := range allItems {
		if int64(item)&allBitfields > 0 {
			items = append(items, item)
		}
	}

	return items
}

// concatenateBitfields takes all 5 item spawn bitfields and concatenates the bits to produce one 5-byte number
// representing each item bitfield. ItemBitfield1 starts from the LSB, and ItemBitfield5 ends at the MSB.
func concatenateBitfields(bf1, bf2, bf3, bf4, bf5 uint8) int64 {
	result := int64(bf5) // ItemBitfield 5 will be at the end of the binary integer.
	result <<= 8
	result += int64(bf4)
	result <<= 8
	result += int64(bf3)
	result <<= 8
	result += int64(bf2)
	result <<= 8
	result += int64(bf1) // ItemBitfield 1 will be the start of the binary integer.

	return result
}
