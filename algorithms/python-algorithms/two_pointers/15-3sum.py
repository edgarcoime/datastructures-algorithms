from typing import List


# 2023-05-09
# TC: O(nlogn + O(n^2)) => O(n^2) - Sorting algo + nested loop algorithm for three sum soln
# SC: O(1) or O(n) - Dependent on sorting algo
class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        result = []
        for index in range(len(nums)):
            if nums[index] > 0:
                break
            if index > 0 and nums[index] == nums[index - 1]:
                continue
            left, right = index + 1, len(nums) - 1
            while left < right:
                if nums[left] + nums[right] < 0 - nums[index]:
                    left += 1
                elif nums[left] + nums[right] > 0 - nums[index]:
                    right -= 1
                else:
                    result.append(
                        [nums[index], nums[left], nums[right]]
                    )  # After a triplet is appended, we try our best to incease the numeric value of its first element or that of its second.
                    left += 1  # The other pairs and the one we were just looking at are either duplicates or smaller than the target.
                    right -= 1  # The other pairs are either duplicates or greater than the target.
                    # We must move on if there is less than or equal to one integer in between the two nums.
                    while nums[left] == nums[left - 1] and left < right:
                        left += 1  # The pairs are either duplicates or smaller than the target.
        return result


# 2023-05-15
# TC: O(nlogn) + O(n^2) - sorting func + nested loop in solving twosum II
# SC: O(1) or O(n) - depending on sorting algorithm in library
class Solution2:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        res = []

        for i, n in enumerate(nums):
            # Base case; since sorted if greater than 0 then impossible to equal 0
            if n > 0:
                break
            if i > 0 and nums[i] == nums[i - 1]:
                continue

            l, r = i + 1, len(nums) - 1
            while l < r:  # find the complement of n (0 - n)
                if nums[l] + nums[r] < 0 - n:
                    l += 1
                elif nums[l] + nums[r] > 0 - n:
                    r -= 1
                else:
                    res.append([n, nums[l], nums[r]])
                    l += 1
                    r -= 1

                    while l < r and nums[l] == nums[l - 1]:
                        l += 1

        return res


def main():
    pass


if __name__ == "__main__":
    main()
