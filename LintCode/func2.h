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

bool rotatedDigitsSub(int n)
{
  bool flag = false;
  while (n > 0)
  {
    switch (n % 10)
    {
    case 3: case 4: case 7: return false;
    case 1: case 0: case 8: break;
    default: flag = true;
    }
    n = n / 10;
  }
  return flag;
}

int rotatedDigits(int N) // 1028. Rotated Digits  --> I have a best answer, but I can't do it because of I don'tt have good state
{
  int num = 0;
  for (int i = 0; i <= N; i++)
  {
    num += rotatedDigitsSub(i) ? 1 : 0;
  }
  return num;
}

int trailingZeroes(int n) // 1347. Factorial Trailing Zero
{
  int res = 0;
  while (n >= 5)
  {
    n /= 5;
    res += n;
  }

  return res;
}

void tree2strSub(TreeNode<int> * t, string &res) // reduce the consume of return
{
  if (!t)
  {
    return;
  }

  res += to_string(t->val);
  if (t->left || t->right)
  {
    res.push_back('(');
    if (t->left)
    {
      tree2strSub(t->left, res);
    }
    res.push_back(')');
    if (t->right)
    {
      res.push_back('(');
      tree2strSub(t->right, res);
      res.push_back(')');
    }
  }
}

string tree2str(TreeNode<int> * t) // 1137. Construct String from Binary Tree
{
  string res;
  tree2strSub(t, res);
  return res;
}

vector<vector<int>> matrixReshape(vector<vector<int>> &nums, int r, int c) // 1170. Reshape the Matrix
{
  if ((r * c < (int)nums.size() * (int)nums[0].size()) || (r == nums.size() || c == nums[0].size()) || (r > (int)nums.size() && c > (int)nums[0].size()))
  {
    return nums;
  }

  vector<vector<int>> res;
  vector<int> curLine;
  for (auto lineNum : nums)
  {
    for (auto num : lineNum)
    {
      curLine.push_back(num);
      if (curLine.size() == c)
      {
        res.push_back(curLine);
        curLine.clear();
      }
    }
  }
  if (!curLine.empty())
  {
    res.push_back(curLine);
  }

  return res;
}

void findModeSub(TreeNode<int> * root, int &curVal, int &curNum, int &maxNum, vector<int> &res)
{
  if (!root)
  {
    return;
  }

  if (root->left)
  {
    findModeSub(root->left, curVal, curNum, maxNum, res);
  }

  if (maxNum == -1)
  {
    maxNum = 0;
    curNum = 1;
    curVal = root->val;
  }
  else
  {
    if (curVal == root->val)
    {
      curNum++;
    }
    else
    {
      if (curNum > maxNum)
      {
        maxNum = curNum;
        res.clear();
        res.push_back(curVal);
      }
      else if (curNum == maxNum)
      {
        res.push_back(curVal);
      }
      curNum = 1;
      curVal = root->val;
    }
  }

  if (root->right)
  {
    findModeSub(root->right, curVal, curNum, maxNum, res);
  }
}

vector<int> findMode(TreeNode<int> * root) // 1203. Find Mode in Binary Search Tree
{
  vector<int> res;
  int curVal = 0;
  int curNum = -1;
  int maxNum = -1;
  findModeSub(root, curVal, curNum, maxNum, res);
  if (maxNum != -1)
  {
    if (curNum == maxNum)
    {
      res.push_back(curVal);
    }
    else if (curNum > maxNum)
    {
      res.clear();
      res.push_back(curVal);
    }
  }
  return res;
}

vector<string> findRelativeRanks(vector<int> &nums) // 1200. Relative Ranks
{
  int n = nums.size(), cnt = 1;
  vector<string> res(n, "");
  map<int, int> m;
  for (int i = 0; i < n; ++i)
  {
    m[nums[i]] = i;
  }
  for (auto it = m.rbegin(); it != m.rend(); ++it)
  {
    if (cnt == 1)
    {
      res[it->second] = "Gold Medal";
    }
    else if (cnt == 2)
    {
       res[it->second] = "Silver Medal";
    }
    else if (cnt == 3)
    {
       res[it->second] = "Bronze Medal";
    }
    else
    {
      res[it->second] = to_string(cnt);
    }
    ++cnt;
  }
  return res;
}

int numJewelsInStones(string &J, string &S) // 1038. Jewels And Stones
{
    int res = 0;
    set<char> charSet;
    for (auto cha : J)
    {
        charSet.insert(cha);
    }

    for (auto cha : S)
    {
        if (charSet.count(cha) > 0)
        {
            res += 1;
        }
    }
    return res;
}

int maximumProduct(vector<int> &nums) // 1119. Maximum Product of Three Numbers
{
    int firstNum = MININTNUM;
    int secondNum = MININTNUM;
    int threadNum = MININTNUM;
    int oneNum = MININTNUM;
    int twoNum = MININTNUM;

    for (auto num : nums)
    {
        if (num > firstNum)
        {
            threadNum = secondNum;
            secondNum = firstNum;
            firstNum = num;
        }
        else if (num > secondNum)
        {
            threadNum = secondNum;
            secondNum = num;
        }
        else if (num > threadNum)
        {
            threadNum = num;
        }
    }
    oneNum = firstNum * secondNum * threadNum;
    firstNum = MAXINTNUM;
    secondNum = MAXINTNUM;
    threadNum = MININTNUM;

    for (auto num : nums)
    {
        if (num < firstNum && num < 0)
        {
            secondNum = firstNum;
            firstNum = num;
        }
        else if (num < secondNum &&num < 0)
        {
            secondNum = num;
        }
        else if (num >threadNum && num > 0)
        {
            threadNum = num;
        }
    }
    if (firstNum < 0 && secondNum < 0 && threadNum > 0)
    {
        twoNum = firstNum * secondNum * threadNum;
    }

    return max(oneNum, twoNum);
}

bool partitiontoEqualSumSubsetsSub(vector<int>& nums, int k, int target, int start, int curSum, vector<bool>& visited) {
    if (k == 1)
    {
        return true;
    }
    
    if (curSum == target)
    {
        return partitiontoEqualSumSubsetsSub(nums, k - 1, target, 0, 0, visited);
    }
    
    for (uint i = start; i < nums.size(); ++i)
    {
        if (visited[i])
        {
            continue;
        }
        visited[i] = true;
        
        if (partitiontoEqualSumSubsetsSub(nums, k, target, i + 1, curSum + nums[i], visited))
        {
            return true;
        }
        visited[i] = false;
    }
    return false;
}

bool partitiontoEqualSumSubsets(vector<int> &nums, int k) // 836. Partition to K Equal Sum Subsets
{
    int sum = 0;
    for (auto num : nums)
    {
        sum += num;
    }
    if (sum % k != 0)
    {
        return false;
    }
    
    vector<bool> visited(nums.size(), false);
    return partitiontoEqualSumSubsetsSub(nums, k, sum / k, 0, 0, visited);
}

string reverseString(string &s) // 1283. Reverse String
{
    bool isContinue = true;
    for (uint i = 0; i < s.length() / 2; i++)
    {
        char tmp = s[i];
        s[i] = s[s.length() - 1 - i];
        s[s.length() - 1 - i] = tmp;
        if (i > 0)
        {
            if (s[i] == '\\')
            {
                s[i] = s[i - 1];
                s[i - 1] = '\\';
            }
            if (isContinue && s[s.length() - i] == '\\')
            {
                isContinue = false;
                s[s.length() - i] = s[s.length() - 1 - i];
                s[s.length() - 1 - i] = '\\';
                continue;
            }
            if (!isContinue)
            {
                isContinue = true;
            }
        }
    }
    return s;
}

int titleToNumber(string &s) // 1348. Excel Sheet Column Number
{
    int curDou = 1;
    int res = 0;
    for (uint i = s.length() - 1; (int)i >= 0; i--)
    {
        res += (s[i] - 'A' + 1) * curDou;
        curDou *= 26;
    }
    return res;
}

double largestTriangleArea(vector<vector<int>> &points) // 1005. Largest Triangle Area
{
    double res = 0;
    for (uint i = 0; i < points.size() - 2; i++)
    {
        for (uint j = i + 1; j < points.size() - 1; j++)
        {
            for (uint z = j + 1; z < points.size(); z++)
            {
                double curArea = 0.5 * abs(points[i][0] * points[j][1] - points[i][1] * points[j][0] + points[j][0] * points[z][1] - points[j][1] * points[z][0] + points[z][0] * points[i][1] - points[z][1] * points[i][0]);
                res = max(res, curArea);
            }
        }
    }
    return res;
}

bool multiSortCompare(vector<int> a, vector<int> b)
{
    if (a[1] == b[1])
    {
        return a[0] < b[0];
    }
    return a[1] < b[1];
}

vector<vector<int>> multiSort(vector<vector<int>> &array) // 846. Multi-keyword Sort
{
    //sort(array.begin(), array.end(), multiSortCompare);
    //return array;

    //the sort funcion is Limited by LintCode -- so using bubble sort instead of quick sort
    vector<int> tmpmax;
    for (uint j = 0; j < array.size() - 1; j++)
    {
        for (uint i = array.size() - 1; i != j; i--)
        {
            if (array[i][1]>array[i - 1][1])
            {
                tmpmax = array[i];
                array[i] = array[i - 1];
                array[i - 1] = tmpmax;
            }
            else if (array[i][1] == array[i - 1][1])
            {
                if (array[i][0]<array[i - 1][0])
                {
                    tmpmax = array[i];
                    array[i] = array[i - 1];
                    array[i - 1] = tmpmax;
                }
            }
        }
    }
    return array;
}

long long largestPalindromeSub(int n) {
    string lastHalf = to_string(n);
    reverse(lastHalf.begin(), lastHalf.end());
    return stoll(to_string(n) + lastHalf);
}

int largestPalindrome(int n) { // 1216. Largest Palindrome Product
    if (n == 1)
    {
        return 9;
    }

    int upper = (int)pow(10, n) - 1;
    int lower = (int)pow(10, n - 1);
    for (int i = upper; i > lower; --i) 
    {
        long long cand = largestPalindromeSub(i);
        for (int j = upper; cand / j < j; --j) 
        {
            if (cand % j == 0)
            {
                return cand % 1337;
            }
        }
    }
    return -1;
}

int findLHS(vector<int> &nums){ // 1148. Longest Harmonious Subsequence
    map<int, int> numMap;
    int res = 0;

    for (auto num : nums)
    {
        numMap[num]++;
    }

    if (numMap.size() < 2)
    {
        return 0;
    }

    for (auto it : numMap)
    {
        if (numMap.count(it.first - 1) > 0)
        {
            res = max(res, numMap[it.first - 1] + it.second);
        }
        if (numMap.count(it.first + 1) > 0)
        {
            res = max(res, numMap[it.first + 1] + it.second);
        }
    }

    return res;
}

map<int, int> numMap;

void add(int number)
{
    numMap[number]++;
}

bool find(int value)
{
    for (auto it : numMap)
    {
        if (value - it.first == it.first)
        {
            if (it.second >= 2)
            {
                return true;
            }
        }
        else if (numMap.count(value - it.first) >= 1)
        {
            return true;
        }
    }
    return false;
}

class NestedInteger
{
public:
    bool isInteger() const;
    int getInteger() const;
    const vector<NestedInteger> &getList() const;
};

/* The function using interfacce of type NestedInteger
int depthSum(const vector<NestedInteger>& nestedList) // 551. Nested List Weight Sum
{
    static int dep = 1;
    int sum = 0;

    for (auto it : nestedList)
    {
        if (it.isInteger())
        {
            sum += dep * it.getInteger();
        }
        else
        {
            dep++;
            sum += depthSum(it.getList());
            dep--;
        }
    }
    return sum;
}
*/

vector<int> findRedundantConnection(vector<vector<int>> &edges) // 1088. Redundant Connection
{
    vector<map<int, bool>> relationMapVec;

    for (auto edge : edges)
    {
        bool isFind = false;
        for (uint i = 0; i < relationMapVec.size(); i++)
        {
            if (relationMapVec[i].count(edge[0]) && relationMapVec[i].count(edge[1]))
            {
                return edge;
            }
            if (relationMapVec[i].count(edge[0]) || relationMapVec[i].count(edge[1]))
            {
                relationMapVec[i][edge[0]] = true;
                relationMapVec[i][edge[1]] = true;
                isFind = true;
                for (uint j = i + 1; j < relationMapVec.size(); j++)
                {
                    if (relationMapVec[j].count(edge[0]) || relationMapVec[j].count(edge[1]))
                    {
                        for (auto it : relationMapVec[j])
                        {
                            relationMapVec[i][it.first] = true;
                        }
                        relationMapVec.erase(relationMapVec.begin() + j);
                        break;
                    }
                }
                break;
            }
        }
        if (!isFind)
        {
            map<int, bool> newMap;
            newMap[edge[0]] = true;
            newMap[edge[1]] = true;
            relationMapVec.push_back(newMap);
        }
    }
    return vector<int>();
}

#endif
