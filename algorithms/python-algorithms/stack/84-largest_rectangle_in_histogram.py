from typing import List


# 2023-05-15
# TC: O(n) - iterate through heights once
# SC: O(n) - Create a stack that could potentially be len(n)
class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        maxArea = 0
        # Increasing value stack where top of stack is highest height
        stack = []  # pair: (index, height)

        for i, h in enumerate(heights):
            start = i
            # Keep decreasing highest value to see how far back possible
            # As popping values calculate their potential height
            while stack and stack[-1][1] > h:
                index, height = stack.pop()
                maxArea = max(maxArea, height * (i - index))
                start = index
            stack.append((start, h))

        for i, h in stack:
            maxArea = max(maxArea, h * (len(heights) - i))

        return maxArea


def main():
    pass


if __name__ == "__main__":
    main()
