package algorithms

func CalculateAccuracy(crits, nears, errors, maxNotes uint32) float64 {
	// HACK: maxNotes is 0 until we can confidently determine it's number
	maxNotes = (crits + nears + errors)
	if maxNotes == 0 {
		maxNotes = 1 // HACK: maxNotes also must be non-zero in the case that everything else is also 0
	}

	critAmt := float64(float64(100) / float64(maxNotes))
	nearAmt := float64(critAmt / 2)
	return (critAmt * float64(crits)) + (nearAmt * float64(nears))
}
