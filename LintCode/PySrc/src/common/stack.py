# isn't default stack
class Stack:

    # conitnue to opt, getMin
    def __init__(self):
        self.arr = []

    def push(self, x) -> None:
        self.arr.append(x)

    def pop(self) -> None:
        l_i = len(self.arr) - 1
        r_d = self.arr[l_i]
        self.arr = self.arr[:l_i]
        return r_d

    def top(self) -> int:
        return self.arr[len(self.arr) - 1]

    def getMin(self) -> int:
        if not self.arr:
            return None

        res = self.arr[0]
        for v in self.arr[1:]:
            res = min(res, v)
        return res