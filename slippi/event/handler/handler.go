package handler

import (
	"github.com/pmcca/go-slippi/slippi"
	"github.com/pmcca/go-slippi/slippi/event"
)

// EventHandler defines the behaviour for parsing Slippi events.
type EventHandler interface {
	Parse(dec *event.Decoder, data *slippi.Data) error
}
