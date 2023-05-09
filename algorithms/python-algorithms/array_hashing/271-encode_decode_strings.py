from typing import List


# 2023-05-09
# TC:
# - encode: O(n) - Iterate through list once to generate str
# - decode: O(n) - Iterate through string once
# SC:
# - encode: O(1) - no ds needed (If not counting return string)
# - decode: O(n) - return ds for string
class Codec:
    def encode(self, strs: List[str]) -> str:
        """Encodes a list of strings to a single string."""
        res = ""
        for s in strs:
            res += str(len(s)) + "#" + s
        return res

    def decode(self, s: str) -> List[str]:
        """Decodes a single string to a list of strings."""
        res, i = [], 0

        while i < len(s):
            j = i
            while s[j] != "#":
                j += 1
            length = int(s[i:j])
            res.append(s[j + 1 : j + 1 + length])
            i = j + 1 + length
        return res


# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.decode(codec.encode(strs))


def main():
    pass


if __name__ == "__main__":
    main()
