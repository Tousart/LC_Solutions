"""
:type l1: Optional[ListNode]
:type l2: Optional[ListNode]
:rtype: Optional[ListNode]
"""

class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution(object):
    def addTwoNumbers(self, l1, l2):
        count = 0

        tail1 = l1
        tail2 = l2

        sum = str(tail1.val + tail2.val + count)
        if len(sum) == 2:
            count = int(sum[0])
            l3 = ListNode(int(sum[1]))
        else:
            count = 0
            l3 = ListNode(int(sum))

        current = l3

        while tail1.next or tail2.next or count == 1:
            if tail1.next:
                tail1 = tail1.next
            else:
                tail1 = ListNode(0)

            if tail2.next:
                tail2 = tail2.next
            else:
                tail2 = ListNode(0)

            sum = str(tail1.val + tail2.val + count)

            if len(sum) == 2:
                count = int(sum[0])
                current.next = ListNode(int(sum[1]))
            else:
                count = 0
                current.next = ListNode(int(sum))

            current = current.next

        return l3