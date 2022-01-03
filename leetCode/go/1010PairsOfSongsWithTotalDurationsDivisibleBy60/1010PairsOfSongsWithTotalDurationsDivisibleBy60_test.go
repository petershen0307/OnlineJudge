package pairsofsongswithtotaldurationsdivisibleby60

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	expect := 3
	input := []int{30, 20, 150, 100, 40}
	// act
	actual := numPairsDivisibleBy60(input)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	expect := 3
	input := []int{60, 60, 60}
	// act
	actual := numPairsDivisibleBy60(input)
	// assert
	assert.Equal(t, expect, actual)
}
