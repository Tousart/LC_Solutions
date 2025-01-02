package LC19_RemoveNthNodeFromEndOfList;

class ListNode {
    int val;
    ListNode next;

    ListNode() {
    }

    ListNode(int val) {
        this.val = val;
    }

    ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }
}


public class Solution {
    public static ListNode removeNthFromEnd(ListNode head, int n) {
        int len = 0;
        int removeIndex = 0;
        ListNode current = head;

        while (current != null) {
            len++;
            current = current.next;
        }

        current = head;
        removeIndex = len - n;

        if (removeIndex == 0) {
            head = head.next;
        }

        for (int i = 0; i < removeIndex; i++) {
            assert current != null;

            if (i == (len - n - 1)) {
                if (current.next.next != null) {
                    current.next = current.next.next;
                } else {
                    current.next = null;
                }
            }

            current = current.next;
        }

        return head;
    }
}


class Test {
    public static void main(String[] args) {
        ListNode list = new ListNode(3);
        list.next = new ListNode(2);
        list.next.next = new ListNode(1);
        Solution.removeNthFromEnd(list, 3);
    }
}