from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


# 2023-05-01
# TC: O(n) - Iterate through all nodes to invert tree
# SC: O(n) - Stack frame calls would potentially be the length of n
class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        if not root:
            return root

        left = self.invertTree(root.right)
        root.right = self.invertTree(root.left)
        root.left = left

        return root


def main():
    pass


if __name__ == "__main__":
    main()
