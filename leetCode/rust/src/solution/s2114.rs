pub struct Solution;

impl Solution {
    pub fn most_words_found(sentences: Vec<String>) -> i32 {
        let mut result: i32 = 0;
        for sentence in sentences {
            let mut count: i32 = 0;
            for char_byte in sentence.as_bytes() {
                if *char_byte == b' ' {
                    count += 1;
                }
            }
            if result < count {
                result = count;
            }
        }
        // only count space interval, should plus 1 for words
        return result+1;
    }
}

#[test]
fn test_1() {
    let sentences = vec![
        String::from("alice and bob love leetcode"),
        String::from("i think so too"),
        String::from("this is great thanks very much"),
    ];
    let expected_output: i32 = 6;
    let output = Solution::most_words_found(sentences);
    assert_eq!(expected_output, output)
}

#[test]
fn test_2() {
    let sentences = vec![
        String::from("please wait"),
        String::from("continue to fight"),
        String::from("continue to win"),
    ];
    let expected_output: i32 = 3;
    let output = Solution::most_words_found(sentences);
    assert_eq!(expected_output, output)
}
