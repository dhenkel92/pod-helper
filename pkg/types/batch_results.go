package types

func BatchResults(batchSize int, results []Result) [][]Result {
	var divided [][]Result
	for i := 0; i < len(results); i += batchSize {
		end := i + batchSize

		if end > len(results) {
			end = len(results)
		}

		divided = append(divided, results[i:end])
	}
	return divided
}
