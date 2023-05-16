from typing import List


# 2023-05-16
# TC: O(logn) - Modified binary search
# SC: O(1) - Pointers are only used
class Solution:
    def findMin(self, nums: List[int]) -> int:
        lo, hi = 0, len(nums) - 1
        res = float("inf")

        while lo <= hi:
            # If hit subarray that's already sorted
            if nums[lo] < nums[hi]:
                res = min(res, nums[lo])
                break

            # If array is not sorted
            m = lo + (hi - lo) // 2
            res = min(res, nums[m])

            if nums[lo] <= nums[m]:
                lo = m + 1
            else:
                hi = m - 1

        return int(res)


def main():
    pass


if __name__ == "__main__":
    main()
