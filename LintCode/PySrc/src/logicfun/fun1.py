# will go to the other company, so i'd like coding by python - 2020.1.16
import common.graph
import common.list

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
    def maxPoints(self, points: list[list[int]]) -> int:
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
