package algorithms

func CalculateAccuracy(crits, nears, errors, maxNotes uint32) float64 {
	// HACK: maxNotes is 0 until we can confidently determine it's number
	maxNotes = (crits + nears + errors)

	critAmt := float64(100 / maxNotes)
	nearAmt := float64(critAmt / 2)
	return (critAmt * float64(crits)) + (nearAmt * float64(nears))
}
