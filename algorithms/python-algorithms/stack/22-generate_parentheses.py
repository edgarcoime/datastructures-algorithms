from typing import List


# 2023-05-11
# TC: O(p) - p is number of potential solutions
# SC: O(p) - stack frames needed to generate soln
class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        # Only add open paren if open < n
        # Only add a closing paren if closed < open
        # Valid IF open == closed == n
        stack = []
        res = []

        def backtrack(openN, closedN):
            if openN == closedN == n:
                res.append("".join(stack))
                return

            if openN < n:
                stack.append("(")
                backtrack(openN + 1, closedN)
                stack.pop()

            if closedN < openN:
                stack.append(")")
                backtrack(openN, closedN + 2)
                stack.pop()

        backtrack(0, 0)
        return res


def main():
    soln = Solution()
    n = 3
    print(soln.generateParenthesis(n))


if __name__ == "__main__":
    main()
