/*
 * Best way to understand the problem is to analyze the examples in a notebook/whiteboard
 * The key idea is:
 *  - The cost of each element is:
 *      cost to n = min([(cost to n-1) + (cost from n-1)], [(cost to n-2) + (cost from n-2)])
 *  - Since we can go one or two steps at a time
 * We start by initializing the min_costs Vector:
 *  - Set first and second element to 0: we can start from either of them
 *  - Iterate (starting from the 2-index) over the elements and apply the above formula
 *  - Add the result to the vector
 *  - After completing the loop, the last element of the min_costs vector is the minimum value to
 *  reach the top
*/
use std::cmp;

struct Solution;

impl Solution {
    pub fn min_cost_climbing_stairs(cost: Vec<i32>) -> i32 {
        let cost_len = cost.len();

        let mut min_costs: Vec<i32> = Vec::new();

        min_costs.push(0);
        min_costs.push(0);
        for n in 2..cost_len+1 {
            let cost1 = cost[n-1] + min_costs[n-1];
            let cost2 = cost[n-2] + min_costs[n-2];
            min_costs.push(cmp::min(cost1, cost2));
        }

        return *min_costs.last().unwrap();
    }
}

#[test]
fn test_example1() {
    let cost: Vec<i32> = vec![10, 15, 20];
    let expected = 15;
    let result = Solution::min_cost_climbing_stairs(cost);
    assert_eq!(expected, result);
}

#[test]
fn test_example2() {
    let cost: Vec<i32> = vec![1,100,1,1,1,100,1,1,100,1];
    let expected = 6;
    let result = Solution::min_cost_climbing_stairs(cost);
    assert_eq!(expected, result);
}
