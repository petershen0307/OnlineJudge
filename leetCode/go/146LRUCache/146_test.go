package lrucache

import "testing"

func Test_1(t *testing.T) {
	cache := Constructor(2)
	r := 0
	cache.Put(1, 1)
	cache.Put(2, 2)
	r = cache.Get(1)
	if r != 1 {
		t.Error("")
	}
	cache.Put(3, 3)
	r = cache.Get(2)
	if r != -1 {
		t.Error("")
	}
	cache.Put(4, 4)
	r = cache.Get(1)
	if r != -1 {
		t.Error("")
	}
	r = cache.Get(3)
	if r != 3 {
		t.Error("")
	}
	r = cache.Get(4)
	if r != 4 {
		t.Error("")
	}
}
