from typing import Optional


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


# 2023-05-04
# TC: O(n) - Iterate through all nodes only once
# SC: O(n) - Call stack for each node in tree
class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        return self.valid(root, float("-inf"), float("inf"))

    def valid(self, node, left, right):
        if not node:
            return True
        if not (left < node.val and node.val < right):
            return False

        return self.valid(node.left, left, node.val) and self.valid(
            node.right, node.val, right
        )


def main():
    pass


if __name__ == "__main__":
    main()
