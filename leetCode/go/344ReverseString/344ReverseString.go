package reversestring

func reverseString(s []byte) {
	for i := 0; i < len(s); i++ {
		lastOrderIndex := len(s) - i - 1
		if i > lastOrderIndex {
			break
		}
		s[i], s[lastOrderIndex] = s[lastOrderIndex], s[i]
	}
}
