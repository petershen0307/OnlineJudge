package searchinrotatedsortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	expect := 4
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	// act
	actual := search(nums, target)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	expect := -1
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3
	// act
	actual := search(nums, target)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_3(t *testing.T) {
	// arrange
	expect := -1
	nums := []int{1}
	target := 0
	// act
	actual := search(nums, target)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_4(t *testing.T) {
	// arrange
	expect := 4
	nums := []int{1, 2, 3, 4, 5}
	target := 5
	// act
	actual := search(nums, target)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_5(t *testing.T) {
	// arrange
	expect := 0
	nums := []int{1, 2, 3, 4, 5}
	target := 1
	// act
	actual := search(nums, target)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_6(t *testing.T) {
	// arrange
	expect := 1
	nums := []int{3, 1}
	target := 1
	// act
	actual := search(nums, target)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_7(t *testing.T) {
	// arrange
	expect := 0
	nums := []int{1, 3}
	target := 1
	// act
	actual := search(nums, target)
	// assert
	assert.Equal(t, expect, actual)
}
