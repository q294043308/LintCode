//  func2.h
//  LintCode
//
//  Created by DongXuxuan on 2018/8/21.
//  Copyright © 2018 DongXuxuan. All rights reserved.

#ifndef func2
#define func2

#include "define.h"

int countPrimeSetBits(int L, int R) // LintCode 1046. Prime Number of Set Bits in Binary Representation
{
    //init prime number arr
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
      //int cnt = __builtin_popcount(i)
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

int minCostClimbingStairs(vector<int> &cost) // 1054. Min Cost Climbing Stairs
{
  int minOneStep = 0;
  int minTwoStep = MAXINTNUM;

  for (uint i = 0; i < cost.size() - 1; i++)
  {
    int tmp = minOneStep;
    minOneStep = min(minTwoStep, minOneStep + cost[i]);
    minTwoStep = tmp + cost[i + 1];
  }
  return min(minOneStep, minTwoStep);
}

int findLUSlength(string &a, string &b) // 1192. Longest Uncommon Subsequence I
{
  if (a.size() != b.size())
  {
    return a.size() > b.size() ? a.size() : b.size();
  }

  int res = 0;
  int tem = 0;
  for (uint i = 0; i < a.size(); i++)
  {
    if (a[i] == b[i])
    {
      res = max(tem, res);
      tem = 0;
      continue;
    }

    tem++;
    res = max(tem, res);
  }

  return res == 0 ? -1 : res;
}

vector<int> constructRectangle(int area) // 1209. Construct the Rectangle
{
  int lenth = (int)sqrt(area);
  int width = lenth;

  for (; width > 0; width--)
  {
    for (int i = lenth; i <= area; i++)
    {
      if (width * i > area)
      {
        break;
      }

      if (width * i == area)
      {
        return vector<int>{i, width};
      }
    }
  }
  return vector<int>{area, 1};
}

void letterCasePermutationSub(string &S, vector<string> &res, vector<int> &chrIndex, int curIndex, bool isBig)
{
  if (isBig)
  {
    if (S[chrIndex[curIndex]] >= 'a' && S[chrIndex[curIndex]] <= 'z')
    {
      S[chrIndex[curIndex]] = S[chrIndex[curIndex]] - 'a' + 'A';
    }
  }
  else
  {
    if (S[chrIndex[curIndex]] >= 'A' && S[chrIndex[curIndex]] <= 'Z')
    {
      S[chrIndex[curIndex]] = S[chrIndex[curIndex]] - 'A' + 'a';
    }
  }

  if (curIndex == chrIndex.size() - 1)
  {
    res.push_back(S);
    return;
  }
  letterCasePermutationSub(S, res, chrIndex, curIndex + 1, true);
  letterCasePermutationSub(S, res, chrIndex, curIndex + 1, false);
  return;
}

vector<string> letterCasePermutation(string &S) // 1032. Letter Case Permutation
{
  vector<string> res;
  vector<int> chrIndex;

  for (uint i = 0; i < S.size(); i++)
  {
    if ((S[i] <= 'z' && S[i] >= 'a')
      || (S[i] <= 'Z' && S[i] >= 'A'))
    {
      chrIndex.push_back(i);
    }
  }

  if (chrIndex.empty())
  {
    res.push_back(S);
    return res;
  }

  letterCasePermutationSub(S, res, chrIndex, 0, true);
  letterCasePermutationSub(S, res, chrIndex, 0, false);

  return res;
}

vector<int> getRow(int rowIndex) // 1354. Pascal's Triangle II
{
  vector<int> out;
  if (rowIndex < 0)
  {
    return out;
  }

  out.assign(rowIndex + 1, 0);
  for (int i = 0; i <= rowIndex; ++i)
  {
    if (i == 0)
    {
      out[0] = 1;
      continue;
    }
    for (int j = rowIndex; j >= 1; --j)
    {
      out[j] = out[j] + out[j - 1];
    }
  }
  return out;
}

vector<int> findWordsInit()
{
  vector<int> lowVec = vector<int>(BIG_ENGLISH_CHAR_NUM, 0);
  string oneCha = "qwertyuiop";
  string twoCha = "asdfghjkl";
  string threeCha = "zxcvbnm";
  for (auto i : oneCha)
  {
    lowVec[i - 'a'] = 1;
  }
  for (auto i : twoCha)
  {
    lowVec[i - 'a'] = 2;
  }
  for (auto i : threeCha)
  {
    lowVec[i - 'a'] = 3;
  }
  return lowVec;
}

vector<string> findWords(vector<string> &words) // 1204. Keyboard Row
{
  vector<int> lowArr = findWordsInit();
  vector<string> res;

  for (auto str : words)
  {
    int line = 0;
    bool isSameLine = true;
    for (auto cha : str)
    {
      int index = 0;

      if (cha <= 'z' && cha >= 'a')
      {
        index = cha - 'a';
      }
      else
      {
        index = cha - 'A';
      }

      if (line == 0)
      {
        line = lowArr[index];
      }
      else if (line != lowArr[index])
      {
        isSameLine = false;
        break;
      }
    }

    if (isSameLine)
    {
      res.push_back(str);
    }
  }
  return res;
}

int maxAreaOfIslandSub(vector<vector<int>> &grid, int i, int j) //
{
  grid[i][j] = 0;
  int curArea = 1;
  if (i - 1 >= 0 && grid[i - 1][j] == 1)
  {
    curArea += maxAreaOfIslandSub(grid, i - 1, j);
  }
  if (j - 1 >= 0 && grid[i][j - 1] == 1)
  {
    curArea += maxAreaOfIslandSub(grid, i, j - 1);
  }
  if (i + 1 < (int)grid.size() && grid[i + 1][j] == 1)
  {
    curArea += maxAreaOfIslandSub(grid, i + 1, j);
  }
  if (j + 1 < (int)grid[i].size() && grid[i][j + 1] == 1)
  {
    curArea += maxAreaOfIslandSub(grid, i, j + 1);
  }
  return curArea;
}

int maxAreaOfIsland(vector<vector<int>> &grid) // 1080. Max Area of Island
{
  int res = 0;

  for (uint i = 0; i < grid.size(); i++)
  {
    for (uint j = 0; j < grid[i].size(); j++)
    {
      if (grid[i][j] == 1)
      {
        res = max(res, maxAreaOfIslandSub(grid, i, j));
      }
    }
  }
  return res;
}

int diameterOfBinaryTreeSub(TreeNode<int> * root, map<TreeNode<int>*, int> &nodeMap, int &res)
{
  if (!root)
  {
    return 0;
  }

  if (nodeMap.count(root))
  {
    return nodeMap[root];
  }

  int left = diameterOfBinaryTreeSub(root->left, nodeMap, res);
  int right = diameterOfBinaryTreeSub(root->right, nodeMap, res);
  res = max(res, left + right);
  nodeMap[root] = max(left, right) + 1;
  return nodeMap[root];
}

int diameterOfBinaryTree(TreeNode<int> * root) // 1181. Diameter of Binary Tree
{
  map<TreeNode<int> *, int> nodeMap;
  int res = 0;

  diameterOfBinaryTreeSub(root, nodeMap, res);
  return res;
}

int arrangeCoins(int n) // 988. Arranging Coins
{
  // D = b^2 - 4ac; x = (-b ± √D) / 2a
  return (int)(sqrt(8 * (double)n + 1) - 1) / 2;
}

string reverseWords(string &s) // 1173. Reverse Words in a String III
{
  if (s.empty())
  {
    return s;
  }

  uint start = 0;
  for (uint i = 0;; ++i)
  {
    if (i == s.size() ||  s[i] == ' ')
    {
      // reverse
      for (uint end = i - 1; start < end; ++start, --end)
      {
        char tmp = s[start];
        s[start] = s[end];
        s[end] = tmp;
      }
      start = i + 1;
    }

    if (i == s.size())
    {
      return s;
    }
  }
  return s;
}

vector<int> nextGreaterElement(vector<int> &nums1, vector<int> &nums2) // 1206. Next Greater Element I
{
  vector<int> res;
  stack<int> st;
  unordered_map<int, int> m;
  for (int num : nums2)
  {
    while (!st.empty() && st.top() < num)
    {
      m[st.top()] = num; st.pop();
    }
    st.push(num);
  }
  for (int num : nums1)
  {
    res.push_back(m.count(num) ? m[num] : -1);
  }
  return res;
}

#endif
