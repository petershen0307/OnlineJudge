use std::ops::Index;

pub struct Solution;
impl Solution {
    pub fn restore_string(s: String, indices: Vec<i32>) -> String {
        let mut result: String = s.clone();
        let mut i: usize = 0;
        while i < indices.len() {
            result.replace_range(
                indices[i] as usize..indices[i] as usize + 1,
                s.chars().nth(i).unwrap().to_string().as_str(),
            );
            i += 1
        }
        return result;
    }
}

#[test]
fn test_1() {
    let s = String::from("codeleet");
    let indices = vec![4, 5, 6, 7, 0, 2, 1, 3];
    let output = Solution::restore_string(s, indices);
    let expected_output = String::from("leetcode");
    assert_eq!(expected_output, output)
}

#[test]
fn test_2() {
    let s = String::from("abc");
    let indices = vec![0, 1, 2];
    let output = Solution::restore_string(s, indices);
    let expected_output = String::from("abc");
    assert_eq!(expected_output, output)
}
