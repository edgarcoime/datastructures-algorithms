from typing import Optional


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


# 2023-05-01
# TC: O(n) - Iterate through all nodes once
# SC: O(n) - DFS call stack has to iterate through all nodes
class Solution:
    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        res = [0]  # Want pass by reference
        self.dfs(root, res)
        return res[0]

    def dfs(self, root, acc):
        if not root:
            return -1

        left = self.dfs(root.left, acc)
        right = self.dfs(root.right, acc)

        acc[0] = max(acc[0], 2 + left + right)

        return 1 + max(left, right)


def main():
    pass


if __name__ == "__main__":
    main()
