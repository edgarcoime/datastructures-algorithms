from typing import List


class Solution:
    # 2023-04-29
    # TC: O(n) - Iterating through elements only once
    # SC: O(1) - Only pointers used
    def trap(self, height: List[int]) -> int:
        l, r = 0, len(height) - 1
        maxL, maxR = height[l], height[r]
        res = 0

        # If it is odd the way I did it I increment at the end
        # Therefore it needs to be evaluated again
        while l <= r:
            if maxL < maxR:
                maxL = max(maxL, height[l])
                res += maxL - height[l]
                l += 1
            else:
                maxR = max(maxR, height[r])
                res += maxR - height[r]
                r += 1

        return res


def main():
    pass


if __name__ == "__main__":
    main()
