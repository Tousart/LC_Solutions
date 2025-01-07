package K11_ContainerWithMostWater

class Solution {
    fun maxArea(height: IntArray): Int {
        var i = 0
        var j = height.size - 1
        var area = minOf(height[i], height[j]) * (j - i)
        var result = area

        while (i != j) {
            if (height[i] < height[j]) {
                i++
            } else {
                j--
            }

            area = minOf(height[i], height[j]) * (j - i)

            if (result < area) {
                result = area
            }
        }

        return result
    }
}


fun main() {
    val sl = Solution()
    val array: IntArray = arrayOf(1, 3, 2, 5, 25, 24, 5).toIntArray()
    println(sl.maxArea(array))
}