"""
:type nums: List[int]
:type k: int
:rtype: int
"""

class Solution(object):
    def subarraySum(self, nums, k):
        count, sm = 0, 0
        d = {0: 1}

        for i in range(nums):
            sm += i
            count += d.get(sm - k, 0)
            d[sm] = d.get(sm, 0) + 1
        
        return count
    

sl = Solution()
print(sl.subarraySum([1, -1, 0]), 0)
        

# class Solution(object):
#     def subarraySum(self, nums, k):
#         i, j, count = 0, 0, 0
#         while j < len(nums):
#             sm = 0
#             while i < len(nums) and sm < k:
#                 sm += nums[i]
#                 i += 1
#             if sm == k:
#                 count += 1
#             j += 1
#             i = j
        
#         return count