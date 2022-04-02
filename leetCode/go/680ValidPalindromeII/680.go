package validpalindromeii

func validPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			delete_j := s[i:j]
			delete_i := s[i+1 : j+1]
			return isPalindrome(delete_i) || isPalindrome(delete_j)
		}
	}
	return true
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
