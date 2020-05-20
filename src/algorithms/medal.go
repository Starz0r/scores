package algorithms

// TODO: change this in for go 1.13 (https://github.com/golang/go/issues/19308)
const (
	gameFlagNone uint32 = iota
	gameFlagHard        = 0x1
)

func CalculateMedal(score uint64, errors uint64, gauge float32, flags uint32) uint8 {
	switch {
	case score == 10000000: // Perfect
		return 5
	case errors == 0: // Full Combo
		return 4
	case (flags&gameFlagHard) != gameFlagNone && gauge > 0: // Hard Clear
		return 3
	case (flags&gameFlagHard) == gameFlagNone && gauge > 0: // Clear
		return 2
	default: // Failed
		return 1
	}
}
