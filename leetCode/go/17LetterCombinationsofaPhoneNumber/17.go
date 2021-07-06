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
	combination(digitKeyMap, digits, &result)
	return result
}

func combination(digitKeyMap map[rune][]string, digits string, result *[]string) {
	if digits == "" {
		return
	}
	combination(digitKeyMap, digits[0:len(digits)-1], result)
	temp := []string{}
	for _, v := range digitKeyMap[rune(digits[len(digits)-1])] {
		if len(digits) == 1 {
			temp = append(temp, v)
		} else {
			for _, r := range *result {
				temp = append(temp, r+v)
			}
		}
	}
	*result = temp
}
