package slippi

import (
	"fmt"
	"github.com/jmank88/ubjson"
	"github.com/pkg/errors"
)

// Metadata represents the parsed metadata element from a .slp file.
type Metadata struct {
	StartAt   string                 `ubjson:"startAt"`
	LastFrame int32                  `ubjson:"lastFrame"`
	Players   map[string]interface{} `ubjson:"players"`
	PlayedOn  string                 `ubjson:"playedOn"`
}

type Player struct {
	Name string
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
			return errors.Wrap(err, "could not decode key")
		}

		switch k {
		case "startAt":
			s, err := o.DecodeString()
			if err != nil {
				return errors.Wrap(err, "could not decode string for startAt")
			}
			c.StartAt = s

		case "lastFrame":
			l, err := o.DecodeInt32()
			if err != nil {
				return errors.Wrap(err, "could not decode string for lastFrame")
			}
			c.LastFrame = l

		case "players":
			//err = o.DecodeObject(func(decoder *ubjson.ObjectDecoder) error {
			//	for i := 0; i < 4; i++ {
			//		a, err := decoder.DecodeKey()
			//		if err != nil {
			//			return errors.Wrap(err, "could not decode key for inside players func")
			//		}
			//
			//		switch a {
			//		case "0","1","2","3":
			//			fmt.Println("key is ", a)
			//
			//		}
			//	}
			//	return nil
			//})
			//if err != nil {
			//	return errors.Wrap(err, "could not decode players object")
			//}

			players := make(map[string]interface{})
			err = o.Decode(&players)
			if err != nil {
				return errors.Wrap(err, "could not decode players into map")
			}
			//p, err := o.DecodeString()
			//if err != nil {
			//	return err
			//}

			c.Players = players

		case "characters":
			s, err := o.DecodeString()
			if err != nil {
				return err
			}

			fmt.Println(s)

		case "playedOn":
			s, err := o.DecodeString()
			if err != nil {
				return errors.Wrap(err, "could not decode string for playedOn")
			}
			c.PlayedOn = s

		default:
			fmt.Printf("k was %s\n", k)
		}
	}

	return nil
}
