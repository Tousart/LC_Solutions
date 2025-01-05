package K33_SearchInRotatedSortedArray

class Solution {
    fun search(nums: IntArray, target: Int): Int {
        var left = 0
        var right = nums.size - 1

        while (left <= right) {
            val midIndex = (left + right) / 2

            if (nums[midIndex] == target) {
                return midIndex
            }

            if (nums[midIndex] < nums[right]) {
                if ((nums[midIndex] < target) && (target <= nums[right])) {
                    left = midIndex + 1
                } else {
                    right = midIndex - 1
                }
            } else {
                if ((nums[left] <= target) && (target < nums[midIndex])) {
                    right = midIndex - 1
                } else {
                    left = midIndex + 1
                }
            }
        }

        return -1
    }
}


fun main() {
    val nums : IntArray = arrayOf(1, 3).toIntArray()
    val sl: Solution = Solution()
    println(sl.search(nums, 3))
}