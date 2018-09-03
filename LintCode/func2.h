//  func2.h
//  LintCode
//
//  Created by 董旭轩 on 2018/8/21.
//  Copyright © 2018年 董旭轩. All rights reserved.

#ifndef func2
#define func2

#include "define.h"

int countPrimeSetBits(int L, int R) // LintCode 1046. Prime Number of Set Bits in Binary Representation
{
    //init prime number arr  -- 备注：理论上这里不应该用这个宏，可题目说了是10^6，30位够了
    int res = 0;
    bool promeArr[THIRTY] = {false};
    promeArr[2] = true;
    promeArr[3] = true;
    promeArr[5] = true;
    promeArr[7] = true;
    promeArr[11] = true;
    promeArr[13] = true;
    promeArr[17] = true;
    promeArr[19] = true;

    for (int i = L; i <= R; ++i)
    {
        int cnt = __builtin_popcount(i);
        res += promeArr[cnt] ? 1 : 0;
    }
    return res;
}

char firstUniqChar(string &str) // 209. First Unique Character in a String
{
    map<char, int> strMap;
    queue<char> strQue;

    for(auto cha : str)
    {
        if(strMap.count(cha) == 0)
        {
            strQue.push(cha);
            strMap[cha] = 1;
        }
        else
        {
            strMap[cha]++;
        }
    }

    while(!strQue.empty())
    {
        char curCha = strQue.front();
        strQue.pop();
        if(strMap[curCha] == 1)
        {
            return curCha;
        }
    }
    return '0';
}

void findSecondMinimumValueSub(TreeNode<int> * root, int& res, int& secondRes)
{
    if(!root)
    {
        return;
    }

    findSecondMinimumValueSub(root->left, res, secondRes);
    if(res > root->val)
    {
        secondRes = res;
        res = root->val;
    }
    else if(res < root->val && root->val < secondRes)
    {
        secondRes = root->val;
    }
    findSecondMinimumValueSub(root->right, res, secondRes);
}

int findSecondMinimumValue(TreeNode<int> * root) // 1094. Second Minimum Node In a Binary Tree
{
    int res = MAXINTNUM;
    int secondRes = MAXINTNUM;
    findSecondMinimumValueSub(root, res, secondRes);

    if(secondRes == MAXINTNUM)
    {
        return -1;
    }

    return secondRes;
}

#endif
