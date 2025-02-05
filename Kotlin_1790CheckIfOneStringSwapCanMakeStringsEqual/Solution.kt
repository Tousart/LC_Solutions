package Kotlin_1790CheckIfOneStringSwapCanMakeStringsEqual

class Solution {
    fun areAlmostEqual(s1: String, s2: String): Boolean {
        var count = 0
        var i = 0
        var ind1 = -1
        var ind2 = -1

        while (i != s1.length) {
            if (s1[i].equals(s2[i])) {
                i++
            } else {
                ind2 = ind1
                ind1 = i
                count++
                i++
            }
        }

        if ((count == 2 &&
            s1[ind1] == s2[ind2] &&
            s1[ind2] == s2[ind1]) || (count == 0)) {
            return true
        }

        return false
    }
}


fun main() {
    val sl = Solution()
    val s1 = "4ff5"
    val s2 = "5ff4"
    println(sl.areAlmostEqual(s1, s2))
}