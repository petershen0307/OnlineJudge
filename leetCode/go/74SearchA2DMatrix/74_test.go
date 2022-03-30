package searcha2dmatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	expect := true
	target := 3
	input := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	// act
	actual := searchMatrix(input, target)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	expect := false
	target := 13
	input := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	// act
	actual := searchMatrix(input, target)
	// assert
	assert.Equal(t, expect, actual)
}
