package palindromicsubstrings

// O(N^3)
func countSubstringsOn3(s string) int {
	runes := []rune(s)
	countPalindromic := 0
	for i, v := range runes {
		countPalindromic++
		subStr := []rune{v}
		for j := i + 1; j < len(runes); j++ {
			subStr = append(subStr, runes[j])
			if isPalindromic(subStr) {
				countPalindromic++
			}
		}
	}
	return countPalindromic
}

func isPalindromic(runes []rune) bool {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

// O(N^2)
func countSubstringsOn2(s string) int {
	runes := []rune(s)
	countPalindromic := 0
	for i := range runes {
		countPalindromic += extendPalindromic(runes, i, i)   // odd palindromic
		countPalindromic += extendPalindromic(runes, i, i+1) // event palindromic
	}
	return countPalindromic
}

func extendPalindromic(runes []rune, i, j int) int {
	count := 0
	for ; i >= 0 && j < len(runes); i, j = i-1, j+1 {
		if runes[i] == runes[j] {
			count++
		} else {
			break
		}
	}
	return count
}
