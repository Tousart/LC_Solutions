"""
:type nums: List[int]
:rtype: int
"""

from collections import defaultdict

class Solution(object):
    def tupleSameProduct(self, nums):
        compositions = defaultdict(int)
        result = 0

        for i in range(0, len(nums) - 1):
            for j in range(i + 1, len(nums)):
                composition = nums[i] * nums[j]
                compositions[composition] += 1

        for i in compositions.values():
            if i > 1:
                result += 4 * i * (i - 1)

        return result



sl = Solution()
print(sl.tupleSameProduct([2, 3, 4, 6, 8, 12]))