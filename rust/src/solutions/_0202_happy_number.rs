/*
 * More of a math problem than a programming one
 *
 * Keys:
 *  - For 1 digit numbers, we know that 1 means happy number -> true
 *  - If while looping any number is repeated -> false: it'll come back again to
 *  this infinitely
 *      - Use hash map to store known numbers and check for cycles
 *  - Iterate over each digit in the number using the modulo operator
 *  - If sum == 1 -> true
 */
use std::collections::HashMap;

struct Solution;

impl Solution {
    pub fn is_happy(n: i32) -> bool {
        let mut visited: HashMap<i32, bool> = HashMap::new();

        let mut value: i32 = n;
        loop {
            if value == 1 {
                return true;
            }
            visited.insert(value, true);
            let mut sum: i32 = 0;
            println!("Result of {}", value);
            loop {
                sum += (value % 10) * (value % 10);
                value = value / 10;
                if value == 0 {
                    break;
                }
            }
            println!(" = {}", sum);
            value = sum;
            if visited.contains_key(&value) {
                return false;
            } else if value == 1 {
                return true;
            }
        }
    }
}

#[test]
fn test_case_1() {
    let n = 19;
    let output = true;
    let result = Solution::is_happy(n);
    assert_eq!(output, result);
}

#[test]
fn test_case_2() {
    let n = 2;
    let output = false;
    let result = Solution::is_happy(n);
    assert_eq!(output, result);
}
