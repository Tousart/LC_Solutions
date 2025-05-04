class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def mergeKLists(self, lists):
        if len(lists) == 0:
            return None
        return self.merge_sort(lists)

    def merge_sort(self, lists):
        if len(lists) > 1:
            mid = len(lists) // 2
            left = self.merge_sort(lists[:mid])
            right = self.merge_sort(lists[mid:])

            result = ListNode()
            curr = result

            while (left is not None) and (right is not None):
                if left.val <= right.val:
                    curr.next = left
                    left = left.next
                else:
                    curr.next = right
                    right = right.next
                curr = curr.next

            while left is not None:
                curr.next = left
                left = left.next
                curr = curr.next

            while right is not None:
                curr.next = right
                right = right.next
                curr = curr.next
            # self.print_lst(result)

            return result.next
        
        return lists[0]

    def print_lst(self, lst):
        curr = lst
        while curr:
            print(curr.val, end=" ")
            curr = curr.next
        print()


lists = [ListNode(1, ListNode(4, ListNode(5))), ListNode(1, ListNode(3, ListNode(4))), ListNode(2, ListNode(6))]
sl = Solution()
sl.print_lst(sl.mergeKLists(lists))