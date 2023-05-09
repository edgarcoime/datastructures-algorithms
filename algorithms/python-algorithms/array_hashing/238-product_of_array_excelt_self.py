from typing import List


# TODO: Fix initialization of prefix and postfix
class Solution2:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        size = len(nums)

        prefix = [1] * size
        postfix = [1] * size
        res = [1] * size

        # Generate prefix
        for i in range(1, size - 1):
            prefix[i] = prefix[i - 1] * nums[i - 1]

        # Generate Postfix
        for i in range(size - 2, -1, -1):
            postfix[i] = postfix[i + 1] * nums[i + 1]

        # Generate result
        for i in range(size):
            res[i] = prefix[i] * postfix[i]

        return res


# 2024-05-09
# TC: O(n) - Iterate through array twice for prefix and postfix/result but no nested iterations
# SC: O(1) - Using output array to store prefix and postfix doesn't count towards SC
class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        res = [1] * len(nums)

        # Gen Prefix
        prefix = 1
        for i in range(len(nums)):
            # Store Prefix before calculation then calculate
            res[i] = prefix
            prefix = prefix * nums[i]

        postfix = 1
        for i in range(len(nums) - 1, -1, -1):
            res[i] = res[i] * postfix
            postfix = postfix * nums[i]

        return res


def main():
    pass


if __name__ == "__main__":
    main()
