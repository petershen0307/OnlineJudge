package minstack

type node struct {
	value int
	next  *node
}

type MinStack struct {
	head *node
}

func Constructor() MinStack {
	return MinStack{head: nil}
}

func (this *MinStack) Push(val int) {
	n := new(node)
	n.value = val
	n.next = this.head
	this.head = n
}

func (this *MinStack) Pop() {
	this.head = this.head.next
}

func (this *MinStack) Top() int {
	if this.head != nil {
		return this.head.value
	}
	return 0
}

func (this *MinStack) GetMin() int {
	var min *int
	for i := this.head; i != nil; i = i.next {
		if min == nil {
			min = new(int)
			*min = i.value
		}
		if *min > i.value {
			*min = i.value
		}
	}
	return *min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
