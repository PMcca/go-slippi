package goslippi

import (
	"bytes"
	"fmt"
	"github.com/PMcca/go-slippi/internal/errutil"
	"github.com/PMcca/go-slippi/internal/logging"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/PMcca/go-slippi/slippi/event"
	"github.com/PMcca/go-slippi/slippi/event/handler"
	"github.com/PMcca/go-slippi/slippi/event/handler/handlers"
	"github.com/toitware/ubjson"
	"os"
)

var (
	eventHandlers = map[event.Code]handler.EventHandler{
		event.EventGameStart:       handlers.GameStartHandler{},
		event.EventFrameStart:      handlers.FrameStartHandler{},
		event.EventPreFrame:        handlers.PreFrameHandler{},
		event.EventMessageSplitter: handlers.MessageSplitterHandler{},
		event.EventPostFrame:       handlers.PostFrameHandler{},
	}
	log = logging.NewLogger()
)

// rawParser contains the parsed Slippi replay and is used as the orchestrator in the parsing process.
type rawParser struct {
	ParsedData slippi.Data
}

// parser wraps a rawParser and slippi.Metadata, and is passed into ubjson.Unmarshal() to begin the parsing process.
type parser struct {
	RawParser rawParser       `ubjson:"raw"`
	Meta      slippi.Metadata `ubjson:"metadata"`
}

// ParseGame reads the .slp file given by filePath and returns the decoded game.
func ParseGame(filePath string) (slippi.Game, error) {
	b, err := readFile(filePath)
	if err != nil {
		return slippi.Game{}, err
	}

	p := parser{
		RawParser: rawParser{
			ParsedData: slippi.Data{
				Frames: map[int]slippi.Frame{},
			},
		},
	}
	if err := ubjson.Unmarshal(b, &p); err != nil {
		return slippi.Game{}, errutil.WithMessagef(err, ErrParsingGame, "filePath: %s", filePath)
	}

	return slippi.Game{
		Data: p.RawParser.ParsedData,
		Meta: p.Meta,
	}, nil
}

// UnmarshalUBJSON implements the ubjson.Unmarshaler interface. It receives the array of bytes from the 'raw' array and
// orchestrates the parsing process. rawParser implements this to separate this logic from slippi.Data.
func (r *rawParser) UnmarshalUBJSON(b []byte) error {
	// Beginning of raw array should always be '$U#l'.
	if !bytes.Equal(b[0:4], []byte("$U#l")) {
		return fmt.Errorf("%w:expected '$U#l', found %s", ErrInvalidRawStart, b[0:4]) // TODO move errors?
	}

	dec := event.Decoder{
		Data: b[8:],
		Size: len(b),
	}
	eventSizes, err := handlers.ParseEventPayloads(&dec) // Skip $U#l and 4 bytes for length.
	if err != nil {
		return err
	}

	startOffset := (eventSizes[event.EventPayloadsEvent] + 1) + 8 // Start reading from the first event after EventPayloads
	dec.Data = b[startOffset:]

	// Main event parsing loop
	for len(dec.Data) > 0 {
		eventCode := event.Code(dec.Read(0x0))
		eventSize, ok := eventSizes[eventCode]
		if !ok {
			return fmt.Errorf("%w:eventCode %X", ErrUnknownEventInEventSizes, eventCode)
		}
		dec.Size = eventSize + 1

		eventHandler, ok := eventHandlers[eventCode]
		if !ok {
			// TODO re-add log message when finished parsers
			//log.Warn().Msgf("Unable to handle unknown event %X. Skipping.", eventCode)
		} else {
			if err := eventHandler.Parse(&dec, &r.ParsedData); err != nil {
				return errutil.WithMessagef(err, ErrFailedEventParsing, "event code: %X", eventCode)
			}
		}

		// Update the window of data, skipping the # of bytes read + the command byte.
		dec.Data = dec.Data[dec.Size:]
	}

	return nil
}

// readFile reads & returns the bytes of the given .slp file.
func readFile(filePath string) ([]byte, error) {
	if filePath == "" {
		return nil, ErrEmptyFilePath
	}

	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errutil.WithMessagef(err, ErrReadingFile, "filePath: %s", filePath)
	}

	return b, nil
}
