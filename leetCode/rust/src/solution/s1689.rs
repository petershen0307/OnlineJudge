pub struct Solution;
impl Solution {
    pub fn min_partitions(n: String) -> i32 {
        let mut max: u8 = 0;
        for b in n.as_bytes() {
            if max < *b {
                max = *b
            }
        }
        return (max - b'0') as i32;
    }
}

#[test]
fn test_1() {
    let n = String::from("32");
    let expected = 3;
    let ans = Solution::min_partitions(n);
    assert_eq!(ans, expected);
}

#[test]
fn test_2() {
    let n = String::from("82734");
    let expected = 8;
    let ans = Solution::min_partitions(n);
    assert_eq!(ans, expected);
}

#[test]
fn test_3() {
    let n = String::from("27346209830709182346");
    let expected = 9;
    let ans = Solution::min_partitions(n);
    assert_eq!(ans, expected);
}
