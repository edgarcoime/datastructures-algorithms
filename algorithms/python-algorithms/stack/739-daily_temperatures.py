from typing import List


# 2023-05-11
# TC: O(n) - Iterate through temps once but keep running tally using stack
# SC: O(n) - Potentially could be same len as temperatures
class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        res = [0 for _ in range(len(temperatures))]
        stack = []  # pair: [temp, idx]

        for i, t in enumerate(temperatures):
            while stack and t > stack[-1][0]:
                stackT, stackIdx = stack.pop()
                res[stackIdx] = i - stackIdx
            stack.append([t, i])
        return res


def main():
    pass


if __name__ == "__main__":
    main()
