func countSubstrings(s string) int {
    runes := []rune(s)
    countPalindromic := 0
    for i, v := range runes {
        countPalindromic++
        subStr := []rune{v}
        for j := i+1; j < len(runes); j++ {
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
