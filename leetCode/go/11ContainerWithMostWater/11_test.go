package containerwithmostwater

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	expect := 49
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// act
	actual := maxArea(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	expect := 1
	input := []int{1, 1}
	// act
	actual := maxArea(input)
	// assert
	assert.Equal(t, expect, actual)
}
