package ransomNote

import (
	"fmt"
	"strings"
)

func canConstruct(ransomNote string, magazine string) bool {
	ret := true
	for _, ransom := range ransomNote {
		if !strings.Contains(magazine, string(ransom)) {
			ret = false
			break
		}
		magazine = strings.Replace(magazine, string(ransom), "", 1)
	}
	return ret
}

func canConstructV2(ransomNote string, magazine string) bool {
	ret := true
	alphabet := make(map[rune]int)
	for _, m := range magazine {
		alphabet[m]++
	}
	for _, ransom := range ransomNote {
		alphabet[ransom]--
	}
	for _, value := range alphabet {
		if value < 0 {
			ret = false
		}
	}
	return ret
}

func main() {
	a := "aa"
	b := "ab"
	fmt.Print("result:", canConstruct(a, b))
}
