from typing import List


# 2023-05-15
# TC: O(logmn) - where m n are the dimensions of the matrix
# SC: O(1) - Pointer logic only
class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        w, h = len(matrix[0]), len(matrix)
        lo, hi = 0, (w * h) - 1

        while lo <= hi:
            mid = lo + (hi - lo) // 2
            x, y = self.getRowCol(mid, w)
            midVal = matrix[y][x]

            # Check if value found
            if target == midVal:
                return True

            if midVal < target:
                lo = mid + 1
            else:
                hi = mid - 1

        return False

    def getRowCol(self, idx, w):
        return [
            idx % w,
            idx // w,
        ]


def main():
    soln = Solution()
    matrix = [[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 60]]
    target = 3
    soln.searchMatrix(matrix, target)


if __name__ == "__main__":
    main()
