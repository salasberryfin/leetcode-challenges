/*
 * Hard to come up with the correct "algorithm" but easy to implement
 * Not a particularly useful problem
 *
 * Keys:
 *  - 0s will always be on the right side
 *  - 1s will always be on the left side
 *  - A string full of 0s or 1s is also valid
 *
 * With these keys:
 *  - Iterate until a '1' is found
 *  - When '1', there are two alternatives:
 *      a. Flip for a '0'
 *      b. Keep the '1'
 *  - Keep a counter of the number of ones in the string
 *  - When a '0' is found, calculate the minimum of:
 *      a. Flipping 0s to 1s: min_flips + 1
 *      b. Flipping 1s to 0s: count_ones (so far)
 *  - With this approach we're testing all possibilities.
 */
use std::cmp;

struct Solution;

impl Solution {
    pub fn min_flips_mono_incr(s: String) -> i32 {
        let mut min_flips: i32 = 0;
        let mut ones: i32 = 0;
        for c in s.chars() {
            if c == '1' {
                ones += 1;
            } else {
                min_flips = cmp::min(min_flips + 1, ones);
            }
        }

        return min_flips;
    }
}

#[test]
fn test_case_1() {
    let s = "00110";
    let output = 1;
    let result = Solution::min_flips_mono_incr(s.to_string());
    assert_eq!(output, result);
}

#[test]
fn test_case_2() {
    let s = "010110";
    let output = 2;
    let result = Solution::min_flips_mono_incr(s.to_string());
    assert_eq!(output, result);
}

#[test]
fn test_case_3() {
    let s = "00011000";
    let output = 2;
    let result = Solution::min_flips_mono_incr(s.to_string());
    assert_eq!(output, result);
}

#[test]
fn test_case_4() {
    let s = "0101100011";
    let output = 3;
    let result = Solution::min_flips_mono_incr(s.to_string());
    assert_eq!(output, result);
}
