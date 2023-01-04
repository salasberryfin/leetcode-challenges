"""
I kept hitting the Time Limit exceeded on a 100000 value array because I was
using lists and list.count() to get the number of tasks of each difficulty.
A dict (hash-table) is the correct solution. Accessing an element is O(n) in the
worst case.
"""
from collections import defaultdict
from typing import List


class Solution:
    def minimumRounds(self, tasks: List[int]) -> int:
        result = 0

        count_map = defaultdict(int)
        for task in tasks:
            count_map[task] += 1

        for k in count_map:
            count = count_map[k]
            if count == 1:
                return -1
            if count%3 == 0:
                result += int(count/3)
            elif count%3 > 0:
                result += int(count/3) + 1
            else:
                result += int(count/2)

        return result
