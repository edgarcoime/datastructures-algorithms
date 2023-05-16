from typing import List


# 2023-05-15
# TC: O(n) - Iterate through target and position at same time, should be same len
# SC: O(n) - Have to create pairs array of len n, and stack
class Solution:
    def carFleet(self, target: int, position: List[int], speed: List[int]) -> int:
        pair = [[p, s] for p, s in zip(position, speed)]
        stack = []

        for p, s in sorted(pair, reverse=True):
            stack.append((target - p) / s)

            # Check collision
            if len(stack) >= 2 and stack[-1] <= stack[-2]:
                stack.pop()

        return len(stack)


def main():
    pass


if __name__ == "__main__":
    main()
