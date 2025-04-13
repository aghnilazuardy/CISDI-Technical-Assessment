package helper

func CalculateRank(records []int, scores []int) []int {
	uniqueScores := removeDuplicatesAndSort(records)

	result := make([]int, len(scores))

	// every scored, calculate the rank
	for i, score := range scores {
		rank := 1
		for j := 0; j < len(uniqueScores); j++ {
			if score < uniqueScores[j] {
				rank++
			} else {
				break
			}
		}
		result[i] = rank
	}

	return result
}

func removeDuplicatesAndSort(scores []int) []int {
	// remove duplicate
	scoreMap := make(map[int]bool)
	for _, score := range scores {
		scoreMap[score] = true
	}

	// convert map to slice
	uniqueScores := make([]int, 0, len(scoreMap))
	for score := range scoreMap {
		uniqueScores = append(uniqueScores, score)
	}

	// sort ranking descending
	for i := 0; i < len(uniqueScores); i++ {
		for j := i + 1; j < len(uniqueScores); j++ {
			if uniqueScores[i] < uniqueScores[j] {
				uniqueScores[i], uniqueScores[j] = uniqueScores[j], uniqueScores[i]
			}
		}
	}

	return uniqueScores
}
