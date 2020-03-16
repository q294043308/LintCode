import common.graph
import common.list
from  common.stack import Stack
import logicfun.fun1
import sys
import json
import datetime

        
def main():
    # node1 = common.list.Node(1)
    # node2 = common.list.Node(2,node1,node1)
    # node1.next = node2
    # node1.random = node2
    # res = logicfun.fun1.Solution().reverseWords("  hello world!  ")
    s = Stack()
    s.push(-2)
    s.push(-3)
    print(s.getMin())
    print(s.pop())
    print(s.top())
    print(s.getMin())
    print(logicfun.fun1.Solution().findMin([1,0,1,1,1]))

if __name__ == '__main__':
    main()