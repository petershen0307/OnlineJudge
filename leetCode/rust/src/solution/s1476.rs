/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * let obj = SubrectangleQueries::new(rectangle);
 * obj.update_subrectangle(row1, col1, row2, col2, newValue);
 * let ret_2: i32 = obj.get_value(row, col);
 */
struct SubrectangleQueries {
    m_rectangle: Vec<Vec<i32>>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl SubrectangleQueries {
    fn new(rectangle: Vec<Vec<i32>>) -> Self {
        SubrectangleQueries {
            m_rectangle: rectangle,
        }
    }

    fn update_subrectangle(&mut self, row1: i32, col1: i32, row2: i32, col2: i32, new_value: i32) {
        for c in col1..col2 + 1 {
            for r in row1..row2 + 1 {
                self.m_rectangle[r as usize][c as usize] = new_value
            }
        }
    }

    fn get_value(&self, row: i32, col: i32) -> i32 {
        return self.m_rectangle[row as usize][col as usize];
    }
}

#[test]
fn test_1() {
    // arrange
    let rectangle = vec![vec![1, 2, 1], vec![4, 3, 4], vec![3, 2, 1], vec![1, 1, 1]];
    let mut sub_rec = SubrectangleQueries::new(rectangle);
    // act
    // assert
    let mut r = sub_rec.get_value(0, 2);
    assert_eq!(r, 1);
    sub_rec.update_subrectangle(0, 0, 3, 2, 5);
    r = sub_rec.get_value(0, 2);
    assert_eq!(r, 5);
}
