/*
 * This is a very straightforward problem when using Kadene's algorithm
 *
 * Keys:
 *  - Iterate over each item in the vector
 *  - Keep track of the golbal maximum sum and the current maximum sum
 *      - The current maximum sum is updated with max(item, current maximum sum + item)
 *      - After updating the current: the global max is updated if the new current is larger
 *  - It's important to properly initialize the variables to support only negative inputs
 */
use std::cmp;

struct Solution;

impl Solution {
    pub fn max_sub_array(nums: Vec<i32>) -> i32 {
        let mut max_sum: i32 = i32::MIN;

        // kadene's algorithm
        let mut current_max: i32 = -1000000;
        for i in nums {
            current_max = cmp::max(i, current_max + i);
            max_sum = cmp::max(max_sum, current_max);
        }

        return max_sum;
    }
}

#[test]
fn test_case_1() {
    let nums = vec![-2, 1, -3, 4, -1, 2, 1, -5, 4];
    let output = 6;
    assert_eq!(output, Solution::max_sub_array(nums));
}

#[test]
fn test_case_2() {
    let nums = vec![1];
    let output = 1;
    assert_eq!(output, Solution::max_sub_array(nums));
}

#[test]
fn test_case_3() {
    let nums = vec![5, 4, -1, 7, 8];
    let output = 23;
    assert_eq!(output, Solution::max_sub_array(nums));
}
