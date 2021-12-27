package basiccalculatorii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	input := "3+2*2"
	expect := 7
	// act
	actual := calculate(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	input := "3-2/12"
	expect := 3
	// act
	actual := calculate(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_3(t *testing.T) {
	// arrange
	input := "3*2-12"
	expect := -6
	// act
	actual := calculate(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_4(t *testing.T) {
	// arrange
	input := " 3 * 2 - 12 "
	expect := -6
	// act
	actual := calculate(input)
	// assert
	assert.Equal(t, expect, actual)
}
