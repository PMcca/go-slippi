package handler

import (
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/event"
)

type EventHandler interface {
	Parse(dec *event.Decoder, data *slippi.Data) error
}
