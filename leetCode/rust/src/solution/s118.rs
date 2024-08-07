struct Solution;
impl Solution {
    pub fn generate(num_rows: i32) -> Vec<Vec<i32>> {
        // bottom up
        let mut result = vec![[1].to_vec()];
        for _i in 2..num_rows + 1 {
            let last_row: &Vec<i32> = &result.last().unwrap();
            let mut new_row = vec![1];
            for i in 0..last_row.len() - 1 {
                new_row.push(last_row[i] + last_row[i + 1]);
            }
            new_row.push(1);
            result.push(new_row);
        }
        result
    }
}
