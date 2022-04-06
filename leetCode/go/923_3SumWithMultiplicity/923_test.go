package sumwithmultiplicity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	expect := 20
	input := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}
	target := 8
	// act
	actual := threeSumMulti(input, target)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	expect := 12
	input := []int{1, 1, 2, 2, 2, 2}
	target := 5
	// act
	actual := threeSumMulti(input, target)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_3(t *testing.T) {
	// arrange
	expect := 7
	input := []int{1, 1, 2, 0, 2, 3, 0}
	target := 5
	// act
	actual := threeSumMulti(input, target)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_4(t *testing.T) {
	// arrange
	expect := 10
	input := []int{0, 0, 0, 0, 0}
	target := 0
	// act
	actual := threeSumMulti(input, target)
	// assert
	assert.Equal(t, expect, actual)
}
