package algorithms

func CalculateVolforce(level uint, score uint, grade string, medal uint8) float32 {
	gradeBonus := *new(float32)
	medalBonus := *new(float32)

	switch grade {
	case "S":
		gradeBonus = 1.05
	case "AAA+":
		gradeBonus = 1.02
	case "AAA":
		gradeBonus = 1
	default:
		gradeBonus = 0.97 // AA+ and below
	}

	switch medal {
	case 5:
		medalBonus = 1.10
	case 4:
		medalBonus = 1.05
	case 3:
		medalBonus = 1.02
	case 2:
		medalBonus = 1
	case 1:
		medalBonus = 0.50
	default:
		medalBonus = 0.50
	}

	return (float32((level * 2)) * float32((score / 10000000)) * gradeBonus * medalBonus)
}
