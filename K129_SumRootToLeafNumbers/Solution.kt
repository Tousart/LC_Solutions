package K129_SumRootToLeafNumbers

class TreeNode(var `val`: Int) {
        var left: TreeNode? = null
        var right: TreeNode? = null
    }


class Solution {
    private var result: Int = 0

    fun sumNumbers(root: TreeNode?): Int {
        if (root == null) {
            return 0
        }

        deep(root, 0)

        return result
    }

    private fun deep(current: TreeNode?, number: Int) {
        if (current == null) {
            return
        }

        val tmp = number * 10 + current.`val`

        if (current.left == null && current.right == null) {
            result += tmp
            return
        }

        deep(current.left, tmp)
        deep(current.right, tmp)
    }
}


fun main() {
    val treeNode: TreeNode = TreeNode(1)
    treeNode.left = TreeNode(2)
    treeNode.right = TreeNode(3)
    val sl: Solution = Solution()
    println(sl.sumNumbers(treeNode))
}