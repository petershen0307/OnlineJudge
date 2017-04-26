package ransomNote

import "testing"

func Test_canConstruct(t *testing.T) {
	if b := canConstruct("aa", "ab"); b != false {
		t.Error("Failed")
	}

	if b := canConstruct("aa", "aab"); b != true {
		t.Error("Failed")
	}

	if b := canConstruct("aaa", "aab"); b != false {
		t.Error("Failed")
	}
}
