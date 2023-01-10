/*
 * The iterative approach is notably faster than the recursive
 * Declare first two elements:
 *      - n1 = 0, n2 = 1
 * Iterate until n:
 *      - current = val1+val2
 *      - val1 = val2
 *      - val2 = current
 * Return current
*/
#include <stdio.h>

int fib(int n) {
    if(n == 0) {
        return 0;
    }

    int val1 = 0;
    int val2 = 1;
    int current = val2;

    for(int i = 1; i < n; i++){
        current = val1 + val2;
        val1 = val2;
        val2 = current;
    }

    return current;
}

int main() {

    // keep count of number of failed test cases
    int failures = 0;

    int n1 = 1;
    int expected1 = 1;
    int result1 = fib(n1);
    if(expected1 != result1) {
        fprintf(stderr, "Test case 1 - Expected %d but got %d\n", expected1, result1);
        failures++;
    } else {
        printf("Test case 1: PASS - Result: %d\n", result1);
    }

    int n2 = 3;
    int expected2 = 2;
    int result2 = fib(n2);
    if(expected2 != result2) {
        fprintf(stderr, "Test case 2 - Expected %d but got %d\n", expected2, result2);
        failures++;
    } else {
        printf("Test case 2: PASS - Result: %d\n", result2);
    }

    int n3 = 4;
    int expected3 = 3;
    int result3 = fib(n3);
    if(expected3 != result3) {
        fprintf(stderr, "Test case 3 - Expected %d but got %d\n", expected3, result3);
        failures++;
    } else {
        printf("Test case 3: PASS - Result: %d\n", result3);
    }

    int n4 = 30;
    int expected4 = 832040;
    int result4 = fib(n4);
    if(expected4 != result4) {
        fprintf(stderr, "Test case 3 - Expected %d but got %d\n", expected4, result4);
        failures++;
    } else {
        printf("Test case 3: PASS - Result: %d\n", result4);
    }

    if (failures > 0) {
        printf("Status: FAILED\n");
    } else {
        printf("Status: PASS\n");
    }


    return 1;
}
