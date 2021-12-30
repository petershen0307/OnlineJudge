package minimumnumberofkconsecutivebitflips

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	expect := 2
	k := 1
	input := []int{0, 1, 0}
	// act
	actual := minKBitFlips(input, k)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	expect := -1
	k := 2
	input := []int{1, 1, 0}
	// act
	actual := minKBitFlips(input, k)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_3(t *testing.T) {
	// arrange
	expect := 3
	k := 3
	input := []int{0, 0, 0, 1, 0, 1, 1, 0}
	// act
	actual := minKBitFlips(input, k)
	// assert
	assert.Equal(t, expect, actual)
}
