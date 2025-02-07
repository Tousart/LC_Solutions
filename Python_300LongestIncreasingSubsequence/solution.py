"""
:type nums: List[int]
:rtype: int
-10^4 <= nums[i] <= 10^4
"""


class Solution(object):
    def lengthOfLIS(self, nums):
        result = 1

        # Можно заполнить с помощью MinValue и MaxValue
        sublist = [-(10 ** 4) - 1] + [nums[0]] + [(10 ** 4) + 1] * (len(nums) - 1)

        for i in range(len(nums)):
            left = 0
            right = result + 1
            while left <= right:
                mid = (left + right) // 2

                if (nums[i] <= sublist[mid]) and (nums[i] > sublist[mid - 1]):
                    break
                
                elif nums[i] > sublist[mid]:
                    left = mid + 1

                else:
                    right = mid - 1

            sublist[mid] = nums[i]
            result = max(result, mid)

        return result



sl = Solution()
nums = [10,9,2,5,3,7,101,18]
print(sl.lengthOfLIS(nums))