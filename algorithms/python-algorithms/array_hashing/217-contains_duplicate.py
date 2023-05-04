from typing import List


# 2023-05-04
# TC: O(n) - Iterate through list of nums once
# SC: O(n) - Have to create a set DS
class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        lookup = set()

        for n in nums:
            if n in lookup:
                return True
            lookup.add(n)

        return False


def main():
    pass


if __name__ == "__main__":
    main()
