/*
 * There's the intuitive solution of iterating over every possibility and brute
 * force the maximum sum, which is very slow for large inputs O(n^3)
 *
 * The optimal solution is to use Kadane's algorithm, which allows to calculate
 * the largest sum of subarray within an array. The original algorithm is used for
 * non cyclic arrays (see problem 53) so we have to adapt its result to work for
 * circular arrays.
 *
 * Keys:
 *  - Using Kadane's algorithm we obtain the maximum sum by:
 *      - Keeping track of the global max sum and the max sum at the current element
 *      - Iterating over all elements i and updating values:
 *          current_max = max(i, current_max + i)
 *          global_max = max(global_max, current_max)
 *      - We are only tracking the non-circular array solution
 *  - If we also get the minimum sum in the array, using the same Kadane's algorithm.
 *      - result is max(total_sum(array)-global_min, global_max)
 *      - still need to check for arrays with only negative numbers:
 *          - in that case: if global_max < 0 -> return global_max
 */
use std::cmp;

struct Solution;

impl Solution {
    pub fn sum_digits(nums: Vec<i32>) -> i32 {
        let mut sum: i32 = 0;

        for i in nums {
            sum += i;
        }

        return sum;
    }

    pub fn brute_force_max_subarray_sum_circular(nums: Vec<i32>) -> i32 {
        let mut double_nums: Vec<i32> = Vec::new();
        let mut max_sum: i32 = i32::MIN;
        let n: usize = nums.len();

        double_nums.extend(nums.iter().cloned());
        double_nums.extend(nums.iter().cloned());

        for i in 0..n {
            for j in i..i + n {
                let sum = Solution::sum_digits(double_nums[i..j + 1].to_vec());
                if sum > max_sum {
                    max_sum = sum;
                }
            }
        }

        return max_sum;
    }

    pub fn max_subarray_sum_circular(nums: Vec<i32>) -> i32 {
        let n: usize = nums.len();

        let mut max_sum: i32 = i32::MIN;
        let mut current_max: i32 = 0;
        let mut min_sum: i32 = i32::MAX;
        let mut current_min: i32 = 100000;
        let mut total_sum: i32 = 0;
        for i in nums.iter() {
            current_max = cmp::max(*i, current_max + *i);
            max_sum = cmp::max(max_sum, current_max);
            current_min = cmp::min(*i, current_min + *i);
            min_sum = cmp::min(min_sum, current_min);
            total_sum += i;
        }
        //println!("Total sum {}", total_sum);
        //println!("Max sum {}", max_sum);
        //println!("Min sum {}", min_sum);
        //println!("Result is max({}, {})", total_sum - min_sum, max_sum);

        if max_sum < 0 {
            return max_sum;
        }

        return cmp::max(total_sum - min_sum, max_sum);
    }
}

#[test]
fn test_case_1() {
    let nums = vec![1, -2, 3, -2];
    let output = 3;
    assert_eq!(Solution::max_subarray_sum_circular(nums), output);
    //assert_eq!(
    //    Solution::brute_force_max_subarray_sum_circular(nums),
    //    output
    //);
}

#[test]
fn test_case_2() {
    let nums = vec![5, -3, 5];
    let output = 10;
    assert_eq!(Solution::max_subarray_sum_circular(nums), output);
    //assert_eq!(
    //    Solution::brute_force_max_subarray_sum_circular(nums),
    //    output
    //);
}

#[test]
fn test_case_3() {
    let nums = vec![-3, -2, -3];
    let output = -2;
    assert_eq!(Solution::max_subarray_sum_circular(nums), output);
    //assert_eq!(
    //    Solution::brute_force_max_subarray_sum_circular(nums),
    //    output
    //);
}

#[test]
fn test_case_4() {
    let nums = vec![3, -1, 2, -1];
    let output = 4;
    assert_eq!(Solution::max_subarray_sum_circular(nums), output);
    //assert_eq!(
    //    Solution::brute_force_max_subarray_sum_circular(nums),
    //    output
    //);
}

#[test]
fn test_case_5() {
    let nums = vec![-5, 4, -6];
    let output = 4;
    assert_eq!(Solution::max_subarray_sum_circular(nums), output);
}
