/*
 * n steps to reach top of stairs
 * how many distinct ways to reach the top? either 1 or 2 steps per turn
 *
 * The general formula is:
 *  - climbStairs(n) = climbStairs(n-1) + climbStais(n-2)
 *
 * It matches a Fibonacci sequence where: steps(n) = fibonacci(n+1)
 *  - For 2 steps = fibonacci(3)
 *  - For 3 steps = fibonacci(4)
 *
*/
#include <stdio.h>

int climbStairs(int n) {
    if(n == 1) {
        return n;
    }

    int first = 1;
    int second = 1;

    int ways = second;
    for(int i = 1; i < n;i++) {
        ways = first + second;
        first = second;
        second = ways;
    }

    return ways;
}

int main() {
    int n, expected, result;

    n = 2;
    expected = 2;
    result = climbStairs(n);
    if(expected != result) {
        fprintf(stderr, "Expected %d but got %d\n", expected, result);
    } else {
        printf("Test case 1 (%d): PASS -> Result %d\n", n, result);
    }

    n = 3;
    expected = 3;
    result = climbStairs(n);
    if(expected != result) {
        fprintf(stderr, "Expected %d but got %d\n", expected, result);
    } else {
        printf("Test case 2 (%d): PASS -> Result %d\n", n, result);
    }

    n = 4;
    expected = 5;
    result = climbStairs(n);
    if(expected != result) {
        fprintf(stderr, "Expected %d but got %d\n", expected, result);
    } else {
        printf("Test case 3 (%d): PASS -> Result %d\n", n, result);
    }

    n = 5;
    expected = 8;
    result = climbStairs(n);
    if(expected != result) {
        fprintf(stderr, "Expected %d but got %d\n", expected, result);
    } else {
        printf("Test case 4 (%d): PASS -> Result %d\n", n, result);
    }

    return 0;
}
