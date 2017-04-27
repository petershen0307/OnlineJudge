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

func Test_canConstructV2(t *testing.T) {
	if b := canConstructV2("aa", "ab"); b != false {
		t.Error("Failed")
	}

	if b := canConstructV2("aa", "aab"); b != true {
		t.Error("Failed")
	}

	if b := canConstructV2("aaa", "aab"); b != false {
		t.Error("Failed")
	}
}
