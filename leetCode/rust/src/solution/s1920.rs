pub struct Solution;
impl Solution {
    pub fn build_array(nums: Vec<i32>) -> Vec<i32> {
        let mut ans: Vec<i32> = Vec::new();
        for (i, num) in nums.iter().enumerate() {
            ans.insert(i, nums[*num as usize]);
        }
        return ans;
    }
}

#[test]
fn test_1() {
    let nums = vec![0, 2, 1, 5, 3, 4];
    let ans = Solution::build_array(nums);
    let expected = vec![0, 1, 2, 4, 5, 3];
    let matching = ans.iter().zip(&expected).filter(|&(a, b)| a == b).count();
    assert_eq!(matching, expected.len());
}

#[test]
fn test_2() {
    let nums = vec![5, 0, 1, 2, 3, 4];
    let ans = Solution::build_array(nums);
    let expected = vec![4, 5, 0, 1, 2, 3];
    let matching = ans.iter().zip(&expected).filter(|&(a, b)| a == b).count();
    assert_eq!(matching, expected.len());
}
