package slippi

type GeckoCode struct {
	Type     uint32
	Address  uint32
	Contents []uint8
}

type GeckoCodeList struct {
	Codes    []GeckoCode
	Contents []uint8
}
