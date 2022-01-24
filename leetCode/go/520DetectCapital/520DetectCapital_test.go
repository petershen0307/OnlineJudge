package detectcapital

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	expect := true
	input := "USA"
	// act
	actual := detectCapitalUse(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	expect := false
	input := "FlaG"
	// act
	actual := detectCapitalUse(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_3(t *testing.T) {
	// arrange
	expect := true
	input := "Google"
	// act
	actual := detectCapitalUse(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_4(t *testing.T) {
	// arrange
	expect := true
	input := "leetcode"
	// act
	actual := detectCapitalUse(input)
	// assert
	assert.Equal(t, expect, actual)
}
