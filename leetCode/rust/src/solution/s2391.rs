pub struct Solution;
impl Solution {
    pub fn garbage_collection(garbage: Vec<String>, travel: Vec<i32>) -> i32 {
        let mut m_count: i32 = 0;
        let mut m_travel: i32 = 0;
        let mut g_count: i32 = 0;
        let mut g_travel: i32 = 0;
        let mut p_count: i32 = 0;
        let mut p_travel: i32 = 0;
        let mut travel_count: i32 = 0;
        for (index, house_garbage) in garbage.iter().enumerate() {
            if index > 0 {
                travel_count += travel[index - 1];
            }
            for unit in house_garbage.chars() {
                match unit {
                    'G' => {
                        g_count += 1;
                        g_travel = travel_count;
                    }
                    'P' => {
                        p_count += 1;
                        p_travel = travel_count;
                    }
                    'M' => {
                        m_count += 1;
                        m_travel = travel_count;
                    }
                    _ => {}
                }
            }
        }
        return g_count + g_travel + p_count + p_travel + m_count + m_travel;
    }
}

#[test]
fn test_1() {
    // arrange
    let garbage: Vec<String> = vec![
        String::from("G"),
        String::from("P"),
        String::from("GP"),
        String::from("GG"),
    ];
    let travel: Vec<i32> = vec![2, 4, 3];
    let expected = 21;
    // act
    let output = Solution::garbage_collection(garbage, travel);
    // assert
    assert_eq!(expected, output);
}

#[test]
fn test_2() {
    // arrange
    let garbage: Vec<String> = vec![String::from("MMM"), String::from("PGM"), String::from("GP")];
    let travel: Vec<i32> = vec![3, 10];
    let expected = 37;
    // act
    let output = Solution::garbage_collection(garbage, travel);
    // assert
    assert_eq!(expected, output);
}
