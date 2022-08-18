pub struct Solution;
impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        let mut digit_array: [u8; 26] = [0; 26];
        for byte in s.as_bytes() {
            digit_array[(*byte) as usize % 26] += 1;
        }
        for byte in t.as_bytes() {
            digit_array[(*byte) as usize % 26] -= 1;
        }
        return digit_array == [0; 26];
    }
}

#[test]
fn test_1() {
    let s = String::from("anagram");
    let t = String::from("nagaram");
    let r = Solution::is_anagram(s, t);
    assert_eq!(r, true);
}
