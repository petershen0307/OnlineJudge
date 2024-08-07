use std::collections::HashMap;

struct Solution;

impl Solution {
    pub fn fib(n: i32) -> i32 {
        let mut cache = [0; 2];
        cache[1] = 1;
        if n < 2 {
            return cache[n as usize];
        }
        for i in 2..n + 1 {
            cache[(i % 2) as usize] = cache[0] + cache[1];
        }
        cache[(n % 2) as usize]
    }
}
