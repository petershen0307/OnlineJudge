package lettercombinationsofaphonenumber

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	digitKeyMap := map[rune][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
	result := []string{}
	for _, v0 := range digitKeyMap[rune(digits[0])] {
		if len(digits) == 1 {
			result = append(result, v0)
		}
		for i1 := 0; len(digits) > 1 && i1 < len(digitKeyMap[rune(digits[1])]); i1++ {
			if len(digits) == 2 {
				result = append(result, v0+digitKeyMap[rune(digits[1])][i1])
			}
			for i2 := 0; len(digits) > 2 && i2 < len(digitKeyMap[rune(digits[2])]); i2++ {
				if len(digits) == 3 {
					result = append(result, v0+digitKeyMap[rune(digits[1])][i1]+digitKeyMap[rune(digits[2])][i2])
				}
				for i3 := 0; len(digits) > 3 && i3 < len(digitKeyMap[rune(digits[3])]); i3++ {
					result = append(result, v0+digitKeyMap[rune(digits[1])][i1]+digitKeyMap[rune(digits[2])][i2]+digitKeyMap[rune(digits[3])][i3])

				}
			}
		}
	}
	return result
}
