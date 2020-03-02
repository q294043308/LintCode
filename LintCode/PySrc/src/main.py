import common.graph
import common.list
import logicfun.fun1
import sys


class P(object):
    def foo(self):
        print("I am a P-foo()")
 
class C(P):
    def foo(self):
        print("I am a C-foo()")
        
def main():
    # node1 = common.list.Node(1)
    # node2 = common.list.Node(2,node1,node1)
    # node1.next = node2
    # node1.random = node2
    res = logicfun.fun1.Solution().reverseWords("  hello world!  ")
    print(res)

if __name__ == '__main__':
    main()