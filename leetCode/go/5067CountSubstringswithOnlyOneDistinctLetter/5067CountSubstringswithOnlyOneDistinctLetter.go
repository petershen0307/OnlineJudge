package countsubstringswithonlyonedistinctletter

func countLetters(S string) int {
	previousRune := ' '
	sameRuneCount := 0
	result := 0
	addToResult := func(tmp int) {
		if tmp%2 == 0 {
			result += (tmp + 1) * (tmp / 2)
		} else {
			result += ((tmp + 1) / 2) * tmp
		}
	}
	for i, a := range S {
		sameRuneCount++
		if previousRune == ' ' {
			previousRune = a
		}
		if previousRune != a {
			// calcute sameRuneCount
			previousRune = a
			addToResult(sameRuneCount - 1)
			sameRuneCount = 1
		}
		if i == len(S)-1 {
			addToResult(sameRuneCount)
		}
	}
	return result
}
