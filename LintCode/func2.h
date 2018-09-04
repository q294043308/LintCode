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
        //int cnt = __builtin_popcount(i); -- 此处是Gcc编译器特有函数，计算二进制非零个数
      int cnt = 1;
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

class Interval{
public:
  int start, end;
  Interval(int start, int end) {
    this->start = start;
    this->end = end;
  }
};

bool canAttendMeetingsSort(Interval a, Interval b)
{
  return a.start <= b.start;
}

bool canAttendMeetings(vector<Interval> &intervals) // 920. Meeting Rooms
{
  sort(intervals.begin(), intervals.end(),
    [](const Interval &a, const Interval &b)
      {
        return a.start <= b.start;
      });
  for (uint i = 1; i < intervals.size(); ++i)
  {
    if (intervals[i].start < intervals[i - 1].end)
    {
      return false;
    }
  }
  return true;
}

bool isSubtreeSub(TreeNode<int> * s, TreeNode<int> * t)
{
  if (t->val != s->val)
  {
    return false;
  }

  if (t->left && s->left)
  {
    if (!isSubtreeSub(s->left, t->left))
    {
      return false;
    }
  }
  else if (t->left || s->left)
  {
    return false;
  }

  if (t->right && s->right)
  {
    if (!isSubtreeSub(s->right, t->right))
    {
      return false;
    }
  }
  else if (t->right || s->right)
  {
    return false;
  }

  return true;
}

bool isSubtree(TreeNode<int> * s, TreeNode<int> * t) // 1165. Subtree of Another Tree
{
  if (!s)
  {
    return false;
  }

  if (s->val == t->val)
  {
    if (isSubtreeSub(s, t))
    {
      return true;
    }
  }

  if (isSubtree(s->left, t))
  {
    return true;
  }
  if (isSubtree(s->right, t))
  {
    return true;
  }

  return false;
}

vector<Interval> mergeTwoInterval(vector<Interval> &list1, vector<Interval> &list2) // 839. Merge Two Sorted Interval Lists
{
  vector<Interval> res;
  sort(list1.begin(), list1.end(),
    [](const Interval &a, const Interval &b)
  {
    return a.start <= b.start;
  });
  sort(list2.begin(), list2.end(),
    [](const Interval &a, const Interval &b)
  {
    return a.start <= b.start;
  });

  uint i = 0, j = 0;
  Interval curVal(-1, -1);
  while (i < list1.size() && j < list2.size())
  {
    if (curVal.start == -1)
    {
      if (list1[i].start < list2[j].start)
      {
        curVal.start = list1[i].start;
        curVal.end = list1[i++].end;
      }
      else
      {
        curVal.start = list2[j].start;
        curVal.end = list2[j++].end;
      }
    }
    else
    {
      if (curVal.end >= list1[i].start)
      {
        curVal.end = max(curVal.end, list1[i++].end);
      }
      else if (curVal.end >= list2[j].start)
      {
        curVal.end = max(curVal.end, list2[j++].end);
      }
      else
      {
        res.push_back(curVal);
        curVal.start = -1;
      }
    }
  }

  while (i < list1.size())
  {
    if (curVal.start == -1)
    {
      curVal.start = list1[i].start;
      curVal.end = list1[i].end;
    }
    else if (curVal.end >= list1[i].start)
    {
      curVal.end = max(curVal.end, list1[i].end);
    }
    else
    {
      res.push_back(curVal);
      curVal.start = -1;
      curVal.start = list1[i].start;
      curVal.end = list1[i].end;
    }
    i++;
  }
  while (j < list2.size())
  {
    if (curVal.start == -1)
    {
      curVal.start = list2[j].start;
      curVal.end = list2[j].end;
    }
    else if (curVal.end >= list2[j].start)
    {
      curVal.end = max(curVal.end, list2[j].end);
    }
    else
    {
      res.push_back(curVal);
      curVal.start = -1;
      curVal.start = list2[j].start;
      curVal.end = list2[j].end;
    }
    j++;
  }

  if (curVal.start != -1)
  {
    res.push_back(curVal);
  }

  return res;
}

void islandPerimeterSub(vector<vector<int>> &grid, vector<vector<bool>> &gridMap, int i, int j, int &res)
{
  gridMap[i][j] = true;
  if (i - 1 >= 0 && grid[i - 1][j] == 1)
  {
    if (!gridMap[i - 1][j])
    {
      islandPerimeterSub(grid, gridMap, i - 1, j, res);
    }
  }
  else
  {
    res++;
  }

  if (j - 1 >= 0 && grid[i][j - 1] == 1)
  {
    if (!gridMap[i][j - 1])
    {
      islandPerimeterSub(grid, gridMap, i, j - 1, res);
    }
  }
  else
  {
    res++;
  }

  if (i + 1 < (int)grid.size() && grid[i + 1][j] == 1)
  {
    if (!gridMap[i + 1][j])
    {
      islandPerimeterSub(grid, gridMap, i + 1, j, res);
    }
  }
  else
  {
    res++;
  }

  if (j + 1 < (int)grid[i].size() && grid[i][j + 1] == 1)
  {
    if (!gridMap[i][j + 1])
    {
      islandPerimeterSub(grid, gridMap, i, j + 1, res);
    }
  }
  else
  {
    res++;
  }

  return;
}

int islandPerimeter(vector<vector<int>> &grid) // 1225. Island Perimeter
{
  int res = 0;
  vector<vector<bool>> gridMap;
  for (uint i = 0; i < grid.size(); i++)
  {
    gridMap.push_back(vector<bool>(grid[0].size(), false));
  }

  for (uint i = 0; i < grid.size(); i++)
  {
    for (uint j = 0; j < grid[i].size(); j++)
    {
      if (grid[i][j] == 1)
      {
        islandPerimeterSub(grid, gridMap, i, j, res);
        return res;
      }
    }
  }
  return -1;
}

int judgeTheLastNumber(string &str) // 1459. Judge the last number
{
  uint i = 0;
  for (; i < str.length() - 1;)
  {
    if (str[i] == '1')
    {
      i += 2;
    }
    else
    {
      i++;
    }
  }
  return str.length() <= i ? 2 : 1;
}



#endif
