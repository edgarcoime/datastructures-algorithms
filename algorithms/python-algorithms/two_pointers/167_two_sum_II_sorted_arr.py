from typing import List


# 2023-05-09
# TC: O(n) - Iterate through list once
# SC: O(1) - Pointers are only used since sorted arr of nums
class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        l = 0
        r = len(numbers) - 1

        while l < r:
            curSum = numbers[l] + numbers[r]

            if target < curSum:
                r -= 1
                continue

            elif curSum < target:
                l += 1
                continue

            # HAS to be the same curSum == target
            else:
                return [l + 1, r + 1]

        return []


def main():
    pass


if __name__ == "__main__":
    main()
