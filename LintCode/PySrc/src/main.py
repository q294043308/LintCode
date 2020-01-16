import common.graph
import logicfun.fun1
import sys

def main():
    node1 = common.graph.Node(1, [])
    node2 = common.graph.Node(2, [node1])
    node1.neighbors.append(node2)
    res = logicfun.fun1.Solution().cloneGraph(node1)
    print(res)

if __name__ == '__main__':
    main()