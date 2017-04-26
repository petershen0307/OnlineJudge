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

func main() {
	a := "aa"
	b := "ab"
	fmt.Print("result:", canConstruct(a, b))
}
