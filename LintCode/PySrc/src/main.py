# encoding: utf-8
# import common.graph
# import common.list
# from  common.stack import Stack
# import logicfun.fun1
import sys
import json
import types
import datetime
import os

def adfsdf(**kwargs):
    for k, v in kwargs.items():
        print(k, v)
        if k == 'aa':
            print(v)

def testss(a):
    a.ll = 20
    a.a[0].append(1)

def tests2(a):
    a[0] = 20

testa = 10000

def main():
    print(str(0))
    arr = {0: [100001, 100002, 100003, 100004, 100005, 100006, 100007], 1000040: [100001]}
    print(json.dumps(arr))
    print(arr[0])
    stra = "{\"0\":[100001,100002,100003,100004,100005,100006,100007],\"1000040\":[100001]}"
    aaa = json.loads(stra)
    print(aaa)
    print(aaa['0'])
    # node1 = common.list.Node(1)
    # node2 = common.list.Node(2,node1,node1)
    # node1.next = node2
    # node1.random = node2
    # res = logicfun.fun1.Solution().reverseWords(\"  hello world!  \")
    connect_str = 'mysql://:@{consul}/{db_name}?charset={charset}&use_unicode=1'.format(consul='123123',
                                                                                        db_name='0000000',
                                                                                        charset='6666666')
    print(connect_str)
    print(testa)

    class test1(object):
        def __init__(self):
            self.ll = 10
            self.a = [[]]
        def __enter__(self):
            print("10")
        def __exit__(self, ex_type, ex_value, traceback):
            print("10")
        def _test(self):
            print("10")
            
        def aaa(self):
            self._test()
        
        ll = 0

    l = test1()
    testss(l)
    print(l.ll)
    print(l.a)
    l = [1]
    tests2(l)
    print(l[0])

    class test2(test1):
        def __init__(self):
            self.l2 = 3
        def _test(self):
            print("20")
        l2 = 0

    a = test2()
    a.aaa()

    a = [1, 1, 2]
    print(list(set(a)))
    ss = '123AD|LIVE|ROOM|CONSUMER|DEADLOCK|ADID|%s123'
    print(ss % 10)

    class Creative():
        def __init__(self, id, data_dict):
            self.id = id
            self.data_dict = data_dict

        test = 10
        id = 0
        data_dict = {}

        @classmethod
        def add(cls):
            cls.test += 1
        
        @staticmethod
        def ssprint():
            print(Creative.test)

        def add1(self):
            self.test += 1
 
    class Ref():
        def __init__(self, item_id, creative_id):
            self.item_id = item_id
            self.creative_id = creative_id

        item_id = 0
        isdel = 0
        creative_id = 0

    creatives = [Creative(id=1, data_dict={'aweme_item_id':33333}), Creative(id=2, data_dict={'aweme_item_id':34567}), Creative(id=3, data_dict={'aweme_item_id':111111})]
    creative_mapper = {c.id: c for c in creatives}
    print(creative_mapper.keys())
    ies_refs = [Ref(item_id=123456, creative_id=1), Ref(item_id=33333, creative_id=1), Ref(item_id=34567, creative_id=2), Ref(item_id=222222, creative_id=2), Ref(item_id=111111, creative_id=3)]
    delete_item_ref = {}
    for ref in ies_refs:
        creative = creative_mapper.get(ref.creative_id)
        if not creative:
            continue
        if not creative.data_dict.get('aweme_item_id') or creative.data_dict.get('aweme_item_id') != ref.item_id:
            if ref.item_id in delete_item_ref:
                delete_item_ref[ref.item_id].append(ref)
            else:
                delete_item_ref[ref.item_id] = [ref]

    if not delete_item_ref:
        print("1111")
        return
    
    all_refs = [Ref(item_id=123456, creative_id=1), Ref(item_id=33333, creative_id=1), Ref(item_id=222222, creative_id=2), Ref(item_id=34567, creative_id=2), Ref(item_id=111111, creative_id=3)]
    ref_cnt = {}
    for ref in all_refs:
        ref_cnt[ref.item_id] = ref_cnt[ref.item_id] + 1 if ref.item_id in ref_cnt else 1

    for (item_id, refs) in delete_item_ref.items():
        # ref 数量不等说明 item 还有 ref 关系, 不能发对应的删除消息
        if item_id in ref_cnt and ref_cnt[item_id] == len(refs):
            print("delete item : ", item_id)

        for ref in refs:
            ref.isdel = 1

    print("over ：",)
    for ref in ies_refs:
        print("{item_id:", ref.item_id, ", is_del:", ref.isdel, "}")


if __name__ == '__main__':
    main()