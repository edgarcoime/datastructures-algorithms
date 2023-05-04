from typing import List, Optional


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


# 2023-05-04
# TC: O(n) - Traversing through all nodes in Tree once
# SC: O(n)? - Traversing through each element in tree requires a call stack
class Solution:
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        preorder = []
        self.genInorder(root, preorder)
        return preorder[k - 1]

    def genInorder(self, root, acc: List):
        if root:
            self.genInorder(root.left, acc)
            acc.append(root.val)
            self.genInorder(root.right, acc)


def main():
    root = TreeNode(
        3, TreeNode(1, None, TreeNode(2, None, None)), TreeNode(4, None, None)
    )
    sol = Solution()
    sol.kthSmallest(root, 2)


if __name__ == "__main__":
    main()
