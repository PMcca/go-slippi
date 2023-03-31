package melee

// InternalCharacterID is the Internal ID of the character which Melee uses.
type InternalCharacterID uint8

const (
	Mario InternalCharacterID = iota // = 0
	Fox
	CaptainFalcon
	DonkeyKong
	Kirby
	Bowser
	Link
	Sheik
	Ness
	Peach
	Popo
	Nana
	Pikachu
	Samus
	Yoshi
	Jigglypuff
	Mewtwo
	Luigi
	Marth
	Zelda
	YoungLink
	DrMario
	Falco
	Pichu
	GameAndWatch
	Ganondorf
	Roy
	MasterHand
	CrazyHand
	WireFrameMale
	WireFrameFemale
	GigaBowser
	Sandbag
)

var characterStrings = map[InternalCharacterID]string{
	Mario:           "Mario",
	Fox:             "Fox",
	CaptainFalcon:   "Captain Falcon",
	DonkeyKong:      "Donkey Kong",
	Kirby:           "Kirby",
	Bowser:          "Bowser",
	Link:            "Link",
	Sheik:           "Sheik",
	Ness:            "Ness",
	Peach:           "Peach",
	Popo:            "Popo",
	Nana:            "Nana",
	Pikachu:         "Pikachu",
	Samus:           "Samus",
	Yoshi:           "Yoshi",
	Jigglypuff:      "Jigglypuff",
	Mewtwo:          "Mewtwo",
	Luigi:           "Luigi",
	Marth:           "Marth",
	Zelda:           "Zelda",
	YoungLink:       "Young Link",
	DrMario:         "Dr. Mario",
	Falco:           "Falco",
	Pichu:           "Pichu",
	GameAndWatch:    "Mr. Game & Watch",
	Ganondorf:       "Ganondorf",
	Roy:             "Roy",
	MasterHand:      "Master Hand",
	CrazyHand:       "Crazy Hand",
	WireFrameMale:   "WireFrame Male",
	WireFrameFemale: "WireFrame Female",
	GigaBowser:      "Giga Bowser",
	Sandbag:         "Sandbag",
}

// String returns the character name of the respective internal character ID
func (c InternalCharacterID) String() string {
	n, ok := characterStrings[c]
	if !ok {
		return "unknown character"
	}
	return n
}
