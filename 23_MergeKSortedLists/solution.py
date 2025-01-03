"""
:type lists: List[Optional[ListNode]]
:rtype: Optional[ListNode]
"""


class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution(object):
    def mergeKLists(self, lists):
        if lists is None or len(lists) == 0:
            return None
        return self.mergeLists(lists, 0, len(lists) - 1)

    def mergeLists(self, lists, start, end):
        if start == end:
            return lists[start]

        mid = start + (end - start) // 2

        # Рекурсия для левой части
        left = self.mergeLists(lists, start, mid)

        # Рекурсия для правой части
        right = self.mergeLists(lists, mid + 1, end)

        # Объединяем отсортированные массивы
        return self.merge(left, right)

    @staticmethod
    def merge(left, right):
        head = ListNode(-1)
        temp = head
        # Пока есть и левая, и правая части
        while left is not None and right is not None:
            if left.val < right.val:
                temp.next = left
                left = left.next
            else:
                temp.next = right
                right = right.next

            temp = temp.next

        # Когда осталась только левая
        while left is not None:
            temp.next = left
            left = left.next
            temp = temp.next

        # Когда осталась только правая
        while right is not None:
            temp.next = right
            right = right.next
            temp = temp.next

        return head.next



def display(lst):
    current = lst
    while current:
        print(current.val, end=" ")
        current = current.next
    print()


sl = Solution()

# lists = [ListNode(1, ListNode(2, ListNode(4))),
#          ListNode(3, ListNode(3, ListNode(6))),
#          ListNode(2, ListNode(98))]

lists = [ListNode(), ListNode(1)]

# lists = [ListNode()]

lists = sl.mergeKLists(lists)