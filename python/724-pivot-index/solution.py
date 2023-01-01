from typing import List


class Solution:

    def pivotIndex(self, nums: List[int]) -> int:
        total = sum(nums)
        left = 0
        for i, val in enumerate(nums):
            if total-left-val == left:
                return i
            left += val

        return -1
