/*
 * This is an extremely difficult problem. Reference here: https://www.youtube.com/watch?v=rv2GBYQm7xM
 * Keys:
 *  - Use Union Find to group the nodes based on the values.
 *  - Union Find reference: https://www.youtube.com/watch?v=FXWRE67PLL0
 *
 * I need to revisit this in the future because it drives me crazy.
 */
use std::cmp::max;

struct Solution;

impl Solution {
    pub fn find(val: usize, root: &Vec<usize>) -> usize {
        if val == root[val] {
            return val;
        }

        return Solution::find(root[val], root);
    }

    pub fn number_of_good_paths(vals: Vec<i32>, edges: Vec<Vec<i32>>) -> i32 {
        let vals_len = vals.len();
        let mut output: i32 = vals_len as i32;
        if vals_len == 1 {
            return 1;
        }

        let mut sorted_edges: Vec<Vec<i32>> = edges;
        sorted_edges.sort_by(|x, y| {
            max(vals[x[0] as usize], vals[x[1] as usize])
                .cmp(&max(vals[y[0] as usize], vals[y[1] as usize]))
        });

        let mut paths = vec![1; vals_len];
        let mut root = (0..vals_len).collect();
        for e in sorted_edges {
            let a = Solution::find(e[0] as usize, &mut root);
            let b = Solution::find(e[1] as usize, &mut root);
            if vals[a] == vals[b] {
                // found path
                output += paths[a] as i32 * paths[b] as i32;
                root[a] = b;
                paths[b] += paths[a];
            } else if vals[a] > vals[b] {
                root[b] = a;
            } else {
                root[a] = b;
            }
        }

        return output;
    }
}

#[test]
fn test_case_1() {
    let vals: Vec<i32> = vec![1, 3, 2, 1, 3];
    let edges: Vec<Vec<i32>> = vec![vec![0, 1], vec![0, 2], vec![2, 3], vec![2, 4]];
    let expected = 6;
    let result = Solution::number_of_good_paths(vals, edges);
    assert_eq!(expected, result)
}
