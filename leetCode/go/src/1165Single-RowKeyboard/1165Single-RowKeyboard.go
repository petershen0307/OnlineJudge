package singlerowkeyboard

func calculateTime(keyboard string, word string) int {
	keyMap := make(map[rune]int, 26)
	for i, c := range keyboard {
		keyMap[c] = i
	}
	lastPosition := 0
	result := 0
	for _, c := range word {
		currentPosition := keyMap[c]
		if currentPosition > lastPosition {
			result += (currentPosition - lastPosition)
		} else {
			result += (lastPosition - currentPosition)
		}
		lastPosition = keyMap[c]
	}
	return result
}
