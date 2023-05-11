# 2023-05-11
# n = number of elements inserted
# SC: O(2n) - Generate 2 stacks instead of 1 to keep track of min
class MinStack:
    def __init__(self):
        self.stack = []
        self.minStack = []

    def push(self, val: int) -> None:
        if len(self.stack) == 0:
            self.stack.append(val)
            self.minStack.append(val)
            return

        # Insert into stack, check if min is less than val
        self.stack.append(val)
        curMin = self.getMin()
        if val < curMin:
            self.minStack.append(val)
        else:
            self.minStack.append(curMin)

    def pop(self) -> None:
        self.stack.pop()
        self.minStack.pop()

    def top(self) -> int:
        return self.stack[-1]

    def getMin(self) -> int:
        return self.minStack[-1]


# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(val)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()


def main():
    pass


if __name__ == "__main__":
    main()
