pub struct Solution;
impl Solution {
    pub fn reverse_bits(x: u32) -> u32 {
        let mut result: u32 = 0;
        for i in 0..32 {
            if (x & (1 << i)) != 0 {
                result = result | (1 << 31 - i);
            }
        }
        return result;
    }
}

#[test]
fn test_1() {
    // arrange
    let x: u32 = 0b00000010100101000001111010011100u32;
    let expected = 0b00111001011110000010100101000000u32;
    // act
    let output = Solution::reverse_bits(x);
    // assert
    assert_eq!(expected, output);
}

#[test]
fn test_2() {
    // arrange
    let x: u32 = 1u32;
    let expected = 0b10000000000000000000000000000000u32;
    // act
    let output = Solution::reverse_bits(x);
    // assert
    assert_eq!(expected, output);
}

#[test]
fn test_3() {
    // arrange
    let x: u32 = 0b11111111111111111111111111111101u32;
    let expected = 0b10111111111111111111111111111111u32;
    // act
    let output = Solution::reverse_bits(x);
    // assert
    assert_eq!(expected, output);
}
