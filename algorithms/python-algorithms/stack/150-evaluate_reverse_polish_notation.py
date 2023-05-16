from typing import List


# 2023-05-15
# TC: O(n) - Iterate through tokens only once
# SC: O(n) - Create a stack that could potentially have all tokens
class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stack = []

        for c in tokens:
            # Order doesn't matter
            if c == "+":
                stack.append(stack.pop() + stack.pop())
            elif c == "*":
                stack.append(stack.pop() * stack.pop())
            # Order matters
            elif c == "-":
                a, b = stack.pop(), stack.pop()
                stack.append(b - a)
            elif c == "/":
                a, b = stack.pop(), stack.pop()
                # Division needs to be towards 0
                # Python decimal (default) division rounds down always
                stack.append(int(b / a))
            else:
                stack.append(int(c))

        return stack.pop()


def main():
    pass


if __name__ == "__main__":
    main()
