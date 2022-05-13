package slippi

import (
	"fmt"
	"github.com/jmank88/ubjson"
)

// Game represents a parsed .slp game.
type Game struct {
	Raw  []byte   `ubjson:"raw"`
	Meta Metadata `ubjson:"metadata"`
	//Meta map[string]interface{} `ubjson:"metadata"`
}

type Metadata struct {
	StartAt   string
	LastFrame int
	Players   map[string]interface{}
	PlayedOn  string
}

func (c *Metadata) UBJSONType() ubjson.Marker {
	return ubjson.ObjectStartMarker
}

func (c *Metadata) MarshalUBJSON(e *ubjson.Encoder) error {
	return nil
}

func (c *Metadata) UnmarshalUBJSON(d *ubjson.Decoder) error {
	o, err := d.Object()
	if err != nil {
		return err
	}

	for o.NextEntry() {
		k, err := o.DecodeKey()
		if err != nil {
			return err
		}

		switch k {
		case "startAt":
			s, err := o.DecodeString()
			if err != nil {
				return err
			}

			c.StartAt = s

		default:
			fmt.Printf("k was %s\n", k)
		}
	}

	return nil
}

type Game2 struct {
	FieldA string `ubjson:"fielda"`
	FieldB string
}

type Game3 struct {
	FieldC string `ubjson:"fielda"`
}
