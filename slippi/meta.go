package slippi

import (
	"fmt"
	"github.com/PMcca/go-slippi/slippi/melee"
	"github.com/jmank88/ubjson"
	"github.com/pkg/errors"
	"strconv"
)

// Metadata represents the parsed metadata element from a .slp file.
type Metadata struct {
	StartAt   string                 `ubjson:"startAt"`
	LastFrame int32                  `ubjson:"lastFrame"`
	Players   map[string]interface{} `ubjson:"players"`
	PlayedOn  string                 `ubjson:"playedOn"`
}

// Character is the Melee character that was present in the match, including how long said player was played.
type Character struct {
	CharacterID  melee.InternalCharacterID
	FramesPlayed int
}

// Player is a single player in the game including their Slippi display name or in-game name (dependent on online/local).
type Player struct {
	Name       string
	Port       int
	Characters []Character
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
			//err = o.DecodeObject(parsePlayers(c, d))
			//if err != nil {
			//	return err
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

// parsePlayers parses the Players object given by the ObjectDecoder and sets it to the Metadata's Players field.
func parsePlayers(m *Metadata, d *ubjson.Decoder) func(decoder *ubjson.ObjectDecoder) error {
	return func(o *ubjson.ObjectDecoder) error {
		for o.NextEntry() {
			p, err := o.DecodeKey()
			if err != nil {
				return errors.Wrap(err, "could not decode players key")
			}

			switch p {
			case "0", "1", "2", "3":
				port, err := strconv.Atoi(p)
				if err != nil {
					return errors.WithMessagef(err, "could not convert port  %s to int", p)
				}
				fmt.Println(port)

			default:
				return errors.New(fmt.Sprintf("unknown key in players object. expected port, got %s", p))
			}
		}

		return nil
	}
}
