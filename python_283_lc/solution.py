"""
:type nums: List[int]
:rtype: None Do not return anything, modify nums in-place instead.
"""

class Solution(object):
    def moveZeroes(self, nums):
        j = 0
        for i in range(len(nums)):
            if nums[i] != 0:
                nums[i], nums[j] = nums[j], nums[i]
                j += 1
        
        print(nums)


sl = Solution()
sl.moveZeroes([0, 0, 1])