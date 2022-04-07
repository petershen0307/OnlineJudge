package laststoneweight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	expect := 1
	input := []int{2, 7, 4, 1, 8, 1}
	// act
	actual := lastStoneWeight(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	expect := 1
	input := []int{1}
	// act
	actual := lastStoneWeight(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_3(t *testing.T) {
	// arrange
	expect := 0
	input := []int{1, 1}
	// act
	actual := lastStoneWeight(input)
	// assert
	assert.Equal(t, expect, actual)
}
