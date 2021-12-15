package designaddandsearchwordsdatastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	// arrange
	wordDict := Constructor()
	// act
	wordDict.AddWord("hello")
	// assert
	assert.True(t, wordDict.Search("hello"))
}

func Test_2(t *testing.T) {
	// arrange
	wordDict := Constructor()
	// act
	wordDict.AddWord("bad")
	wordDict.AddWord("dad")
	wordDict.AddWord("mad")
	// assert
	assert.False(t, wordDict.Search("pad"))
	assert.True(t, wordDict.Search("bad"))
}

func Test_3(t *testing.T) {
	// arrange
	wordDict := Constructor()
	// act
	wordDict.AddWord("bad")
	wordDict.AddWord("dad")
	wordDict.AddWord("mad")
	// assert
	assert.False(t, wordDict.Search("pad"))
	assert.True(t, wordDict.Search("bad"))
	assert.True(t, wordDict.Search(".ad"))
	assert.True(t, wordDict.Search("b.."))
}

func Test_4(t *testing.T) {
	// arrange
	wordDict := Constructor()
	// act
	wordDict.AddWord("bad")
	wordDict.AddWord("dad")
	wordDict.AddWord("mad")
	// assert
	assert.False(t, wordDict.Search("b"))
}

func Test_5(t *testing.T) {
	// arrange
	wordDict := Constructor()
	// act
	wordDict.AddWord("b")
	wordDict.AddWord("bad")
	wordDict.AddWord("dad")
	wordDict.AddWord("mad")
	// assert
	assert.True(t, wordDict.Search("b"))
}

func Test_6(t *testing.T) {
	// arrange
	wordDict := Constructor()
	// act
	wordDict.AddWord("b")
	wordDict.AddWord("b")
	// assert
	assert.True(t, wordDict.Search("b"))
	assert.True(t, wordDict.Search("."))
	assert.False(t, wordDict.Search("bb"))
	assert.True(t, wordDict.Search("b"))
	assert.False(t, wordDict.Search(".b"))
	assert.False(t, wordDict.Search("b."))
}

func Test_7(t *testing.T) {
	// arrange
	wordDict := Constructor()
	// act
	wordDict.AddWord("b")
	wordDict.AddWord("bad")
	wordDict.AddWord("badd")
	wordDict.AddWord("dad")
	wordDict.AddWord("mad")
	// assert
	assert.True(t, wordDict.Search("b"))
	assert.True(t, wordDict.Search("bad"))
	assert.True(t, wordDict.Search("badd"))
}
