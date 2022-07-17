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

// String returns the character name of the respective internal character ID
func (c InternalCharacterID) String() string {
	switch c {
	case Mario:
		return "Mario"
	case Fox:
		return "Fox"
	case CaptainFalcon:
		return "Captain Falcon"
	case DonkeyKong:
		return "Donkey Kong"
	case Kirby:
		return "Kirby"
	case Bowser:
		return "Bowser"
	case Link:
		return "Link"
	case Sheik:
		return "Sheik"
	case Ness:
		return "Ness"
	case Peach:
		return "Peach"
	case Popo:
		return "Popo"
	case Nana:
		return "Nana"
	case Pikachu:
		return "Pikachu"
	case Samus:
		return "Samus"
	case Yoshi:
		return "Yoshi"
	case Jigglypuff:
		return "Jigglypuff"
	case Mewtwo:
		return "Mewtwo"
	case Luigi:
		return "Luigi"
	case Marth:
		return "Marth"
	case Zelda:
		return "Zelda"
	case YoungLink:
		return "Young Link"
	case DrMario:
		return "Dr. Mario"
	case Falco:
		return "Falco"
	case Pichu:
		return "Pichu"
	case GameAndWatch:
		return "Game & Watch"
	case Ganondorf:
		return "Ganondorf"
	case Roy:
		return "Roy"
	case MasterHand:
		return "Master Hand"
	case CrazyHand:
		return "Crazy Hand"
	case WireFrameMale:
		return "WireFrame Male"
	case WireFrameFemale:
		return "WireFrame Female"
	case GigaBowser:
		return "Giga Bowser"
	case Sandbag:
		return "Sandbag"
	default:
		return "unknown character"
	}
}

//// internalCharacterIDs holds a mapping of internal character ID -> character name.
//var internalCharacterIDs = map[int]string{
//	0:  "Mario",
//	1:  "Fox",
//	2:  "Captain Falcon",
//	3:  "Donkey Kong",
//	4:  "Kirby",
//	5:  "Bowser",
//	6:  "Link",
//	7:  "Sheik",
//	8:  "",
//	9:  "",
//	10: "",
//	11: "",
//	12: "",
//	13: "",
//	14: "",
//	15: "",
//	16: "",
//	17: "",
//	18: "",
//	19: "",
//	20: "",
//	21: "",
//	22: "",
//	23: "",
//	24: "",
//	25: "",
//	26: "",
//	27: "",
//	28: "",
//	29: "",
//	30: "",
//	31: "",
//	32: "",
//	33: "",
//}
