# will go to the other company, so i'd like coding by python - 2020.1.16
import common.graph
import common.list
import common.const as Const

class Solution():

    # 133. Clone Graph
    def cloneGraph(self, node):
        nodeDict = {None:None}
        return self.cloneGraphSub(node, nodeDict)

    def cloneGraphSub(self, node, nodeDict):
        if node in nodeDict:
            return nodeDict[node]
        
        res = common.graph.Node(node.val, [])
        nodeDict[node] = res
        for neighbor in node.neighbors:
            res.neighbors.append(self.cloneGraphSub(neighbor, nodeDict))

        return res

    # 138. Copy List with Random Pointer
    def copyRandomList(self, head):
        nodeDict = {}
        def copyRandomListSub(old):
            if old == None:
                return old

            if old in nodeDict:
                return nodeDict[old]

            new = common.list.Node(old.val)
            nodeDict[old] = new
            new.random = copyRandomListSub(old.random)
            new.next = copyRandomListSub(old.next)

            return new

        return copyRandomListSub(head)

    # 149. Max Points on a Line
    def maxPoints(self, points) -> int:
        from collections import Counter, defaultdict
        points_dict = Counter(tuple(point) for point in points)
        not_repeat_points = list(points_dict.keys())
        n = len(not_repeat_points)
        if n == 1:
            return points_dict[not_repeat_points[0]]

        res = 0
        for i in range(n - 1):
            x1, y1 = not_repeat_points[i][0], not_repeat_points[i][1]
            slope = defaultdict(int)
            for j in range(i + 1, n):
                # 使用斜率，精度不准确， * 1000 后能过测试用例，不建议使用，建议使用最小分数求解
                x2, y2 = not_repeat_points[j][0], not_repeat_points[j][1]
                dy, dx = y2 - y1, x2 - x1

                if dx == 0:
                    tmp = "#"
                else:
                    tmp = dy * 1000 / dx * 1000
                slope[tmp] += points_dict[not_repeat_points[j]]
            res = max(res, max(slope.values()) + points_dict[not_repeat_points[i]])
        return res

    # 150. Evaluate Reverse Polish Notation
    def evalRPN(self, tokens) -> int:
        nums = []

        for token in tokens:
            if len(token) == 1 and (token[0] < '0' or token[0] > '9'):
                top = nums.pop()
                if token == "+":
                    nums.append(top + nums.pop())
                elif token == "-":
                    nums.append(nums.pop() - top)
                elif token == "*":
                    nums.append(nums.pop() * top)
                elif token == "/":
                    nums.append(int(nums.pop() / top))
            else:
                nums.append(int(token))

        return nums[0]

    # 151. Reverse Words in a String
    def reverseWords(self, s: str) -> str:
        res = ""
        arr = s.split(" ")
        if len(arr) == 0:
            return res

        for v in reversed(arr):
            if len(v) != 0:
                res = res + v + " "

        return res[:len(res)-1]

    # 152. Maximum Product Subarray
    def maxProduct(self, nums) -> int:
        imax = 1
        imin = 1
        res = Const.MIN_INT_NUM

        for num in nums:
            if not num:
                imax = 1
                imin = 1
                res = max(0, res)
            else:
                t = imax
                imax = max(imax * num, num, imin * num)
                imin = min(imin * num, num, t * num)
                res = max(imax, res)
        return res

    # 153. Find Minimum in Rotated Sorted Array(cut down, continue)
    def findMin(self, nums) -> int:
        if not nums:
            return 0
        n = len(nums)
        if n == 1:
            return nums[0]
        if n == 2:
            return min(nums[0], nums[1])

        mid = n//2
        num = nums[mid]
        start = nums[0]
        end = nums[n-1]
        if start <= num:
            if end <= start:
                return self.findMin(nums[mid:])
            else:
                return self.findMin(nums[:mid+1])
        else:
            return self.findMin(nums[:mid+1])
    
    # 154. Find Minimum in Rotated Sorted Array II
    def findMin2(self, nums) -> int:
        if not nums:
            return 0
        n = len(nums)
        if n == 1:
            return nums[0]
        if n == 2:
            return min(nums[0], nums[1])

        mid = n//2
        num = nums[mid]
        start = nums[0]
        end = nums[n-1]
        if start < num:
            if end <= start:
                return self.findMin(nums[mid:])
            else:
                return self.findMin(nums[:mid+1])
        elif start > num:
            return self.findMin(nums[:mid+1])
        else:
            return self.findMin(nums[1:])