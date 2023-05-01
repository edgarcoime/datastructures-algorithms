from typing import Optional


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


# 2023-05-01
# TC: O(n) - Iterate through all nodes at least once
# SC: O(n) - Stack frame call for potentially worse case all nodes
class Solution:
    def minDepth(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0

        children = [root.left, root.right]
        # Check to see if in leaf node
        if not any(children):
            return 1

        min_depth = float("inf")
        for c in children:
            if c:
                min_depth = min(self.minDepth(c), min_depth)

        return int(min_depth + 1)


def main():
    pass


if __name__ == "__main__":
    main()
