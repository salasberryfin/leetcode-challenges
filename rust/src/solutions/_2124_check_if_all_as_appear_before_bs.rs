struct Solution {}

impl Solution {
    pub fn check_string(s: String) -> bool {
        let mut found_b = false;
        for c in s.chars() {
            if c == 'b' {
                found_b = true;
            } else if found_b {
                return false;
            }
        }
        true
    }
}

#[test]
fn test1() {
    let s: String = String::from("aaabbb");
    let expected: bool = true;
    let result = Solution::check_string(s);
    assert_eq!(result, expected);
}

#[test]
fn test2() {
    let s: String = String::from("abab");
    let expected: bool = false;
    let result = Solution::check_string(s);
    assert_eq!(result, expected);
}

#[test]
fn test3() {
    let s: String = String::from("bbb");
    let expected: bool = true;
    let result = Solution::check_string(s);
    assert_eq!(result, expected);
}
