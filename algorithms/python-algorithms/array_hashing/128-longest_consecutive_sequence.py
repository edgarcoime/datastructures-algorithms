from typing import List


# 2023-05-04
# TC: O(n) - Iterating through list of nums only once
# SC: O(n) - Create a new set DS that could potentially be len of nums
class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        numSet = set(nums)
        longest = 0

        for n in nums:
            # check if its the start of a sequence
            length = 0
            if (n - 1) not in numSet:
                while (n + length) in numSet:
                    length += 1
                longest = max(length, longest)

        return longest


def main():
    pass


if __name__ == "__main__":
    main()
