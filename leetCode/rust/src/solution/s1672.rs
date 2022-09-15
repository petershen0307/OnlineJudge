pub struct Solution;
impl Solution {
    pub fn maximum_wealth(accounts: Vec<Vec<i32>>) -> i32 {
        let mut max_wealth: i32 = 0;
        for customer_bank_wealth_list in accounts {
            let customer_wealth:i32 = customer_bank_wealth_list.iter().sum();
            if max_wealth < customer_wealth {
                max_wealth = customer_wealth;
            }
        }
        return  max_wealth;
    }
}

#[test]
fn test_1() {
    let accounts = vec![vec![1, 2, 3], vec![3, 2, 1]];
    let output = Solution::maximum_wealth(accounts);
    let expect_output = 6;
    assert_eq!(expect_output, output);
}

#[test]
fn test_2() {
    let accounts = vec![vec![1, 5], vec![7, 3], vec![3, 5]];
    let output = Solution::maximum_wealth(accounts);
    let expect_output = 10;
    assert_eq!(expect_output, output);
}

#[test]
fn test_3() {
    let accounts = vec![vec![2, 8, 7], vec![7, 1, 3], vec![1, 9, 5]];
    let output = Solution::maximum_wealth(accounts);
    let expect_output = 17;
    assert_eq!(expect_output, output);
}
