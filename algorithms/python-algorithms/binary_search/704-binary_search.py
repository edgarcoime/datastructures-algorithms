from typing import List


# 2023-05-16
# TC: O(logn) - Cut search area in half every iteration
# SC: O(1) - only pointers needed
class Solution:
    def search(self, nums: List[int], target: int) -> int:
        lo, hi = 0, len(nums) - 1

        while lo <= hi:
            mid = lo + (hi - lo) // 2
            midVal = nums[mid]

            if midVal == target:
                return mid

            # Take left side
            if target < midVal:
                hi = mid - 1
            # Take right side
            else:
                lo = mid + 1

        return -1


def main():
    pass


if __name__ == "__main__":
    main()
