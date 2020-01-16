package plates

var barbellWeight = 45.0

func plateCalc(totalWeight float64) []float64 {
	adjustedWeight := totalWeight - barbellWeight
	if adjustedWeight < 0 {
		return []float64{}
	}

	plates := []float64{}
	remainingWeight := adjustedWeight / 2
	for remainingWeight > 0 {
		plate := getLargestAvailablePlate(remainingWeight)
		plates = append(plates, plate)

		remainingWeight = remainingWeight - plate
	}

	return plates
}

var availableWeights = []float64{45, 35, 25, 15, 10, 5, 2.5, 1.25}

func getLargestAvailablePlate(target float64) float64 {
	for _, weight := range availableWeights {
		diff := target - weight
		if diff > 0 || diff == 0 {
			return weight
		}
	}

	return 0
}
