# will go to the other company(bytedance), so i'd like coding by python - 2020.1.16
import common.graph
import common.list

class Solution():

    # 133. Clone Graph
    def cloneGraph(self, node):
        nodeDict = {}
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