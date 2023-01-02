#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

/*
 * Check HEX values of capital and non-capital characters
*/

bool detectCapitalUse(char *word) {
    /*
     * RULES
     * - All letters are capitals
     * - All letters are not capitals
     * - Only the first letter is capital
    */

    int count = 0;
    int countLower = 0;
    int len = 0;
    char first = *word;
    bool capital = false;
    if(first >= 'A' && first <= 'Z'){
        char *second = word + 1;
        if(*second >= 'A' && *second <= 'Z'){
            capital = true;
        }
    }

    for(char *i = word+1; *i != '\0'; i++){
        if(!capital){
            if(*i >= 'A' && *i <= 'Z'){
                return false;
            }
        } else {
            if(*i >= 'a' && *i <= 'z'){
                return false;
            }
        }
    }

    return true;
}

int main() {
    char *word = "Gaaaa";
    printf("%s\n", word);
    bool result = detectCapitalUse(word);
    bool expected = true;
    if(result != expected){
        printf("Failed to get the correct result: expected %d but got %d\n", expected, result);
    }
    word = "USA";
    printf("%s\n", word);
    result = detectCapitalUse(word);
    expected = true;
    if(result != expected){
        printf("Failed to get the correct result: expected %d but got %d\n", expected, result);
    }
    word = "FlaG";
    printf("%s\n", word);
    result = detectCapitalUse(word);
    expected = false;
    if(result != expected){
        printf("Failed to get the correct result: expected %d but got %d\n", expected, result);
    }
    word = "flaG";
    printf("%s\n", word);
    result = detectCapitalUse(word);
    expected = false;
    if(result != expected){
        printf("Failed to get the correct result: expected %d but got %d\n", expected, result);
    }
    printf("It worked!\n");
}
