package helpers

// Word generator generates words from letters in the characters array using nth Cartesan Product
func WordGen(input string, n int) []string {
	if n <= 0 {
		return nil
	}

	prod := make([]string, len(input))
	for i, char := range input {
		prod[i] = string(char)
	}

	for i := 1; i < n; i++ {
		next := make([]string, 0, len(input)*len(prod))
		for _, word := range prod {
			for _, char := range input {
				next = append(next, word+string(char))
			}
		}

		prod = next
	}

	return prod
}
