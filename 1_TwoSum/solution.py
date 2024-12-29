class Solution(object):
    def twoSum(self, nums, target):
        nums_copy = nums[:]
        nums_copy.sort()

        ind_start = 0
        ind_end = len(nums_copy) - 1

        while (nums_copy[ind_start] + nums_copy[ind_end]) != target:
            if nums_copy[ind_start] + nums_copy[ind_end] > target:
                ind_end -= 1
            else:
                ind_start += 1

        elem1 = nums_copy[ind_start]
        elem2 = nums_copy[ind_end]

        ind1 = nums.index(elem1)

        if elem1 == elem2:
            nums[ind1] = -1

        ind2 = nums.index(elem2)

        return [ind1, ind2]




sl = Solution()
print(sl.twoSum([0,3,-3,4,-1], -1))