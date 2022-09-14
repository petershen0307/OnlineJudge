pub struct Solution;
impl Solution {
    pub fn final_value_after_operations(operations: Vec<String>) -> i32 {
        let mut x = 0;
        for operation in operations {
            match operation.as_str() {
                "++X" => x += 1,
                "X++" => x += 1,
                "--X" => x -= 1,
                "X--" => x -= 1,
                _ => {}
            }
        }
        return x;
    }
}

#[test]
fn test_1() {
    let operations = vec![
        String::from("--X"),
        String::from("X++"),
        String::from("X++"),
    ];
    let expected_output = 1;
    let output = Solution::final_value_after_operations(operations);
    assert_eq!(output, expected_output);
}

#[test]
fn test_2() {
    let operations = vec![
        String::from("++X"),
        String::from("X++"),
        String::from("X++"),
    ];
    let expected_output = 3;
    let output = Solution::final_value_after_operations(operations);
    assert_eq!(output, expected_output);
}

#[test]
fn test_3() {
    let operations = vec![
        String::from("++X"),
        String::from("X++"),
        String::from("X--"),
        String::from("--X"),
    ];
    let expected_output = 0;
    let output = Solution::final_value_after_operations(operations);
    assert_eq!(output, expected_output);
}
