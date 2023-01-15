/*
 * This is not the most efficient way, but seems sensible and intuitive
 * Keys:
 *  - Store the count for each character in both strings: HashMap
 *  - Find the bulls first: equivalente characters
 *  - Subtract from the count maps each time a character is visited
 *  - Update the count with the remaining matches, the cows
 */
use std::collections::HashMap;

struct Solution;

impl Solution {
    pub fn get_hint(secret: String, guess: String) -> String {
        let result: String;
        let mut bulls = 0;
        let mut cows = 0;

        let mut secret_map: HashMap<char, i32> = HashMap::new();
        for (_, n) in secret.chars().enumerate() {
            *secret_map.entry(n).or_default() += 1;
        }
        let mut guess_map: HashMap<char, i32> = HashMap::new();
        for (_, n) in guess.chars().enumerate() {
            *guess_map.entry(n).or_default() += 1;
        }
        for (i, c) in guess.chars().enumerate() {
            if secret_map.contains_key(&c) && secret_map[&c] > 0 {
                if c == secret.chars().nth(i).unwrap() {
                    *secret_map.entry(c).or_default() -= 1;
                    *guess_map.entry(c).or_default() -= 1;
                    bulls = bulls + 1;
                }
            }
        }
        for (k, v) in secret_map {
            if guess_map.contains_key(&k) {
                if v == guess_map[&k] || v < guess_map[&k] {
                    cows += v;
                } else if v > guess_map[&k] {
                    cows += guess_map[&k];
                }
            }
        }

        let str_bulls = bulls.to_string() + "A";
        let str_cows = cows.to_string() + "B";
        result = str_bulls.to_string() + &str_cows;

        return result;
    }
}

#[test]
fn test_case_1() {
    let secret = "1807";
    let guess = "7810";
    let output = "1A3B";
    let result = Solution::get_hint(secret.to_string(), guess.to_string());
    assert_eq!(output, result);
}

#[test]
fn test_case_2() {
    let secret = "1123";
    let guess = "0111";
    let output = "1A1B";
    let result = Solution::get_hint(secret.to_string(), guess.to_string());
    assert_eq!(output, result);
}

#[test]
fn test_case_3() {
    let secret = "11";
    let guess = "11";
    let output = "2A0B";
    let result = Solution::get_hint(secret.to_string(), guess.to_string());
    assert_eq!(output, result);
}

#[test]
fn test_case_4() {
    let secret = "1122";
    let guess = "1222";
    let output = "3A0B";
    let result = Solution::get_hint(secret.to_string(), guess.to_string());
    assert_eq!(output, result);
}
