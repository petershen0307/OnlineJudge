package minstack

type node struct {
	value int
	next  *node
}

type MinStack struct {
	head *node
	min  *node
}

func Constructor() MinStack {
	return MinStack{head: nil}
}

func (this *MinStack) Push(val int) {
	n := new(node)
	n.value = val
	n.next = this.head
	this.head = n
	// track the current min value
	// input: 1,2,3,1,2,3 then pop
	// min stack: 1,1 we will store two 1 here
	// when pop min stack, the input stack head should be the same with min stack head
	if this.min == nil {
		this.min = new(node)
		this.min.value = val
	} else if this.GetMin() >= val {
		m := new(node)
		m.next = this.min
		m.value = val
		this.min = m
	}
}

func (this *MinStack) Pop() {
	if this.head != nil && this.head.value == this.GetMin() {
		this.min = this.min.next
	}
	this.head = this.head.next
}

func (this *MinStack) Top() int {
	if this.head != nil {
		return this.head.value
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if this.min != nil {
		return this.min.value
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
