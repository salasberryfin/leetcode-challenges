struct Solution;

impl Solution {
    pub fn sort(left: Vec<i32>, right: Vec<i32>) -> Vec<i32> {
        let mut i = 0;
        let mut j = 0;
        let mut result: Vec<i32> = Vec::new();

        while i < left.len() && j < right.len() {
            if right[j] < left[i] {
                result.push(right[j]);
                j += 1;
            } else {
                result.push(left[i]);
                i += 1;
            }
        }

        while i < left.len() {
            result.push(left[i]);
            i += 1;
        }
        while j < right.len() {
            result.push(right[j]);
            j += 1;
        }

        return result;
    }

    pub fn merge_sort(vector: Vec<i32>) -> Vec<i32> {
        if vector.len() <= 1 {
            return vector;
        }

        let mid = vector.len() / 2;
        let left = Self::merge_sort(vector[0..mid].to_vec());
        let right = Self::merge_sort(vector[mid..].to_vec());

        return Self::sort(left, right);
    }

    pub fn num_rescue_boats(people: Vec<i32>, limit: i32) -> i32 {
        let sorted = Self::merge_sort(people);
        println!("The sorted array is {:?}", sorted);

        let mut l = 0;
        let mut r = sorted.len() - 1;
        let mut result = 0;
        while l <= r {
            result += 1;
            if sorted[l] + sorted[r] <= limit {
                l += 1;
            }
            if r == 0 {
                break;
            }
            r -= 1;
        }

        result
    }
}

#[test]
fn test_case_1() {
    let people: Vec<i32> = vec![1, 2];
    let limit: i32 = 3;
    let output: i32 = 1;
    let result = Solution::num_rescue_boats(people, limit);

    assert_eq!(output, result);
}

#[test]
fn test_case_2() {
    let people: Vec<i32> = vec![3, 2, 2, 1];
    let limit: i32 = 3;
    let output: i32 = 3;
    let result = Solution::num_rescue_boats(people, limit);

    assert_eq!(output, result);
}

#[test]
fn test_case_3() {
    let people: Vec<i32> = vec![3, 5, 3, 4];
    let limit: i32 = 5;
    let output: i32 = 4;
    let result = Solution::num_rescue_boats(people, limit);

    assert_eq!(output, result);
}
