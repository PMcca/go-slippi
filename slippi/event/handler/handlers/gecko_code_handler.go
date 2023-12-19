package handlers

import (
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/event"
)

type GeckoCodeHandler struct{}

func (g GeckoCodeHandler) Parse(dec *event.Decoder, data *slippi.Data) error {
	// TODO implement Gecko parsing
	return nil
}
