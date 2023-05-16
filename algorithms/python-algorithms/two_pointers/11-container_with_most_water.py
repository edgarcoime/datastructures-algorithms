from typing import List


# 2023-05-16
# TC: O(n) - Use two pointers to iterate through entire arr
# SC: O(1) - Pointer logic only no extra space
class Solution:
    def maxArea(self, height: List[int]) -> int:
        res = 0
        lo, hi = 0, len(height) - 1

        while lo < hi:
            area = (hi - lo) * min(height[lo], height[hi])
            res = max(res, area)

            if height[lo] < height[hi]:
                lo += 1
            else:
                hi -= 1

        return res

    def maxAreaBruteForce(self, height: List[int]) -> int:
        res = 0

        for lo in range(len(height)):
            for hi in range(lo + 1, len(height)):
                area = (hi - lo) * min(height[lo], height[hi])
                res = max(res, area)

        return res


def main():
    pass


if __name__ == "__main__":
    main()
