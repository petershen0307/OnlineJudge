package q1189

func maxNumberOfBalloons(text string) int {
	balloonCount := make(map[rune]int, 0)
	for _, char := range text {
		balloonCount[char]++
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	minVal := min(balloonCount['b'], balloonCount['a'])
	minVal = min(minVal, balloonCount['l']/2)
	minVal = min(minVal, balloonCount['o']/2)
	minVal = min(minVal, balloonCount['n'])
	return minVal
}
