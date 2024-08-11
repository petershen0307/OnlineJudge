struct Solution;
impl Solution {
    pub fn word_break(s: String, word_dict: Vec<String>) -> bool {
        // bfs
        use std::iter;
        let mut stack = Vec::new();
        let mut visited: Vec<bool> = iter::repeat(false).take(s.len()).collect();
        // start from first character, if the first character can't match any word, then we don't need to check on word which start from 2,3...
        stack.push(0);
        while let Some(i) = stack.pop() {
            for word in word_dict.iter() {
                let j = i + word.len() - 1;
                if j >= s.len() {
                    // exceed s length
                    continue;
                }
                let sub = &s[i..=j];
                if (sub == word.as_str()) && (j == s.len() - 1) {
                    return true;
                }
                if j + 1 < s.len() && sub == word.as_str() && !visited[j + 1] {
                    // we find a word, so we put the next character index to the stack
                    stack.push(j + 1);
                    visited[j + 1] = true;
                }
            }
        }
        false
    }
}

#[test]
fn test_1() {
    assert!(Solution::word_break(
        "leetcode".to_string(),
        vec!["leet".to_string(), "code".to_string()]
    ));
    assert!(Solution::word_break(
        "applepenapple".to_string(),
        vec!["apple".to_string(), "pen".to_string()]
    ));
    assert!(!Solution::word_break(
        "catsandog".to_string(),
        vec![
            "cats".to_string(),
            "dog".to_string(),
            "sand".to_string(),
            "and".to_string(),
            "cat".to_string()
        ]
    ));
}

#[test]
fn test_2() {
    assert!(Solution::word_break(
        "cars".to_string(),
        vec!["car".to_string(), "ca".to_string(), "rs".to_string()]
    ));
}
