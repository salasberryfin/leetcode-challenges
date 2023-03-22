struct Solution;

impl Solution {
    pub fn is_palindrome(s: String) -> bool {
        let s = s
            .to_lowercase()
            .replace(|c: char| !c.is_alphanumeric(), "")
            .into_bytes();
        let mut left = 0;
        let mut right = s.len() - 1;

        if s.is_empty() {
            return true;
        }

        while left < right {
            if s[left] != s[right] {
                return false;
            }
            left += 1;
            right -= 1;
        }

        return true;
    }
}

#[test]
fn test_case_1() {
    let s = String::from("A man, a plan, a canal; Panama");
    let output = true;
    assert_eq!(Solution::is_palindrome(s), output);
}

#[test]
fn test_case_2() {
    let s = String::from("race a car");
    let output = false;
    assert_eq!(Solution::is_palindrome(s), output);
}
