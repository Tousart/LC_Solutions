package Kotlin_153FindMinimumInRotatedSortedArray

class Solution {
    fun findMin(nums: IntArray): Int {
        var left = 0
        var right = nums.size - 1
        var mid: Int

        while (left < right) {
            mid = (left + right) / 2

            if (nums[right] > nums[mid]) {
                right = mid
            }

            else {
                left = mid + 1
            }
        }

        return nums[left]
    }
}



fun main() {
    val nums: IntArray = intArrayOf(5, 1, 2, 3, 4)
    val sl = Solution()
    println(sl.findMin(nums))
}