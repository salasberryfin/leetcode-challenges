/*
 * Another hell of hells
 * Implemented a naive solution using recursion that seems to be correct but
 * is unusable for larger inputs because it goes over the same square in the grid
 * multiple times. In order to avoid exceeding the time limit, there must be a
 * solution that stores the value of the already checked squares and we do not go
 * go over them again.
 *
 * Keys:
 *      - First, move from recursion to iteration.
 *      - Identify that last bottom and last right have a single solution: you can only reach the
 *      end either going always right or always down.
 *      - Each cell's number of paths is equal to the sum of the cells that you can travel to:
 *          #paths[(x, y)] = #paths[x+1, y] + #paths[x, y+1]
 *          Considering (x+1, y) and (x, y+1) are both within the boundaries of the grid, which is
 *          always for cells out of the last bottom and last right
 * - My solution is far from ideal but gives the correct solution
 * - The idea is to create a HashMap that stores the #paths associated with each cell
 * - Initialize the bottom and right cells to all 1s
 * - Start the loop from the penultimate cell, that is [m-2, n-2]
 * - Keep going backwards and storing the values in the HashMap
 * - Check if the key is already stored in the HashMap first, and avoid repeating squares
 * - Return the value of the [0, 0] key in the HashMap
*/
use std::collections::HashMap;

struct Solution;

impl Solution {
    pub fn get_paths_slow(m: i32, n: i32, pos: [i32; 2]) -> i32{
        let mut moves: i32 = 0;

        if pos[0] == m-1 && pos[1] == n-1 {
            moves = 1;
            return moves;
        }
        // down
        if pos[0] + 1 < m {
            moves = moves + Solution::get_paths_slow(m, n, [pos[0]+1, pos[1]]);
        }
        // right
        if pos[1] + 1 < n {
            moves = moves + Solution::get_paths_slow(m, n, [pos[0], pos[1]+1]);
        }

        return moves;
    }

    pub fn unique_paths(m: i32, n: i32) -> i32 {
        let mut known: HashMap<[i32; 2], i32> = HashMap::new();
        let mut stack: Vec<[i32; 2]> = Vec::new();
        let mut current_pos: [i32; 2];
        let pos: [i32; 2] = [m-1, n-1];

        if m == 1 || n == 1 {
            return 1;
        }

        for i in 0..n {
            known.insert([m-1, i], 1);
        }
        for i in 0..m {
            known.insert([i, n-1], 1);
        }

        stack.push([pos[0]-1, pos[1]-1]);
        while stack.len() > 0 {
            current_pos = stack.first().cloned().unwrap();
            stack.remove(0);
            if !known.contains_key(&current_pos) {
                let value = known[&[current_pos[0]+1, current_pos[1]]]+known[&[current_pos[0], current_pos[1]+1]];
                known.insert(current_pos, value);
                if current_pos[0]-1 >= 0 && current_pos[1]-1 >= 0 {
                    stack.push([current_pos[0]-1, current_pos[1]]);
                    stack.push([current_pos[0], current_pos[1]-1]);
                } else if current_pos[0]-1 >= 0 {
                    stack.push([current_pos[0]-1, current_pos[1]]);
                } else if current_pos[1]-1 >= 0 {
                    stack.push([current_pos[0], current_pos[1]-1]);
                }
            }
        }

        return known[&[0, 0]];
    }
}

#[test]
fn test_example1() {
    let m = 3;
    let n = 7;
    let expected = 28;
    let result = Solution::unique_paths(m, n);
    assert_eq!(expected, result);
}

#[test]
fn test_example2() {
    let m = 3;
    let n = 2;
    let expected = 3;
    let result = Solution::unique_paths(m, n);
    assert_eq!(expected, result);
}

#[test]
fn test_example3() {
    let m = 1;
    let n = 2;
    let expected = 1;
    let result = Solution::unique_paths(m, n);
    assert_eq!(expected, result);
}
