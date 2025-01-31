package Kotlin_74SearchA2DMatrix

class Solution {
    fun searchMatrix(matrix: Array<IntArray>, target: Int) : Boolean {
        val rows: Int = matrix.size
        val cols: Int = matrix[0].size

        var left = 0
        var right = rows * cols - 1
        var mid: Int
        var x: Int
        var y: Int

        while (left <= right) {
            // Получение среднего индекса матрицы
            mid = (left + right) / 2

            // Нахождение по среднему индексу точные координаты элемента матрицы
            x = mid / cols
            y = mid % cols

            if (matrix[x][y] == target) {
                return true
            } else if (matrix[x][y] > target) {
                right = mid - 1
            } else {
                left = mid + 1
            }
        }

        return false
    }
}


fun main() {
    val sl = Solution()
    val a: Array<IntArray> = Array(3) {IntArray(0)}
    a[0] = intArrayOf(1, 3, 5, 7)
    a[1] = intArrayOf(10, 11, 16, 20)
    a[2] = intArrayOf(23, 30, 34, 60)
    println(sl.searchMatrix(a, 3))
}