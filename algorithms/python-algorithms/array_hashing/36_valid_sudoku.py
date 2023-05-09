import collections
from typing import List


# 2023-05-09
# TC: O(9^2) - Traverse through matrix through nested for loops
# SC: O(9*3) - Generate 3 dictionaries size 9 for checking dups
class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        cols = collections.defaultdict(set)
        rows = collections.defaultdict(set)
        squares = collections.defaultdict(set)

        for r in range(9):
            for c in range(9):
                # If empty skip
                if board[r][c] == ".":
                    continue
                if (
                    board[r][c] in rows[r]
                    or board[r][c] in cols[c]
                    or board[r][c] in squares[r // 3, c // 3]
                ):
                    return False
                cols[c].add(board[r][c])
                rows[r].add(board[r][c])
                squares[(r // 3, c // 3)].add(board[r][c])
        return True


def main():
    pass


if __name__ == "__main__":
    main()
