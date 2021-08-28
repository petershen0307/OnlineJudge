package lrucache

type LRUCache struct {
	capacity int
	// hash table
	data map[int]*struct {
		value  int
		weight int
	}
	weightCounter int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:      capacity,
		weightCounter: 0,
		data: make(map[int]*struct {
			value  int
			weight int
		}),
	}
}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.data[key]; ok {
		this.data[key].weight = this.weightCounter
		this.weightCounter++
		return value.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.data[key]; ok {
		this.data[key].value = value
		this.data[key].weight = this.weightCounter
		this.weightCounter++
		return
	}
	if len(this.data) == this.capacity {
		// remove oldest value
		minWeight := int(^uint(0) >> 1)
		removedKey := -1
		for k, v := range this.data {
			if v.weight < minWeight {
				minWeight = v.weight
				removedKey = k
			}
		}
		delete(this.data, removedKey)
	}
	this.data[key] = &struct {
		value  int
		weight int
	}{value: value, weight: this.weightCounter}
	this.weightCounter++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
