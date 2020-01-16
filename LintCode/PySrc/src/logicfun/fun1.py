# will go to the other company(bytedance), so i'd like coding by python - 2020.1.16
import common.graph

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