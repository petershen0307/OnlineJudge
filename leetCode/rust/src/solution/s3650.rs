struct Solution;
impl Solution {
    pub fn min_cost(n: i32, edges: Vec<Vec<i32>>) -> i32 {
        // https://leetcode.com/problems/minimum-cost-path-with-edge-reversals/solutions/7527912/it-is-just-a-simple-dijkstra-by-balepavl-wsxl/?envType=daily-question&envId=2026-01-27
        // 根據提供的解法，我們可以直接將所有邊直接反轉存回 edges
        // 會不會遇到同一個點走到兩個反轉邊呢? 不會，因為每個邊都為正數，走成迴圈就不會是最小路徑
        // 利用 vec random access 加快查找速度
        let mut new_edges: Vec<Vec<(i32, i32)>> = vec![Vec::new(); n as usize]; // index: start_node, value (end_node, edge_weight)
        for edge in edges {
            // push original
            new_edges[edge[0] as usize].push((edge[1], edge[2]));
            // push reverse
            new_edges[edge[1] as usize].push((edge[0], edge[2].checked_mul(2).unwrap_or(i32::MAX)));
        }
        // go Dijkstra algorithm
        Self::dijkstra(0, n - 1, &new_edges)
    }

    // the edges structure definition index: start_node, value (end_node, edge_weight)
    fn dijkstra(start: i32, target: i32, edges: &[Vec<(i32, i32)>]) -> i32 {
        use std::collections::BinaryHeap;
        // 使用 min heap 來幫助每次取得最短路徑的 node
        // min heap 存的結構是(n, weight) 到達點(n) 所需要的 weight
        // node_records 紀錄點是否有走過與到該點的最短 weight
        // 1. 先將起始點放入 min heap (0, 0)
        // 2. loop 結束條件 當 min heap 為空時
        // 2.1 從 min heap 取出一個點(n)，標註點 n 已經走訪過
        // 2.2 把該點(n)連接的點(a,b)都取出
        // 2.3 若點a, b都沒有走訪過，則都放入 min heap，a.weight = n.weight+w(n,a)，若走訪過則跳過處理
        // w(n,a) => 從 n 走到 a 需要的 weight
        // n.weight => 從起始點走到 n 的最小 weight
        #[derive(Clone, Eq)]
        struct NodeWeight {
            node: i32,
            weight_from_start_node: i32,
        }
        impl Ord for NodeWeight {
            fn cmp(&self, other: &Self) -> std::cmp::Ordering {
                self.weight_from_start_node
                    .cmp(&(other.weight_from_start_node))
            }
        }
        impl PartialOrd for NodeWeight {
            fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
                Some(self.cmp(other))
            }
        }

        impl PartialEq for NodeWeight {
            fn eq(&self, other: &Self) -> bool {
                self.node == other.node
            }
        }
        #[derive(Clone)]
        struct NodeRecord {
            is_visited: bool,
            weight: i32,
        }
        let mut min_heap = BinaryHeap::new();
        min_heap.push(std::cmp::Reverse(NodeWeight {
            node: start,
            weight_from_start_node: 0,
        }));
        let mut node_records = vec![
            NodeRecord {
                is_visited: false,
                weight: i32::MAX,
            };
            (target + 1) as usize
        ];
        while !min_heap.is_empty() {
            let node = match min_heap.pop() {
                Some(n) => n.0,
                None => continue,
            };
            // skip the visited node here for more efficient
            if node_records[node.node as usize].is_visited {
                continue;
            }
            node_records[node.node as usize].is_visited = true;
            if node_records[node.node as usize].weight > node.weight_from_start_node {
                node_records[node.node as usize].weight = node.weight_from_start_node;
            }
            // find related node in edges
            for edge in edges[node.node as usize].clone() {
                if node_records[edge.0 as usize].is_visited {
                    // the node is visited, so skip
                    continue;
                }
                min_heap.push(std::cmp::Reverse(NodeWeight {
                    node: edge.0,
                    weight_from_start_node: node.weight_from_start_node + edge.1,
                }));
            }
        }
        if !node_records[target as usize].is_visited {
            return -1;
        }
        node_records[target as usize].weight
    }
}

#[test]
fn test_1() {
    let input_n = 4;
    let input_edges = vec![vec![0, 1, 3], vec![3, 1, 1], vec![2, 3, 4], vec![0, 2, 2]];
    let result = Solution::min_cost(input_n, input_edges);
    assert_eq!(5, result)
}

#[test]
fn test_2() {
    let input_n = 4;
    let input_edges = vec![vec![0, 2, 1], vec![2, 1, 1], vec![1, 3, 1], vec![2, 3, 3]];
    let result = Solution::min_cost(input_n, input_edges);
    assert_eq!(3, result)
}

#[test]
fn test_3() {
    let input_n = 3;
    let input_edges = vec![vec![2, 0, 21], vec![1, 0, 12], vec![1, 2, 5]];
    let result = Solution::min_cost(input_n, input_edges);
    assert_eq!(29, result)
}
