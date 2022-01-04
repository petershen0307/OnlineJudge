package findthetownjudge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	expect := 2
	n := 2
	trust := [][]int{{1, 2}}
	// act
	actual := findJudge(n, trust)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	expect := -1
	n := 3
	trust := [][]int{{1, 3}, {2, 3}, {3, 1}}
	// act
	actual := findJudge(n, trust)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_3(t *testing.T) {
	// arrange
	expect := 3
	n := 3
	trust := [][]int{{1, 3}, {2, 3}}
	// act
	actual := findJudge(n, trust)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_4(t *testing.T) {
	// arrange
	expect := 1
	n := 1
	trust := [][]int{}
	// act
	actual := findJudge(n, trust)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_5(t *testing.T) {
	// arrange
	expect := -1
	n := 2
	trust := [][]int{}
	// act
	actual := findJudge(n, trust)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_6(t *testing.T) {
	// arrange
	expect := -1
	n := 4
	trust := [][]int{{1, 3}, {1, 4}, {2, 3}}
	// act
	actual := findJudge(n, trust)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_7(t *testing.T) {
	// arrange
	expect := 3
	n := 4
	trust := [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}
	// act
	actual := findJudge(n, trust)
	// assert
	assert.Equal(t, expect, actual)
}

func Test_8(t *testing.T) {
	// arrange
	expect := -1
	n := 3
	trust := [][]int{{1, 2}, {2, 3}}
	// act
	actual := findJudge(n, trust)
	// assert
	assert.Equal(t, expect, actual)
}
