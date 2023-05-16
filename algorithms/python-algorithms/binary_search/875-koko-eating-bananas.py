import math
from typing import List


# 2023-05-16
# TC: O(log(max(p)) * p) - have to calculate rate inner loop but rate is not much less than nested for
# SC: O(1) - no extra memory needed
class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        lo, hi = 1, max(piles)
        res = hi  # could use inf but highest possible should be max(p)

        while lo <= hi:
            k = lo + (hi - lo) // 2
            hours = 0
            for p in piles:
                hours += math.ceil(p / k)

            if hours <= h:
                res = min(res, k)
                hi = k - 1
            else:
                lo = k + 1

        return res


def main():
    pass


if __name__ == "__main__":
    main()
