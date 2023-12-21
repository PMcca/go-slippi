package slippi

// GeckoCode is a single Gecko Code in use.
type GeckoCode struct {
	Type     uint32
	Address  uint32
	Contents []uint8
}

// GeckoCodeList contains a list of GeckoCodes and all the contents of them.
type GeckoCodeList struct {
	Codes    []GeckoCode
	Contents []uint8
}
