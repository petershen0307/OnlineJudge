package validmountainarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	expect := false
	input := []int{2, 1}
	// act
	actual := validMountainArray(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	expect := false
	input := []int{3, 5, 5}
	// act
	actual := validMountainArray(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_3(t *testing.T) {
	// arrange
	expect := true
	input := []int{0, 3, 2, 1}
	// act
	actual := validMountainArray(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_4(t *testing.T) {
	// arrange
	expect := false
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// act
	actual := validMountainArray(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_5(t *testing.T) {
	// arrange
	expect := false
	input := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	// act
	actual := validMountainArray(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_6(t *testing.T) {
	// arrange
	expect := false
	input := []int{0}
	// act
	actual := validMountainArray(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_7(t *testing.T) {
	// arrange
	expect := false
	input := []int{1, 2}
	// act
	actual := validMountainArray(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_8(t *testing.T) {
	// arrange
	expect := true
	input := []int{1, 3, 2}
	// act
	actual := validMountainArray(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_9(t *testing.T) {
	// arrange
	expect := false
	input := []int{1, 1, 1, 1, 1, 1, 1, 2, 1}
	// act
	actual := validMountainArray(input)
	// assert
	assert.Equal(t, expect, actual)
}
