package minimumaddtomakeparenthesesvalid

func minAddToMakeValid(s string) int {
	openCount := 0
	closeCount := 0
	for _, v := range s {
		if v == '(' {
			openCount++
		}
		if v == ')' {
			if openCount > 0 {
				openCount--
			} else {
				closeCount++
			}
		}
	}
	return openCount + closeCount
}
