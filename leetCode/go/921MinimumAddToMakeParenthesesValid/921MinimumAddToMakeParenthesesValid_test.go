package minimumaddtomakeparenthesesvalid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	expect := 1
	input := "())"
	// act
	actual := minAddToMakeValid(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	expect := 3
	input := "((("
	// act
	actual := minAddToMakeValid(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_3(t *testing.T) {
	// arrange
	expect := 4
	input := "()))(("
	// act
	actual := minAddToMakeValid(input)
	// assert
	assert.Equal(t, expect, actual)
}
