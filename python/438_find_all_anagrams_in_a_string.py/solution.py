"""
This is very easy using the sorted strings approach
    - Sort the p string: the target
    - Define the window: this is a string to which characters will be appended each cycle
    - Set the start of the current potential anagram to 0
    - Loop over the characters in s
        - For each cycle, append the character to window
        - When len(window) == len(p) -> this may be an anagram
            - Sort window and compare to sorted_p
            - If anagram, add start to result list
            - Pop first item from window
            - Increase start index by 1
"""
from typing import List


class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:
        if len(p) > len(s):
            return []

        result = []
        sorted_p = "".join(sorted(p))

        window = ""
        start = 0
        for char in s:
            window += char
            if len(window) == len(p):
                sorted_window = "".join(sorted(window))
                if sorted_window == sorted_p:
                    result.append(start)
                window = window[1:]
                start = start + 1

        return result

sol = Solution()

s = "cbaebabacd"
p = "abc"
expected = [0, 6]
result = sol.findAnagrams(s, p)
assert expected == result

s = "abab"
p = "ab"
expected = [0, 1, 2]
result = sol.findAnagrams(s, p)
assert expected == result
