struct Solution {}

impl Solution {
    pub fn length_of_lis(nums: Vec<i32>) -> i32 {
        let mut result = 0;
        let size = nums.len();
        let mut arr = vec![1; size];

        for (index, num) in nums.iter().rev().enumerate() {
            let reversed_index = size - index - 1;
            let mut best = 0;
            for (j, _) in arr.iter().enumerate().skip(reversed_index + 1) {
                if nums[j] > *num {
                    if arr[j] > best {
                        best = arr[j];
                    }
                }
            }
            arr[reversed_index] += best;
            if arr[reversed_index] > result {
                result = arr[reversed_index];
            }
        }

        return result;
    }
}

#[test]
fn test1() {
    let nums: Vec<i32> = vec![10, 9, 2, 5, 3, 7, 101, 18];
    let expected: i32 = 4;
    let result = Solution::length_of_lis(nums);
    assert_eq!(result, expected);
}

fn test2() {
    let nums: Vec<i32> = vec![0, 1, 0, 3, 2, 3];
    let expected: i32 = 4;
    let result = Solution::length_of_lis(nums);
    assert_eq!(result, expected);
}
