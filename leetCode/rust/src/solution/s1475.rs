pub struct Solution;

impl Solution {
    pub fn final_prices(prices: Vec<i32>) -> Vec<i32> {
        let mut result = prices.clone();
        let mut i = 0;
        while i < prices.len() {
            for price in prices[i + 1..].iter() {
                if result[i] >= *price {
                    result[i] = result[i] - price;
                    break;
                }
            }
            i += 1;
        }
        result
    }
}

#[test]
fn test_1() {
    let prices = vec![8, 4, 6, 2, 3];
    let expected = vec![4, 2, 4, 2, 3];
    let ans = Solution::final_prices(prices);
    assert_eq!(ans, expected);
}

#[test]
fn test_2() {
    let prices = vec![1, 2, 3, 4, 5];
    let expected = vec![1, 2, 3, 4, 5];
    let ans = Solution::final_prices(prices);
    assert_eq!(ans, expected);
}

#[test]
fn test_3() {
    let prices = vec![10, 1, 1, 6];
    let expected = vec![9, 0, 1, 6];
    let ans = Solution::final_prices(prices);
    assert_eq!(ans, expected);
}
