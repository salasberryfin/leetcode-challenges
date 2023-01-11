/*
 * This was hell!
 *
 * A 'relatively' efficient solution is to check for palindromes using center characters
 *  - Each palindrome is mirrored from its center element
 *  - Iterate over each character and treat it as a center
 *      - For this center, expand the radius to contiguous
 *      - Keep expanding while the characters match on each side of the center
 *      - When we get to the first different item: we have the size of the palindrome and the left
 *      and right positions
 *  - Get the substring from those positions and return
*/
use std::convert::TryFrom;
use std::iter::FromIterator;

struct Solution;

impl Solution {
    pub fn expand(vec_s: Vec<char>, mut l: usize, mut r: usize) -> usize {
        let mut left: i32 = i32::try_from(l).unwrap();
        let mut right: i32 = i32::try_from(r).unwrap();
        let len_vec_s: i32 = i32::try_from(vec_s.len()).unwrap();
        while left >= 0
            && right < len_vec_s
            && vec_s[usize::try_from(left).unwrap()] == vec_s[usize::try_from(right).unwrap()]
        {
            left = left - 1;
            right = right + 1;
        }

        return usize::try_from(right - left - 1).unwrap();
    }

    pub fn longest_palindrome(s: String) -> String {
        let len_s = s.len();
        if len_s == 1 {
            return s;
        }

        let vec_s: Vec<char> = s.chars().collect();

        if len_s == 2 && vec_s[0] != vec_s[1] {
            return String::from(vec_s[0]);
        }

        let mut left = 0;
        let mut right = 0;
        for i in 0..len_s {
            let len_pal1 = Solution::expand(vec_s.clone(), i, i);
            let len_pal2 = Solution::expand(vec_s.clone(), i, i + 1);
            let len_pal = std::cmp::max(len_pal1, len_pal2);
            if len_pal > right - left {
                left = i - (len_pal - 1) / 2;
                right = i + len_pal / 2;
            }
        }

        return String::from_iter(&vec_s[left..=right]);
    }
}

#[test]
fn test1() {
    let s: String = "babad".to_string();
    let expected: String = "aba".to_string();
    let result = Solution::longest_palindrome(s);
    assert_eq!(result, expected);
}

#[test]
fn test2() {
    let s: String = "cbbd".to_string();
    let expected: String = "bb".to_string();
    let result = Solution::longest_palindrome(s);
    assert_eq!(result, expected);
}

#[test]
fn test3() {
    let s: String = "ac".to_string();
    let expected: String = "a".to_string();
    let result = Solution::longest_palindrome(s);
    assert_eq!(result, expected);
}

#[test]
fn test4() {
    let s: String = "abb".to_string();
    let expected: String = "bb".to_string();
    let result = Solution::longest_palindrome(s);
    assert_eq!(result, expected);
}

#[test]
fn test5() {
    let s: String = "ccd".to_string();
    let expected: String = "cc".to_string();
    let result = Solution::longest_palindrome(s);
    assert_eq!(result, expected);
}
