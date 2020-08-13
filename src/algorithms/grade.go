package algorithms

func CalculateGrade(score uint32) string {
	switch {
	case score >= 9900000:
		return "S"
	case score >= 9800000:
		return "AAA+"
	case score >= 9700000:
		return "AAA"
	case score >= 9500000:
		return "AA+"
	case score >= 9300000:
		return "AA"
	case score >= 9000000:
		return "A+"
	case score >= 8700000:
		return "A"
	case score >= 7500000:
		return "B"
	case score >= 6500000:
		return "C"
	default:
		return "D"
	}
}
