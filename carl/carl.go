package carl

type Carl uint8

const (
	Herp Carl = iota
	Derp
	Merp
)

func (x Carl) String() string {
	switch x {
	case Herp:
		return "Herp"
	case Derp:
		return "Derp"
	case Merp:
		return "Merp"
	default:
		return "bleh"
	}
}

type Jimbob uint8

const (
	Kazerp Jimbob = iota
	Blargh
	Ugh
)
