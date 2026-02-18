package solution

import "container/heap"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	/*
		把 equations 當作點與邊
		values 看為邊的權重
		queries 就是找到從 點a 出發找到 點b 的路徑
		equations 不會出現 (a,b) (b,c) (a,d) (d,c) 這種 (a,c) 有兩種到達可能的情況
	*/
	edges := map[string]map[string]float64{}
	for i, n := range equations {
		s, e := n[0], n[1]
		if _, ok := edges[s]; !ok {
			edges[s] = map[string]float64{}
		}
		// 把反向的方程式當作邊一樣加入
		edges[s][e] = values[i]
		if _, ok := edges[e]; !ok {
			edges[e] = map[string]float64{}
		}
		edges[e][s] = 1 / values[i]

	}
	results := []float64{}
	for _, q := range queries {
		results = append(results, dijkstra(q[0], q[1], edges))
	}
	return results
}

type node struct {
	name            string
	weightFromStart float64
	isVisited       bool
}

type minHeap []node

func (n minHeap) Len() int           { return len(n) }
func (n minHeap) Less(i, j int) bool { return n[i].weightFromStart < n[j].weightFromStart }
func (n minHeap) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n *minHeap) Push(x any) {
	*n = append(*n, x.(node))
}
func (n *minHeap) Pop() any {
	t := (*n)[len(*n)-1]
	*n = (*n)[0 : len(*n)-1]
	return t
}

// map[string]map[string]float64
// (a,b) node_a, node_b, weight
func dijkstra(start, end string, edges map[string]map[string]float64) float64 {
	_, sok := edges[start]
	_, eok := edges[end]
	if !sok || !eok {
		return -1
	}
	nodeRecords := map[string]node{
		start: {
			name:            start,
			weightFromStart: 1.0,
		},
	}
	h := minHeap{
		{
			name:            start,
			weightFromStart: 1.0,
		},
	}
	heap.Init(&h)
	for h.Len() != 0 {
		n := heap.Pop(&h).(node)
		// 因為是 min heap 所以拿出來的是路徑最短的點, 若該點已經造訪過, 則表示當前的權重已經不是最小, 略過這次執行, 沒有造訪過才執行
		if v, ok := nodeRecords[n.name]; ok && v.isVisited {
			continue
		}
		nodeRecords[n.name] = node{
			name:            n.name,
			isVisited:       true,
			weightFromStart: nodeRecords[n.name].weightFromStart,
		}
		// 尋找連接的邊
		for e, w := range edges[n.name] {
			eNode := node{
				name:            e,
				weightFromStart: n.weightFromStart * w,
			}
			// 判斷 e 點是否需要更新權重
			if _, ok := nodeRecords[e]; ok && nodeRecords[e].weightFromStart > (n.weightFromStart*w) {
				nodeRecords[e] = node{
					name:            e,
					weightFromStart: n.weightFromStart * w,
					isVisited:       nodeRecords[e].isVisited,
				}
			} else if !ok {
				// 無 e 點紀錄，更新權重
				nodeRecords[e] = eNode
			}
			heap.Push(&h, eNode)
		}
	}
	if v, ok := nodeRecords[end]; ok && v.isVisited {
		return v.weightFromStart
	}
	return -1.0
}
