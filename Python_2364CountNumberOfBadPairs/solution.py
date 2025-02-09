"""
:type nums: List[int]
:rtype: int
"""

class Solution(object):
    def countBadPairs(self, nums):
        count_of_good_pairs = 0
        dict_of_pairs = {}

        for i, num in enumerate(nums):
            count_of_good_pairs += dict_of_pairs.get(num - i, 0)
            dict_of_pairs[num - i] = dict_of_pairs.get(num - i, 0) + 1

        return len(nums) * (len(nums) - 1) // 2 - count_of_good_pairs



sl = Solution()
print(sl.countBadPairs([1, 2, 3, 5]))