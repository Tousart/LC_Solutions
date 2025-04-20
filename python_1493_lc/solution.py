"""
:type nums: List[int]
:rtype: int
"""


class Solution(object):
    def longestSubarray(self, nums):
        length1 = 0
        length2 = 0
        result = 0

        for i in nums:
            if i == 0:
                result = max(result, length1 + length2)
                length1 = length2
                length2 = 0
            else:
                length2 += 1
            
            # print("l1:", length1, "l2:", length2, "res:", result)
        
        result = max(result, length1 + length2)
        
        if result == len(nums):
            return result - 1

        return result


sl = Solution()
print(sl.longestSubarray([1,1,0,1]))




# class Solution(object):
#     def longestSubarray(self, nums):
#         nums.append(0)
#         length1 = 0
#         length2 = 0
#         result = 0

#         for i in nums:
#             if i == 0:
#                 result = max(result, length1 + length2)
#                 length1 = length2
#                 length2 = 0
#             else:
#                 length2 += 1
            
#             print("l1:", length1, "l2:", length2, "res:", result)
        
#         if result == (len(nums) - 1):
#             return result - 1

#         return result