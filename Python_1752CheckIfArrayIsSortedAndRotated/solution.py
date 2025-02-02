"""
:type nums: List[int]
:rtype: bool
"""

class Solution(object):
    def check(self, nums):
        # if len(nums) <= 2:
        #     return True

        left = 0
        right = len(nums) - 1

        while left <= right:
            mid = (left + right) // 2

            if left == right:
                return True

            elif nums[mid] <= nums[right] and nums[mid] >= nums[left]:
                right = mid

            elif (nums[mid] >= nums[right] and nums[mid] >= nums[left]) or \
                (nums[mid] <= nums[right] and nums[mid] <= nums[left]):
                if nums[left] >= nums[right]:
                    left = mid + 1
                else:
                    return False

            else:
                return False

        return False


sl = Solution()
a = [5, 1, 3, 2]
print(sl.check(a))