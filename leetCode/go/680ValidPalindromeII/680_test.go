package validpalindromeii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	expect := true
	input := "aba"
	// act
	actual := validPalindrome(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	expect := true
	input := "abca"
	// act
	actual := validPalindrome(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_3(t *testing.T) {
	// arrange
	expect := false
	input := "abc"
	// act
	actual := validPalindrome(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_4(t *testing.T) {
	// arrange
	expect := true
	input := "abccba"
	// act
	actual := validPalindrome(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_5(t *testing.T) {
	// arrange
	expect := true
	input := "abcdba"
	// act
	actual := validPalindrome(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_6(t *testing.T) {
	// arrange
	expect := false
	input := "abcgdba"
	// act
	actual := validPalindrome(input)
	// assert
	assert.Equal(t, expect, actual)
}
