pub struct Solution;
impl Solution {
    pub fn get_concatenation(nums: Vec<i32>) -> Vec<i32> {
        let mut r: Vec<i32> = vec![0; nums.len() * 2];
        for (index, num) in nums.iter().enumerate() {
            r[index] = *num;
            r[index + nums.len()] = *num;
        }
        return r;
    }
}

#[test]
fn test_1() {
    let nums: Vec<i32> = vec![1, 2, 3];
    let r = Solution::get_concatenation(nums);
    assert_eq!(r, vec![1, 2, 3, 1, 2, 3]);
}
