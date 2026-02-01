struct Solution;
impl Solution {
    pub fn _min_cost(n: i32, edges: Vec<Vec<i32>>) -> i32 {
        /*
        this method only find the way from 0 to n-1 not the minimum path
        use `paths` to record every travel edge
        if there is no edge for current node
            1. try to reverse the edge
            2. still no edge, go back the previous node
        stop condition
            1. find n-1 node
            2. there is no edge for current node(which should be 0) and no previous node
            3. edges is empty
        */
        let mut edges = edges;
        let mut paths: Vec<(i32, i32, i32)> = Vec::new(); // start, end, weight
        let mut weight = 0;
        let mut current_node = 0;
        loop {
            match Self::_find_edge(current_node, &edges) {
                None => {
                    // stop condition - paths is empty and current_node is 0 and no edge for current_node
                    // stop condition - edges is empty
                    if current_node == 0 && paths.is_empty() || edges.is_empty() {
                        return -1;
                    }
                    // go back to previous node
                    current_node = match paths.pop() {
                        None => 0,
                        Some(v) => {
                            weight -= v.2;
                            v.0
                        }
                    }
                }
                Some(r) => {
                    let edge = edges.swap_remove(r.0);
                    if r.1 {
                        paths.push((edge[1], edge[0], 2 * edge[2]));
                    } else {
                        paths.push((edge[0], edge[1], edge[2]));
                    }
                    weight += paths.last().unwrap().2;
                    current_node = paths.last().unwrap().1;
                }
            }
            // stop condition - found n-1
            if current_node == n - 1 {
                break;
            }
        }
        weight
    }

    // return the edge index, edge is reversed
    fn _find_edge(n: i32, edges: &[Vec<i32>]) -> Option<(usize, bool)> {
        // find direct edge with first match
        for (index, edge) in edges.iter().enumerate() {
            if n == edge[0] {
                return Some((index, false));
            }
        }
        // find possible reverse edge
        for (index, edge) in edges.iter().enumerate() {
            if n == edge[1] {
                return Some((index, true));
            }
        }
        None
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
