package detectcapital

import "regexp"

func detectCapitalUse_regexp(word string) bool {
	re, _ := regexp.Compile(`^[A-Z]+$|^[a-z]+$|^[A-Z]{1}[a-z]+$`)
	return re.Match([]byte(word))
}

func detectCapitalUse(word string) bool {
	const none = 0
	const upper = 1
	const lower = 2
	state := none
	isUpper := func(w byte) bool {
		return w >= "A"[0] && w <= "Z"[0]
	}
	if isUpper(word[0]) {
		state = upper
	} else {
		state = lower
	}
	for i := 1; i < len(word); i++ {
		if state == upper && isUpper(word[i]) {
			continue
		}
		if state == upper && !isUpper(word[i]) && i == 1 {
			state = lower
		}
		if state == lower && !isUpper(word[i]) {
			continue
		}
		state = none
		break
	}
	return state != none
}
