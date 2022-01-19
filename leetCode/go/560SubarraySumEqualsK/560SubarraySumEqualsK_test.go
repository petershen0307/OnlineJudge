package subarraysumequalsk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	expect := 2
	k := 2
	input := []int{1, 1, 1}
	// act
	actual := subarraySum(input, k)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	expect := 2
	k := 3
	input := []int{1, 2, 3}
	// act
	actual := subarraySum(input, k)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_3(t *testing.T) {
	// arrange
	expect := 1
	k := 0
	input := []int{-1, -1, 1}
	// act
	actual := subarraySum(input, k)
	// assert
	assert.Equal(t, expect, actual)
}
