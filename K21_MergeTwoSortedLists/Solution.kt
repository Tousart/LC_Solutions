package K21_MergeTwoSortedLists

class ListNode(var `val`: Int) {
    var next: ListNode? = null
}

class Solution {
    fun mergeTwoLists(list1: ListNode?, list2: ListNode?): ListNode? {
        val resultList: ListNode
        var i = list1
        var j = list2

        if (i == null)
            return j

        if (j == null)
            return i

        if (i.`val` < j.`val`) {
            resultList = ListNode(i.`val`)
            i = i.next
        } else {
            resultList = ListNode(j.`val`)
            j = j.next
        }

        var current = resultList

        while (i != null && j != null) {
            if (i.`val` < j.`val`) {
                current.next = i
                i = i.next
            } else {
                current.next = j
                j = j.next
            }

            current = current.next!!
        }

        while (i != null) {
            current.next = i
            i = i.next
            current = current.next!!
        }

        while (j != null) {
            current.next = j
            j = j.next
            current = current.next!!
        }

        display(resultList)

        return resultList
    }

    private fun display(list: ListNode?) {
        var current = list
        while (current != null) {
            println(current.`val`)
            current = current.next
        }
    }
}



fun main() {
    val list1 = ListNode(1)
    list1.next = ListNode(2)
    list1.next!!.next = ListNode(3)

    val list2 = ListNode(4)
    list2.next = ListNode(5)
    list2.next!!.next = ListNode(6)

    val sl = Solution()
    sl.mergeTwoLists(list1, list2)
}