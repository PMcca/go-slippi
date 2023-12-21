package handlers

import (
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/event"
)

type GeckoCodeHandler struct{}

func (h GeckoCodeHandler) Parse(dec *event.Decoder, data *slippi.Data) error {
	var codes []slippi.GeckoCode
	i := 1
	for i < dec.Size {
		word1 := dec.ReadUint32(i)
		codeType := (word1 >> 24) & 0xfe
		address := (word1 & 0x01ffffff) + 0x80000000

		offset := uint32(8) // Default code length that applies to most codes.
		switch codeType {
		case 0xc0, 0xc2:
			lineCount := dec.ReadUint32(i + 4)
			offset = 8 + lineCount*8
		case 0x06:
			byteLen := dec.ReadUint32(i + 4)
			offset = 8 + (byteLen+7)&0xfffffff8
		case 0x08:
			offset = 16
		}

		codes = append(codes, slippi.GeckoCode{
			Type:     codeType,
			Address:  address,
			Contents: dec.ReadN(i, int(uint32(i)+offset)),
		})
		i += int(offset)
	}

	data.GeckoCodes = slippi.GeckoCodeList{
		Codes:    codes,
		Contents: dec.Data[1:],
	}
	return nil
}
