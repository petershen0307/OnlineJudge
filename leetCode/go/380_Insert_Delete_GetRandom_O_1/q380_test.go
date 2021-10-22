package q380

import "testing"

func Test_case(t *testing.T) {
	set := Constructor()
	if !set.Insert(0) {
		t.Error("inert failed")
	}
	if !set.Insert(1) {
		t.Error("inert failed")
	}
	if !set.Remove(0) {
		t.Error("remove failed")
	}
	if !set.Insert(2) {
		t.Error("inert failed")
	}
	if !set.Remove(1) {
		t.Error("remove failed")
	}
	if set.GetRandom() != 2 {
		t.Error("get failed")
	}
}
