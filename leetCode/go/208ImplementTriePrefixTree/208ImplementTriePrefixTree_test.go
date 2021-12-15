package implementtrieprefixtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	wordDict := Constructor()
	// act
	wordDict.Insert("hello")
	// assert
	assert.True(t, wordDict.Search("hello"))
}

func Test_2(t *testing.T) {
	// arrange
	wordDict := Constructor()
	// act
	wordDict.Insert("bad")
	wordDict.Insert("dad")
	wordDict.Insert("mad")
	// assert
	assert.False(t, wordDict.Search("pad"))
	assert.True(t, wordDict.Search("bad"))
}

func Test_3(t *testing.T) {
	// arrange
	wordDict := Constructor()
	// act
	wordDict.Insert("bad")
	wordDict.Insert("dad")
	wordDict.Insert("mad")
	// assert
	assert.False(t, wordDict.Search("pad"))
	assert.True(t, wordDict.Search("bad"))
	assert.True(t, wordDict.StartsWith("ba"))
	assert.True(t, wordDict.StartsWith("m"))
}

func Test_4(t *testing.T) {
	// arrange
	wordDict := Constructor()
	// act
	wordDict.Insert("bad")
	wordDict.Insert("dad")
	wordDict.Insert("mad")
	// assert
	assert.False(t, wordDict.Search("b"))
}

func Test_5(t *testing.T) {
	// arrange
	wordDict := Constructor()
	// act
	wordDict.Insert("b")
	wordDict.Insert("bad")
	wordDict.Insert("dad")
	wordDict.Insert("mad")
	// assert
	assert.True(t, wordDict.Search("b"))
}

func Test_6(t *testing.T) {
	// arrange
	wordDict := Constructor()
	// act
	wordDict.Insert("b")
	wordDict.Insert("b")
	// assert
	assert.True(t, wordDict.Search("b"))
	assert.True(t, wordDict.StartsWith("b"))
	assert.False(t, wordDict.Search("bb"))
	assert.True(t, wordDict.Search("b"))
	assert.False(t, wordDict.Search("bb"))
}

func Test_7(t *testing.T) {
	// arrange
	wordDict := Constructor()
	// act
	wordDict.Insert("b")
	wordDict.Insert("bad")
	wordDict.Insert("badd")
	wordDict.Insert("dad")
	wordDict.Insert("mad")
	// assert
	assert.True(t, wordDict.Search("b"))
	assert.True(t, wordDict.Search("bad"))
	assert.True(t, wordDict.Search("badd"))
}
