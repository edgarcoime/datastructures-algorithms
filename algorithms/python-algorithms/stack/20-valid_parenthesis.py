# 2023-05-11
# TC: O(n) - Iterate through length of string
# SC: O(n) - Stack could potentially be as long as string if invalid
class Solution:
    def isValid(self, s: str) -> bool:
        if len(s) <= 1:
            return False

        stack = []
        pairs = {
            ")": "(",
            "]": "[",
            "}": "{",
        }

        for c in s:
            pair = pairs.get(c)

            # If not found then opening brace
            if not pair:
                stack.append(c)
                continue

            # If found but stack is empty then invalid
            if len(stack) == 0:
                return False

            # Check if valid
            if pair != stack[-1]:
                return False

            # Valid
            stack.pop()

        return len(stack) == 0


def main():
    pass


if __name__ == "__main__":
    main()
