"""
:type nums: List[int]
:rtype: List[str]
"""


class Solution(object):
    def summaryRanges(self, nums):
        if len(nums) == 0:
            return []

        result = []
        elem1, elem2 = nums[0], nums[0]
        i = 0
        while i <= (len(nums) - 2):
            if nums[i + 1] == (elem2 + 1):
                i += 1
            else:
                if elem1 != elem2:
                    result.append(str(elem1) + "->" + str(elem2))
                else:
                    result.append(str(elem1))
                i += 1
                elem1 = nums[i]           
            elem2 = nums[i]
        
        if elem1 != elem2:
            result.append(str(elem1) + "->" + str(elem2))
        else:
            result.append(str(elem1))
        
        return result


sl = Solution()
print(sl.summaryRanges([0,1,2,4,5,7]))