from typing import List


# 2023-05-04
# TC: O(n) - Iterate through nums array only once
# SC: O(1) - Storing max integers only, no extra ds needed
class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        res = nums[0]
        curMin, curMax = 1, 1

        for n in nums:
            tmp = curMax * n  # Use old curMax

            curMax = max(n, tmp, n * curMin)
            curMin = min(n, tmp, n * curMin)

            res = max(res, curMax)

        return res


def main():
    pass


if __name__ == "__main__":
    main()
