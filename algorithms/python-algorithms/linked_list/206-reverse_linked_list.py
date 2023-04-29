from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    # 2023-04-29
    # TC: O(n) - Iterate through linked list once
    # SC: O(n) - Stack frames need to be created n times len of list
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head or not head.next:
            return head

        revLL = self.reverseList(head.next)
        head.next.next = head
        head.next = None

        return revLL

    # 2023-04-29
    # TC: O(n) - Iterate through linked list
    # SC: O(1) - Pointer logic only no need for extra DS
    def reverseListIterative(self, head: Optional[ListNode]) -> Optional[ListNode]:
        curr = head
        prev = None

        while curr:
            next = curr.next
            curr.next = prev
            prev = curr
            curr = next

        return prev


def main():
    pass


if __name__ == "__main__":
    main()
