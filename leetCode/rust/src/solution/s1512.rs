pub struct Solution;
impl Solution {
    pub fn num_identical_pairs(nums: Vec<i32>) -> i32 {
        let mut good: i32 = 0;
        for (i, v) in nums.iter().enumerate() {
            for (j, v2) in nums.iter().enumerate() {
                if v == v2 && i < j {
                    good += 1;
                }
            }
        }
        return good;
    }
}

#[test]
fn test_1() {
    // arrange
    let nums = vec![1, 2, 3, 1, 1, 3];
    let expected = 4;
    // act
    let output = Solution::num_identical_pairs(nums);
    // assert
    assert_eq!(expected, output)
}

#[test]
fn test_2() {
    // arrange
    let nums = vec![1, 1, 1, 1];
    let expected = 6;
    // act
    let output = Solution::num_identical_pairs(nums);
    // assert
    assert_eq!(expected, output)
}

#[test]
fn test_3() {
    // arrange
    let nums = vec![1, 2, 3];
    let expected = 0;
    // act
    let output = Solution::num_identical_pairs(nums);
    // assert
    assert_eq!(expected, output)
}
