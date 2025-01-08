class Solution(object):
    def rotate(self, nums, k):
        k = (len(nums) - k) % len(nums)
        nums[:] = nums[k:] + nums[:k]


sl = Solution()
nums = [1, 2, 3]
sl.rotate(nums, 4)
print(nums)