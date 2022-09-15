pub struct Solution;
impl Solution {
    pub fn shuffle(nums: Vec<i32>, n: i32) -> Vec<i32> {
        let mut result: Vec<i32> = Vec::new();
        for i in 0..n {
            result.push(nums[i as usize]);
            result.push(nums[n as usize + i as usize]);
        }
        return result;
    }
}

#[test]
fn test_1() {
    let nums = vec![2, 5, 1, 3, 4, 7];
    let n = 3;
    let output = Solution::shuffle(nums, n);
    let expected_output = vec![2, 3, 5, 4, 1, 7];
    assert_eq!(expected_output, output)
}

#[test]
fn test_2() {
    let nums = vec![1, 2, 3, 4, 4, 3, 2, 1];
    let n = 4;
    let output = Solution::shuffle(nums, n);
    let expected_output = vec![1, 4, 2, 3, 3, 2, 4, 1];
    assert_eq!(expected_output, output)
}

#[test]
fn test_3() {
    let nums = vec![1, 1, 2, 2];
    let n = 2;
    let output = Solution::shuffle(nums, n);
    let expected_output = vec![1, 2, 1, 2];
    assert_eq!(expected_output, output)
}
