pub struct Solution;

impl Solution {
    pub fn convert_temperature(celsius: f64) -> Vec<f64> {
        let kelvin: f64 = celsius + 273.15;
        let fahrenheit: f64 = ((celsius * 1.8 + 32.0) * 10_f64.powf(5.0)).round() / 10_f64.powf(5.0);
        vec![kelvin, fahrenheit]
    }
}

#[test]
fn test_1() {
    // arrange
    let celsius = 0.0;
    let kelvin = 273.15;
    let fahrenheit = 32.0;
    // act
    let result = Solution::convert_temperature(celsius);
    // assert
    assert_eq!(result.len(), 2);
    assert_eq!(result[0], kelvin);
    assert_eq!(result[1], fahrenheit);
}

#[test]
fn test_2() {
    // arrange
    let celsius = 19.99999;
    let kelvin = 293.14999;
    let fahrenheit = 67.99998;
    // act
    let result = Solution::convert_temperature(celsius);
    // assert
    assert_eq!(result.len(), 2);
    assert_eq!(result[0], kelvin);
    assert_eq!(result[1], fahrenheit);
}
