pub struct Solution;
impl Solution {
    pub fn running_sum(nums: Vec<i32>) -> Vec<i32> {
        let mut ans: Vec<i32> = Vec::new();
        for (i, num) in nums.iter().enumerate() {
            if i == 0 {
                ans.push(*num);
                continue;
            }
            ans.push(ans[i-1] + *num)
        }
        return ans;
    }
}

#[test]
fn test_1() {
    let nums = vec![1, 2, 3, 4];
    let expected = [1, 3, 6, 10];
    let ans = Solution::running_sum(nums);
    let matching = ans.iter().zip(&expected).filter(|&(a, b)| a == b).count();
    assert_eq!(matching, expected.len());
}

#[test]
fn test_2() {
    let nums = vec![1, 1, 1, 1, 1];
    let expected = [1, 2, 3, 4, 5];
    let ans = Solution::running_sum(nums);
    let matching = ans.iter().zip(&expected).filter(|&(a, b)| a == b).count();
    assert_eq!(matching, expected.len());
}

#[test]
fn test_3() {
    let nums = vec![3, 1, 2, 10, 1];
    let expected = [3, 4, 6, 16, 17];
    let ans = Solution::running_sum(nums);
    let matching = ans.iter().zip(&expected).filter(|&(a, b)| a == b).count();
    assert_eq!(matching, expected.len());
}
