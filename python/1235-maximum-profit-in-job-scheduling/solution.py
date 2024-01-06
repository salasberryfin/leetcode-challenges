import bisect

from typing import List
class Solution:
    def jobScheduling(self, startTime: List[int], endTime: List[int], profit: List[int]) -> int:
        compact = sorted(zip(startTime, endTime, profit))

        stored = {}
        def recursive(current):
            if current == len(compact):
                return 0
            # avoid iterating more than necessary using memoization
            if current in stored:
                return stored[current]

            # this job is not selected
            result = recursive(current + 1)

            # this job is selected
            next = bisect.bisect(compact, (compact[current][1], -1, -1))
            stored[current] = result = max(result, compact[current][2] + recursive(next))

            return result


        return recursive(0)
