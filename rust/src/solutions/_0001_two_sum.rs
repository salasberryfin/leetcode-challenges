use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut visited = HashMap::new();

        for (idx, element) in nums.iter().enumerate() {
            let current = target - element;
            if visited.contains_key(&current) {
                println!("found solution for {}, {}", visited[&current], idx);
                return vec![visited[&current], idx as i32];
            }
            visited.insert(element, idx as i32);
        }

        return vec![];
    }
}

#[test]
fn test1() {
    let nums: Vec<i32> = vec![2, 7, 11, 15];
    let target: i32 = 9;
    let expected: Vec<i32> = vec![0, 1];
    let result = Solution::two_sum(nums,target);
    assert_eq!(result, expected);
}

#[test]
fn test2() {
    let nums: Vec<i32> = vec![3, 2, 4];
    let target: i32 = 6;
    let expected: Vec<i32> = vec![1, 2];
    let result = Solution::two_sum(nums,target);
    assert_eq!(result, expected);
}
