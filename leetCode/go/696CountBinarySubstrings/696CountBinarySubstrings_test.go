package countbinarysubstrings

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	s := "00110011"
	expect := 6
	// act
	actual := countBinarySubstrings(s)
	// assert
	assert.True(t, reflect.DeepEqual(expect, actual))
	assert.Equal(t, expect, actual)
}

func Test_2(t *testing.T) {
	// arrange
	s := "10101"
	expect := 4
	// act
	actual := countBinarySubstrings(s)
	// assert
	assert.True(t, reflect.DeepEqual(expect, actual))
	assert.Equal(t, expect, actual)
}

func Test_3(t *testing.T) {
	// arrange
	s := "1"
	expect := 0
	// act
	actual := countBinarySubstrings(s)
	// assert
	assert.True(t, reflect.DeepEqual(expect, actual))
	assert.Equal(t, expect, actual)
}
