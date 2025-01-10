"""
:type nums: List[int]
:rtype: int
"""



class Solution(object):
    def removeDuplicates(self, nums):
        new_nums = sorted(list(set(nums)))
        k = len(new_nums)
        nums[:k] = new_nums
        return k


sl = Solution()
nums = [1, 1, 2, 2, 3]
print(sl.removeDuplicates(nums))
print(nums)