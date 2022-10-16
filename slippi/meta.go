package slippi

import (
	"fmt"
	"github.com/PMcca/go-slippi/internal/sentinel"
	"github.com/PMcca/go-slippi/slippi/melee"
	"github.com/jmank88/ubjson"
	"github.com/pkg/errors"
	"log"
	"strconv"
)

// Metadata represents the parsed metadata element from a .slp file.
type Metadata struct {
	StartAt   string    `ubjson:"startAt"`
	LastFrame int       `ubjson:"lastFrame"`
	Players   []*Player `ubjson:"players"`
	PlayedOn  string    `ubjson:"playedOn"`
}

// Character is the Melee character that was present in the match, including how long said player was played.
type Character struct {
	CharacterID  melee.InternalCharacterID
	FramesPlayed int
}

// Player is a single player in the game including their Slippi display name or in-game name (dependent on online/local).
type Player struct {
	Name       Names `ubjson:"names"`
	Port       int
	Characters []Character
}

type Names struct {
	Name       string `ubjson:"netplay"`
	SlippiCode string `ubjson:"code"`
}

func (m *Metadata) UBJSONType() ubjson.Marker {
	return ubjson.ObjectStartMarker
}

func (m *Metadata) MarshalUBJSON(e *ubjson.Encoder) error {
	//TODO: Support encoding later?
	return nil
}

func (m *Metadata) UnmarshalUBJSON(d *ubjson.Decoder) error {
	tempM := *m

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
				return sentinel.Wrap(err, ErrDecodingField)
			}
			tempM.StartAt = s

		case "lastFrame":
			l, err := o.DecodeInt()
			if err != nil {
				return sentinel.Wrap(err, ErrDecodingField)
				//return errors.Wrap(err, "could not decode string for lastFrame")
			}
			tempM.LastFrame = l

		case "players":
			err := o.DecodeObject(parsePlayers(&tempM))
			if err != nil {
				return errors.Wrap(err, "could not parse players")
			}
			//players := make(map[string]interface{})
			//err = o.Decode(&players)
			//if err != nil {
			//	return errors.Wrap(err, "could not decode players into map")
			//}

		case "playedOn":
			s, err := o.DecodeString()
			if err != nil {
				return errors.Wrap(err, "could not decode string for playedOn")
			}
			tempM.PlayedOn = s

		default:
			log.Printf("Unexpected key in meta: %s\n", k)
		}
	}

	*m = tempM
	return nil
}

// parsePlayers parses the Players object given by the ObjectDecoder and sets it to the Metadata's Players field.
func parsePlayers(m *Metadata) func(decoder *ubjson.ObjectDecoder) error {
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

				player := Player{
					Port: port,
				}

				// Parse this player's fields
				err = o.DecodeObject(func(d *ubjson.ObjectDecoder) error {
					for d.NextEntry() {
						k, err := d.DecodeKey()
						if err != nil {
							return errors.Wrap(err, "could not decode key in players object")
						}

						switch k {
						case "characters":
							err = d.DecodeObject(parseCharacters(&player))
							if err != nil {
								return errors.Wrap(err, "could not decode characters object")
							}

						case "names":
							names := Names{}
							err = d.Decode(&names)
							if err != nil {
								return errors.Wrap(err, "could not decode names")
							}
							player.Name = names
						}
					}
					return d.End()
				})
				if err != nil {
					return errors.Wrap(err, "could not parse characters object")
				}

				m.Players = append(m.Players, &player)

			default:
				return errors.New(fmt.Sprintf("unknown key in players object. expected port, got %s", p))
			}
		}

		return o.End()
	}
}

// parseCharacters parses the characters UBJSON object in the players parent object.
func parseCharacters(player *Player) func(decoder *ubjson.ObjectDecoder) error {
	return func(o *ubjson.ObjectDecoder) error {
		for o.NextEntry() {
			id, err := o.DecodeKey()
			if err != nil {
				return errors.Wrap(err, "could not decode characterID key")
			}

			characterID, err := strconv.Atoi(id)
			if err != nil {
				return errors.Wrap(err, "could not convert characterID to int")
			}

			framesPlayed, err := o.DecodeInt()
			if err != nil {
				return errors.Wrap(err, "could not decode frames played")
			}

			player.Characters = append(player.Characters, Character{
				CharacterID:  melee.InternalCharacterID(characterID),
				FramesPlayed: framesPlayed,
			})
		}

		return o.End()
	}
}
