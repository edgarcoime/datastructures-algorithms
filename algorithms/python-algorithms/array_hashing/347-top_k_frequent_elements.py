from typing import List


# 2023-05-11
# TC: O(n) - Iterate 3 times; generate count; generate buckets; and iterate through generated buckets to get result
# SC: O(n) - 2 DS needed (not including res) that could be length of nums
class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        count = {}
        buckets = [[] for _ in range(len(nums) + 1)]

        # Count occurences
        for n in nums:
            count[n] = 1 + count.get(n, 0)

        # Store in buckets
        for key, val in count.items():
            buckets[val].append(key)

        # Generate result
        res = []
        for i in range(len(buckets) - 1, -1, -1):
            for n in buckets[i]:
                res.append(n)
                if len(res) >= k:
                    return res

        return res


def main():
    soln = Solution()
    nums = [1, 1, 1, 2, 2, 3]
    k = 2
    soln.topKFrequent(nums, k)


if __name__ == "__main__":
    main()
