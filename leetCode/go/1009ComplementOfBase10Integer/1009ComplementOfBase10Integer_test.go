package complementofbase10integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	expect := 2
	n := 5
	// act
	actual := bitwiseComplement(n)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	expect := 0
	n := 7
	// act
	actual := bitwiseComplement(n)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_3(t *testing.T) {
	// arrange
	expect := 5
	n := 10
	// act
	actual := bitwiseComplement(n)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_4(t *testing.T) {
	// arrange
	expect := 1
	n := 0
	// act
	actual := bitwiseComplement(n)
	// assert
	assert.Equal(t, expect, actual)
}
