""""
1. Two Sum

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
"""


# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        # this is a head node, allows improvements to code readability by avoiding a do-while-like structure
        answer = ListNode(0)
        cur_node = answer
        carry = 0

        while l1 or l2:
            val1 = l1.val if l1 else 0
            val2 = l2.val if l2 else 0
            cur_sum = val1 + val2

            if carry == 1:
                carry = 0
                cur_sum += 1

            if cur_sum > 9:
                carry = 1
                cur_sum = cur_sum % 10

            cur_node.next = ListNode(cur_sum)
            cur_node = cur_node.next

            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None

        if carry == 1:
            cur_node.next = ListNode(1)

        return answer.next
