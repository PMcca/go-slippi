package handlers

import (
	"github.com/pmcca/go-slippi/slippi"
	"github.com/pmcca/go-slippi/slippi/event"
)

// GameEndHandler handles the parsing of GameEnd events.
type GameEndHandler struct{}

// Parse implements the handler.EventHandler interface. It parses a GameEnd event and puts its output into the
// given slippi.Data struct.
func (h GameEndHandler) Parse(dec *event.Decoder, data *slippi.Data) error {
	var placements []slippi.PlayerPlacement
	for i := 0; i < 4; i++ {
		placements = append(placements, slippi.PlayerPlacement{
			PlayerIndex: int8(i),
			Placement:   dec.ReadInt8(0x3 + i),
		})
	}

	data.GameEnd = slippi.GameEnd{
		GameEndMethod:    slippi.GameEndMethod(dec.Read(0x1)),
		LRASInitiatior:   dec.ReadInt8(0x2),
		PlayerPlacements: placements,
	}
	return nil
}
