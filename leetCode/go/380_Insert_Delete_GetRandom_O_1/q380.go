package q380

// reference solution : https://leetcode.com/problems/insert-delete-getrandom-o1/discuss/85422/AC-C%2B%2B-Solution.-Unordered_map-%2B-Vector
import "math/rand"

type RandomizedSet struct {
	vector []int       // for random get, value is the inserted value
	m      map[int]int // key is the inserted value, value is the vector's index
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m:      make(map[int]int),
		vector: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.vector = append(this.vector, val)
	this.m[val] = len(this.vector) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}
	lastVal := this.vector[len(this.vector)-1]
	this.vector[this.m[val]] = lastVal
	this.m[lastVal] = this.m[val]

	this.vector = this.vector[:len(this.vector)-1]
	delete(this.m, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.vector[rand.Int()%len(this.vector)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
