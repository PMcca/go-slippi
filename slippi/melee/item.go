package melee

// Item represents an item which can be bitmasked against to check if it is enabled.
type Item int

const (
	MetalBox         Item = 1
	CloakingDevice   Item = 1 << 1
	PokeBall         Item = 1 << 2
	UnknownItemBit4  Item = 1 << 3
	UnknownItemBit5  Item = 1 << 4
	UnknownItemBit6  Item = 1 << 5
	UnknownItemBit7  Item = 1 << 6
	UnknownItemBit8  Item = 1 << 7
	Fan              Item = 1 << 8
	FireFlower       Item = 1 << 9
	SuperMushroom    Item = 1 << 10
	PoisonMushroom   Item = 1 << 11
	Hammer           Item = 1 << 12
	WarpStar         Item = 1 << 13
	ScrewAttack      Item = 1 << 14
	BunnyHood        Item = 1 << 15
	RayGun           Item = 1 << 16
	Freezie          Item = 1 << 17
	Food             Item = 1 << 18
	MotionSensorBomb Item = 1 << 19
	Flipper          Item = 1 << 20
	SuperScope       Item = 1 << 21
	StarRod          Item = 1 << 22
	LipsStick        Item = 1 << 23
	HeartContainer   Item = 1 << 24
	MaximTomato      Item = 1 << 25
	Starman          Item = 1 << 26
	HomeRunBat       Item = 1 << 27
	BeamSword        Item = 1 << 28
	Parasol          Item = 1 << 29
	GreenShell       Item = 1 << 30
	RedShell         Item = 1 << 31
	Capsule          Item = 1 << 32
	Box              Item = 1 << 33
	Barrel           Item = 1 << 34
	Egg              Item = 1 << 35
	PartyBall        Item = 1 << 36
	BarrelCannon     Item = 1 << 37
	BobOmb           Item = 1 << 38
	MrSaturn         Item = 1 << 39
)

// allItems is an in-order collection of all the items in Melee which are not character-specific (e.g. Link's bombs).
var allItems = []Item{
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

// GetEnabledItems takes the 5 ItemSpawnBitfields and returns their corresponding enabled items.
func GetEnabledItems(
	itemBitfield1,
	itemBitfield2,
	itemBitfield3,
	itemBitfield4,
	itemBitfield5 uint8) []Item {
	allBitfields := concatenateBitfields(itemBitfield1, itemBitfield2, itemBitfield3, itemBitfield4, itemBitfield5)

	var items []Item
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
