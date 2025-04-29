# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseList(self, head):
        if head is None:
            return
        
        current = head
        while current.next:
            nextt = current.next
            current.next = nextt.next
            nextt.next = head
            head = nextt
        
        return head

    def print_lst(self, head):
        current = head
        while current:
            print(current.val, end="")
            current = current.next
        print()


sl = Solution()
lst = ListNode(1, ListNode(2, ListNode(3, ListNode(4, ListNode(5)))))
sl.print_lst(lst)
lst = sl.reverseList(lst)
sl.print_lst(lst)
