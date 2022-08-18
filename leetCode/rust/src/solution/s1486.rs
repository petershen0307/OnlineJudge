pub struct Solution;
impl Solution {
    pub fn xor_operation(n: i32, start: i32) -> i32 {
        let mut result = 0;
        for i in 0..n {
            result ^= start + i * 2;
        }
        return result;
    }
}

#[test]
fn test_1() {
    let n: i32 = 5;
    let start: i32 = 0;
    let r = Solution::xor_operation(n, start);
    assert_eq!(r, 8);
}
