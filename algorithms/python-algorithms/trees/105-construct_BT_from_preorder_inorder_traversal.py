from typing import List, Optional


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


# 2023-05-04
# TC: O(n) - Call stack will have to be created for each node in final tree
# SC: O(n) - Call stack will be created for each node
class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        if not preorder or not inorder:
            return None

        root = TreeNode(preorder[0])
        mid = inorder.index(preorder[0])
        # Partition Tree | preorder skip 0th el and include middle | inorder
        # Partition Tree:
        #   - Preorder: Skip 0th el to inclusive of mid
        #   - Inorder: Start of list to exclusive of mid
        root.left = self.buildTree(preorder[1 : mid + 1], inorder[:mid])
        root.right = self.buildTree(preorder[mid + 1 :], inorder[mid + 1 :])
        return root


def main():
    preorder = [3, 9, 11, 8, 20, 15, 7]
    inorder = [11, 9, 8, 3, 15, 20, 7]
    sol = Solution()
    print(sol.buildTree(preorder, inorder))


if __name__ == "__main__":
    main()
