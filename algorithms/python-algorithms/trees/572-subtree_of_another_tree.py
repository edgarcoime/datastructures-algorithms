from typing import Optional


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


# 2023-04-29
# TC: O(s*t)
#     - s: Full root trree
#     - t: Sub tree to compare if insde
# SC: (O(s)) - Call stack for isSubtree main array
class Solution:
    def isSubtree(self, root: Optional[TreeNode], subRoot: Optional[TreeNode]) -> bool:
        if not subRoot:
            return True
        if not root:
            return False
        if self.isSameTree(root, subRoot):
            return True

        return self.isSubtree(root.left, subRoot) or self.isSubtree(root.right, subRoot)

    def isSameTree(self, s, t) -> bool:
        if not s and not t:
            return True
        if s and t and s.val == t.val:
            return self.isSameTree(s.left, t.left) and self.isSameTree(s.right, t.right)
        return False


def main():
    pass


if __name__ == "__main__":
    main()
