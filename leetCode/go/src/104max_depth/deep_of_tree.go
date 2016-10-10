package main

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

var depth int
var maxDepthNum int

func goDeep(root *treeNode) {
	if nil != root {
		depth++
	} else {
		if depth > maxDepthNum {
			maxDepthNum = depth
		}
		return
	}
	goDeep(root.Left)
	goDeep(root.Right)
	depth--
}

func maxDepth(root *treeNode) int {
	goDeep(root)
	r := maxDepthNum
	maxDepthNum = 0
	depth = 0
	return r
}

func main() {

}
