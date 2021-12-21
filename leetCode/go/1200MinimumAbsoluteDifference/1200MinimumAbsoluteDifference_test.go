package minimumabsolutedifference

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	expect := [][]int{{1, 2}, {2, 3}, {3, 4}}
	arr := []int{4, 2, 1, 3}
	// act
	actual := minimumAbsDifference(arr)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	expect := [][]int{{1, 3}}
	arr := []int{1, 3, 6, 10, 15}
	// act
	actual := minimumAbsDifference(arr)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_3(t *testing.T) {
	// arrange
	expect := [][]int{{-14, -10}, {19, 23}, {23, 27}}
	arr := []int{3, 8, -10, 23, 19, -4, -14, 27}
	// act
	actual := minimumAbsDifference(arr)
	// assert
	assert.Equal(t, expect, actual)
}
