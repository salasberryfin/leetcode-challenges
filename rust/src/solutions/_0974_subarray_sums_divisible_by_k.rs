/*
* This is just a math problem.
* Impossible that I'd come up with a solution -other than the brute force option-
* to this problem
*
* Keys:
*  - Prefix Sum -> The sum between elements i and j: sum(i, j) can be obtained by:
*      sum(i, j) = sum(j) - sum(i-1)
*  - We know that a valid solution has to satisfy sum(i, j) % k = 0, so:
*      (sum(j) - sum(i-1)) % k = 0
*      -> sum(j) % k = sum(i-1) % k
*  - To simplify the iteration and make the solutio O(n):
*      - Keep the modulos of each element in the vector in a vector of mods (size k)
*      - For repetitions of each value of mod, we increase the result
*      - To make it work for negative numbers, apply the double modulo:
*          prefix = (((prefix + n) % k) + k) % k;
*/
struct Solution;

impl Solution {
    pub fn subarrays_div_by_k(nums: Vec<i32>, k: i32) -> i32 {
        let mut result: i32 = 0;

        let mut mods: Vec<i32> = vec![0; k as usize];
        mods[0] = 1;
        let mut prefix: i32 = 0;
        for n in nums.iter() {
            prefix = (((prefix + n) % k) + k) % k;
            result += mods[prefix as usize];
            mods[prefix as usize] += 1;
        }

        return result;
    }
}

#[test]
fn test_case_1() {
    let nums = vec![4, 5, 0, -2, -3, 1];
    let k = 5;
    let output = 7;
    assert_eq!(output, Solution::subarrays_div_by_k(nums, k))
}

#[test]
fn test_case_2() {
    let nums = vec![5];
    let k = 9;
    let output = 0;
    assert_eq!(output, Solution::subarrays_div_by_k(nums, k))
}
