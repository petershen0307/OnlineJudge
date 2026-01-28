struct Solution;
impl Solution {
    pub fn minimum_abs_difference(arr: Vec<i32>) -> Vec<Vec<i32>> {
        let mut arr = arr;
        arr.sort();
        let mut diffs = Vec::new();
        let mut i = 0;
        // calculate the diffs
        while i + 1 < arr.len() {
            diffs.push(arr[i + 1] - arr[i]);
            i += 1;
        }
        // find the minimum diff
        let mut min_diff_value = diffs.first().unwrap(); // index, value
        let mut result = Vec::new();
        for v in diffs.iter() {
            if v < min_diff_value {
                min_diff_value = v
            }
        }
        // convert to arr index value
        for (index, v) in diffs.iter().enumerate() {
            if min_diff_value == v {
                result.push(vec![arr[index], arr[index + 1]]);
            }
        }
        result
    }
}

#[test]
fn test_1() {
    let n = vec![4, 2, 1, 3];
    let expected = vec![vec![1, 2], vec![2, 3], vec![3, 4]];
    let ans = Solution::minimum_abs_difference(n);
    assert_eq!(ans, expected);
}

#[test]
fn test_2() {
    let n = vec![1, 3, 6, 10, 15];
    let expected = vec![vec![1, 3]];
    let ans = Solution::minimum_abs_difference(n);
    assert_eq!(ans, expected);
}

#[test]
fn test_3() {
    let n = vec![3, 8, -10, 23, 19, -4, -14, 27];
    let expected = vec![vec![-14, -10], vec![19, 23], vec![23, 27]];
    let ans = Solution::minimum_abs_difference(n);
    assert_eq!(ans, expected);
}
