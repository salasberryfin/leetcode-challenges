from typing import List

class Solution:
    def runningSum(self, nums: List[int]) -> List[int]:
        result = []
        prev = 0
        for num in nums:
            result.append(num+prev)
            prev = num + prev

        return result
