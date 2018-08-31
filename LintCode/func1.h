//  func1.h
//  LintCode
//
//  Created by 董旭轩 on 2018/8/21.
//  Copyright © 2018年 董旭轩. All rights reserved.

#ifndef func1.h
#define func1.h

#include "define.h"

int MinimumTotal(vector<vector<int>> &triangle) //  LintCode 109  ˝◊÷»˝Ω«–Œ
{
    int n = triangle.size();
    vector<int> MinNum;
    
    for (int i = 0; i < n; i++)
    {
        MinNum.insert(MinNum.end(), triangle[n - 1][i]);
    }
    
    for (int i = n - 2; i >= 0; i--)
    {
        for (uint j = 0; j < triangle[i].size(); j++)
        {
            MinNum[j] = triangle[i][j] + GetMinValue(MinNum[j], MinNum[j + 1]);
        }
    }
    return MinNum[0];
}

ListNode<int> * DeleteDuplicates(ListNode<int> * head) // LintCode 112 …æ≥˝≈≈–Ú¡¥±Ì÷–µƒ÷ÿ∏¥‘™Àÿ
{
    if (head == NULL || head->next == NULL)
    {
        return head;
    }
    ListNode<int> * last_node = head;
    ListNode<int> * now_node = head->next;
    
    while (now_node != NULL)
    {
        if (last_node->val == now_node->val)
        {
            ListNode<int> * delete_node = now_node;
            now_node = now_node->next;
            last_node->next = now_node;
            delete delete_node;
        }
        else
        {
            now_node = now_node->next;
            last_node = last_node->next;
        }
    }
    return head;
}

int RemoveDuplicates(vector<int> &nums) // LintCode 101 …æ≥˝≈≈–Ú ˝◊È÷–µƒ÷ÿ∏¥ ˝◊÷
{
    if (nums.size() == 0)
    {
        return 0;
    }
    
    int cur_value_num = 1;
    for (uint i = 0; i < nums.size() - 1; i++)
    {
        if (nums[i] == nums[i + 1])
        {
            if (cur_value_num == 2)
            {
                nums.erase(nums.begin() + i);
                i--;
            }
            else
            {
                cur_value_num++;
            }
        }
        else
        {
            cur_value_num = 1;
        }
    }
    return nums.size();
}

template<class val_type> ListNode<val_type> * RemoveNthFromEnd(ListNode<val_type> * head, int n) // …æ≥˝¡¥±Ì÷–µπ ˝µ⁄n∏ˆΩ⁄µ„
{
    if (head == NULL)
    {
        return head;
    }
    
    ListNode<val_type> *early_head = head;
    ListNode<val_type> *delay_head = head;
    ListNode<val_type> *last_head = delay_head;
    
    while (n-- > 0)
    {
        early_head = early_head->next;
    }
    
    while (early_head != NULL)
    {
        last_head = delay_head;
        delay_head = delay_head->next;
        early_head = early_head->next;
    }
    
    if (last_head == delay_head)
    {
        head = head->next;
        delete delay_head;
    }
    else
    {
        last_head->next = delay_head->next;
        delete delay_head;
    }
    
    return head;
}

bool Anagram(string s, string t) // LintCode 158  ¡Ω∏ˆ◊÷∑˚¥Æ «±‰Œª¥
{
    if (s.size() != t.size())
    {
        return false;
    }
    
    int char_num[CHARNUM] = { 0 };
    for (uint i = 0; i < s.size(); i++)
    {
        char_num[s[i]]++;
    }
    
    for (uint i = 0; i < t.size(); i++)
    {
        char_num[t[i]]--;
    }
    
    for (int i = 0; i < CHARNUM; i++)
    {
        if (char_num[i] != 0)
        {
            return false;
        }
    }
    return true;
}

string ReverseWords_error(string &s) // LintCode 53 ∑≠◊™◊÷∑˚¥Æ(error –¥¥Ì≥…∑¥◊™√ø∏ˆµ•¥ )
{
    char last_char = s[1];
    for (int i = s.size() - 1; i >= 0;)
    {
        if (i == s.size() - 1 && s[i] == ' ')
        {
            s.erase(s.begin() + i);
            i--;
        }
        else
        {
            break;
        }
    }
    for (uint i = 0; i < s.size();)
    {
        if (i == 0 && s[i] == ' ')
        {
            s.erase(s.begin() + i);
        }
        else
        {
            break;
        }
    }
    for (uint i = 1; i < s.size();)
    {
        if (s[i] == ' ' && last_char == ' ')
        {
            s.erase(s.begin() + i);
        }
        else
        {
            last_char = s[i];
            i++;
        }
    }
    
    int word_start = 0;
    int word_end = 0;
    for (uint i = 1; i < s.size(); i++)
    {
        if ((i == s.size() - 1) || (s[i] == ' '))
        {
            // do reverse
            if (i == s.size() - 1)
            {
                word_end = i;
            }
            while (word_end > word_start)
            {
                char tmp = s[word_start];
                s[word_start] = s[word_end];
                s[word_end] = tmp;
                word_end--;
                word_start++;
            }
            i++;
            word_start = i;
        }
        word_end = i;
    }
    return s;
}

string ReverseWords(string &s) // LintCode 53 ∑≠◊™◊÷∑˚¥Æ
{
    char last_char = s[0];
    for (int i = s.size() - 1; i >= 0;)
    {
        if (i == s.size() - 1 && s[i] == ' ')
        {
            s.erase(s.begin() + i);
            i--;
        }
        else
        {
            break;
        }
    }
    for (uint i = 0; i < s.size();)
    {
        if (i == 0 && s[i] == ' ')
        {
            s.erase(s.begin() + i);
        }
        else
        {
            break;
        }
    }
    for (uint i = 1; i < s.size();)
    {
        if (s[i] == ' ' && last_char == ' ')
        {
            s.erase(s.begin() + i);
        }
        else
        {
            last_char = s[i];
            i++;
        }
    }
    
    string result;
    int word_start = 0;
    int word_end = 0;
    for (uint i = 1; i < s.size(); i++)
    {
        if ((i == s.size() - 1) || (s[i] == ' '))
        {
            // do reverse
            if (i == s.size() - 1)
            {
                word_end = i;
            }
            
            result.insert(result.begin(), s.begin() + word_start, s.begin() + word_end + 1);
            if (i != s.size() - 1)
            {
                result.insert(0, " ");
            }
            
            i++;
            word_start = i;
        }
        word_end = i;
    }
    return result;
}

int UniquePaths(int m, int n) // LintCode 114 ≤ªÕ¨µƒ¬∑æ∂
{
    int **path_argv = (int **)malloc(sizeof(int*)* m);
    for (int i = 0; i < m; i++)
    {
        path_argv[i] = (int *)malloc(sizeof(int)* n);
    }
    
    for (int i = 0; i < m; i++)
    {
        for (int j = 0; j < n; j++)
        {
            if (i == 0 || j == 0)
            {
                path_argv[i][j] = 1;
            }
            else
            {
                path_argv[i][j] = path_argv[i - 1][j] + path_argv[i][j - 1];
            }
        }
    }
    return path_argv[m - 1][n - 1];
}

vector<int> TwoSum(vector<int> &numbers, int target) // LintCode 56 ¡Ω ˝÷Æ∫Õ
{
    vector<int> result;
    map<int, int> hash_map;
    
    for (uint i = 0; i < numbers.size(); i++)
    {
        hash_map.insert(pair<int, int>(numbers[i], i));
    }
    
    for (uint i = 0; i < numbers.size(); i++)
    {
        int del = target - numbers[i];
        if (hash_map.count(del) != 0 && hash_map[del] != i)
        {
            result.push_back(i);
            result.push_back(hash_map[del]);
            break;
        }
    }
    
    if (result[0] > result[1])
    {
        int tmp = result[0];
        result[0] = result[1];
        result[1] = tmp;
    }
    
    return result;
}

int* KMPGetNext(string base_str)
{
    int *next = new int[base_str.length()];
    next[0] = -1;
    uint j = 0;
    int k = -1;
    while (j < base_str.length() - 1)
    {
        if (k == -1 || base_str[j] == base_str[k])
        {
            next[++j] = ++k;
        }
        else
        {
            k = next[k];
        }
    }
    return next;
}

int KMP(const char *source, const char *target)// KMP À„∑®
{
    if (source == NULL || target == NULL)
    {
        return -1;
    }
    string base_str(source);
    string find_str(target);
    if (base_str.length() == 0)
    {
        if (find_str.length() != 0)
        {
            return -1;
        }
        return 0;
    }
    uint i = 0;
    uint j = 0;
    int *next = KMPGetNext(base_str);
    while (i < base_str.length() && j < (int)find_str.length())
    {
        if (j == -1 || base_str[i] == find_str[j])
        {
            i++; j++;
        }
        else
        {
            j = next[j];
        }
    }
    if (j == find_str.length())
    {
        return i - j;
    }
    else
    {
        return -1;
    }
}

int LongestPalindrome(string s) // LintCode 627 ◊Ó≥§ªÿŒƒ¥Æ
{
    uint result_len = 0;
    bool appear[ENGLISH_CHAR_NUM] = { false };
    for (uint i = 0; i < s.length(); i++)
    {
        char base_cha = 'A';
        char del = s[i] - base_cha;
        if (s[i] <= 'z' && s[i] >= 'a')
        {
            base_cha = 'a';
            del = s[i] - base_cha + ENGLISH_CHAR_NUM / 2;
        }
        if (appear[del])
        {
            appear[del] = false;
            result_len += 2;
        }
        else
        {
            appear[del] = true;
        }
    }
    if (s.length() > result_len)
    {
        result_len++;
    }
    return result_len;
}

string Rearrange(string &str) // LintCode 720. Rearrange a String With Integers
{
    if (str.empty())
    {
        return str;
    }
    int sum = 0;
    for (uint i = 0; i < str.length();)
    {
        if ((str[i] <= 'Z' && str[i] >= 'A') || (str[i] <= 'z' && str[i] >= 'a'))
        {
            i++;
        }
        else
        {
            sum += str[i] - '0';
            str.erase(str.begin() + i);
        }
    }
    char * tmp = (char *)str.c_str();
    int str_len = str.length();
    sort(tmp, tmp + str_len);
    str = str + to_string(sum);
    return str;
}

int Sqrt(uint x) // LintCode 141. xµƒ∆Ω∑Ω∏˘
{
    uint result = 0;
    while (true)
    {
        if (result * result < x)
        {
            result++;
        }
        else
        {
            if (result * result > x)
            {
                return result - 1;
            }
            return result;
        }
    }
}

vector<int> SubarraySum(vector<int> &nums) // LintCode 138. ◊” ˝◊È÷Æ∫Õ
{
    vector<int> result;
    int n = 0;
    if ((n = nums.size()) == 0)
    {
        return result;
    }
    
    map<int, int> num_map;
    num_map[0] = -1;
    int num = 0;
    for (int i = 0; i < n; i++)
    {
        num += nums[i];
        if (num_map.count(num))
        {
            result.push_back(num_map[num] + 1);
            result.push_back(i);
            return result;
        }
        num_map[num] = i;
    }
    return result;
}

void RecoverRotatedSortedArray(vector<int> &nums) // LintCode 39 ª÷∏¥–˝◊™≈≈–Ú ˝◊È
{
    int n = 0;
    int dex = 0;
    int start_dex = 0;
    int end_dex = 0;
    if ((n = nums.size()) == 0)
    {
        return;
    }
    for (int i = 1; i < n; i++)
    {
        if (nums[i] < nums[i - 1])
        {
            dex = i;
            break;
        }
    }
    
    end_dex = dex - 1;
    Recover(nums, start_dex, end_dex);
    
    start_dex = dex;
    end_dex = n - 1;
    Recover(nums, start_dex, end_dex);
    Recover(nums, 0, n - 1);
}

string ConcatenetedString(string &s1, string &s2) // LintCode 702. ¡¨Ω”¡Ω∏ˆ◊÷∑˚¥Æ÷–µƒ≤ªÕ¨◊÷∑˚
{
    vector<int> s1_map[CHARNUM];
    vector<int> s1_deletes;
    for (uint i = 0; i < s1.size(); i++)
    {
        s1_map[s1[i]].push_back(i);
    }
    
    for (uint i = 0; i < s2.size();)
    {
        if (s1_map[s2[i]].size() > 0)
        {
            for (uint j = 0; j < s1_map[s2[i]].size(); j++)
            {
                vector<int> cur_vec = s1_map[s2[i]];
                s1_deletes.push_back(cur_vec[j]);
            }
            s2.erase(s2.begin() + i);
        }
        else
        {
            i++;
        }
    }
    sort(s1_deletes.begin(), s1_deletes.end());
    int last_delete_dex = -1;
    for (int i = s1_deletes.size() - 1; i >= 0; i--)
    {
        if (last_delete_dex != s1_deletes[i])
        {
            s1.erase(s1.begin() + s1_deletes[i]);
            last_delete_dex = s1_deletes[i];
        }
    }
    return s1 + s2;
}

bool IsValidSudoku(vector<vector<char>> &board) // LintCode 389. ≈–∂œ ˝∂¿ «∑Ò∫œ∑®
{
    if (board.size() != 9)
    {
        return false;
    }
    for (uint i = 0; i < board.size(); i++)
    {
        int a[9] = { 0 };
        int b[9] = { 0 };
        if (board[i].size() != 9)
        {
            return false;
        }
        for (uint j = 0; j < board[i].size(); j++)
        {
            if (board[i][j] < '0' && board[i][j] > '9' && board[i][j] != '.')
            {
                return false;
            }
            if (board[i][j] != '.')
            {
                if (a[board[i][j] - '1'] != 0)
                {
                    return false;
                }
                else
                {
                    a[board[i][j] - '1'] = 1;
                }
            }
            if (board[j][i] != '.')
            {
                if (b[board[j][i] - '1'] != 0)
                {
                    return false;
                }
                else
                {
                    b[board[j][i] - '1'] = 1;
                }
            }
        }
    }
    
    uint block_i = 3;
    uint block_j = 3;
    int a[9] = { 0 };
    uint j = 0;
    for (uint i = 0; i < block_i; i++)
    {
        for (; j < block_j; j++)
        {
            if (board[i][j] != '.')
            {
                if (a[board[i][j] - '1'] != 0)
                {
                    return false;
                }
                else
                {
                    a[board[i][j] - '1'] = 1;
                }
            }
        }
        if (j % 3 == 0)
        {
            if ((i + 1) % 3 == 0)
            {
                for (int z = 0; z < 9; z++)
                {
                    a[z] = 0;
                }
                if (block_j + 3 > board[0].size())
                {
                    if (block_i + 3 > board.size())
                    {
                        i = block_i;
                        break;
                    }
                    else
                    {
                        block_i += 3;
                    }
                }
                else
                {
                    block_j += 3;
                    i -= 3;
                }
            }
            else
            {
                j -= 3;
            }
        }
    }
    return true;
}

bool IsSubtree(TreeNode<int> * T1, TreeNode<int> * T2) //LintCode 245. ◊” ˜
{
    if (T1 == NULL)
    {
        if (T2 == NULL)
        {
            return true;
        }
        else
        {
            return false;
        }
    }
    if (T2 == NULL)
    {
        return true;
    }
    
    vector<TreeNode<int> *> NodeList;
    FindNode(T1, T2, NodeList);
    
    for (uint i = 0; i < NodeList.size(); i++)
    {
        TreeNode<int> * T1Node = NodeList[i];
        if (IsEqual(T1Node, T2))
        {
            return true;
        }
    }
    return false;
}

int FindElements(vector<vector<int>> &Matrix) // LintCode 737. Find Elements in Matrix
{
    map<int, int> ElementMap;
    for (uint i = 0; i < Matrix.size(); i++)
    {
        for (uint j = 0; j < Matrix[i].size(); j++)
        {
            if (i == 0)
            {
                ElementMap.insert(pair<int, int>(Matrix[i][j], i));
            }
            else if (ElementMap.count(Matrix[i][j]) && i > 0)
            {
                if (ElementMap[Matrix[i][j]] != (i - 1))
                {
                    continue;
                }
                if (i == Matrix.size() - 1)
                {
                    return Matrix[i][j];
                }
                ElementMap[Matrix[i][j]] = i;
            }
        }
    }
    return -1;
}

bool SearchMatrix(vector<vector<int>> &matrix, int target) // LintCode 28. À—À˜∂˛Œ¨æÿ’Û
{
    if (matrix.size() == 0 || matrix[0].size() == 0)
    {
        return false;
    }
    
    int i_len = matrix.size();
    int j_len = matrix[0].size();
    coordinate<int> start_coor;
    coordinate<int> end_coor;
    end_coor.i = i_len - 1;
    end_coor.j = j_len - 1;
    while (start_coor.i != end_coor.i || start_coor.j != end_coor.j)
    {
        int del_coor = (end_coor.i - start_coor.i) * j_len + end_coor.j - start_coor.j;
        if (del_coor == 1)
        {
            return (matrix[start_coor.i][start_coor.j] == target) || (matrix[end_coor.i][end_coor.j] == target);
        }
        int middle_dex = del_coor / 2;
        int i_middle = middle_dex / j_len;
        int j_middle = middle_dex % j_len;
        coordinate<int> middle_coor;
        middle_coor.i = start_coor.i + i_middle + ((j_middle + start_coor.j >= j_len) ? 1 : 0);
        middle_coor.j = (start_coor.j + j_middle) % j_len;
        if (target < matrix[middle_coor.i][middle_coor.j])
        {
            end_coor.i = middle_coor.i;
            end_coor.j = middle_coor.j;
        }
        else if (target > matrix[middle_coor.i][middle_coor.j])
        {
            start_coor.i = middle_coor.i;
            start_coor.j = middle_coor.j;
        }
        else
        {
            return true;
        }
    }
    
    return matrix[start_coor.i][start_coor.j] == target;
}

int RemoveElement(vector<int> &A, int elem) //LintCode 172. …æ≥˝‘™Àÿ
{
    for (int i = A.size() - 1; i >= 0; i--)
    {
        if (A[i] == elem)
        {
            A.erase(A.begin() + i);
        }
    }
    return A.size();
}

int RemoveDuplicates2(vector<int> &nums) // LintCode 100. …æ≥˝≈≈–Ú ˝◊È÷–µƒ÷ÿ∏¥ ˝◊÷
{
    if (nums.size() == 0)
    {
        return 0;
    }
    
    int last_num = nums[nums.size() - 1];
    for (int i = nums.size() - 2; i >= 0; i--)
    {
        if (last_num == nums[i])
        {
            nums.erase(nums.begin() + i);
        }
        else
        {
            last_num = nums[i];
        }
    }
    return nums.size();
}

bool IsValidParentheses(string &s) //LintCode 423. ”––ßµƒ¿®∫≈–Ú¡–
{
    stack<char> str_sta;
    for (uint i = 0; i < s.size(); i++)
    {
        if (s[i] == '(' || s[i] == '{' || s[i] == '[')
        {
            str_sta.push(s[i]);
        }
        else if (s[i] == ')' || s[i] == '}' || s[i] == ']')
        {
            char str_tmp = str_sta.top();
            str_sta.pop();
            if ((str_tmp == '(' && s[i] == ')') || (str_tmp == '{' && s[i] == '}') || (str_tmp == '[' && s[i] == ']'))
            {
                continue;
            }
            else
            {
                return false;
            }
        }
    }
    return str_sta.size() == 0;
}

void ErgodicTreePaths(TreeNode<int>* root, vector<string> &path_map, vector<int> &cur_path)
{
    if (root == NULL)
    {
        return;
    }
    cur_path.push_back(root->val);
    
    if (root->left == NULL && root->right == NULL)
    {
        string cur = to_string(cur_path[0]);
        for (uint i = 1; i < cur_path.size(); i++)
        {
            cur += string("->");
            cur += to_string(cur_path[i]);
        }
        path_map.push_back(cur);
        cur_path.pop_back();
        return;
    }
    if (root->left != NULL)
    {
        ErgodicTreePaths(root->left, path_map, cur_path);
    }
    if (root->right != NULL)
    {
        ErgodicTreePaths(root->right, path_map, cur_path);
    }
    cur_path.pop_back();
}

vector<string> BinaryTreePaths(TreeNode<int>* root) // LintCode 480. ∂˛≤Ê ˜µƒÀ˘”–¬∑æ∂
{
    vector<string> result;
    vector<int> cur_route;
    if (root == NULL)
    {
        return result;
    }
    
    ErgodicTreePaths(root, result, cur_route);
    return result;
}

bool IsRotateWords(string word1, string word2)
{
    if (word1.length() != word2.length()) //¥À¥¶ π”√size() ‘⁄LintCodeæÕª·±®¥Ì£¨≤¬≤‚ «Stl‘¥¬Îø‚∞Ê±æ≤ªÕ¨µƒŒ Ã‚£¨æª≥∂µ≠~~~ πÌ÷™µ¿ ≤√¥‘≠“Ú°£
    {
        return false;
    }
    word1 = word1 + word1;
    if (word1.find(word2) != string::npos)
    {
        return true;
    }
    else
    {
        return false;
    }
}

int CountRotateWords(vector<string> words) // LintCode 671. —≠ª∑µ•¥
{
    if (words.empty())
    {
        return 0;
    }
    
    int count = words.size();
    bool *flag = new bool[count];
    memset(flag, false, count);
    for (int i = 0; i < count - 1; i++)
    {
        if (flag[i] == true)
        {
            continue;
        }
        for (int j = i + 1; j < count; j++)
        {
            if (IsRotateWords(words[i], words[j]))
            {
                flag[j] = true;
            }
        }
    }
    for (uint i = 0; i < words.size(); i++)
    {
        if (flag[i])
        {
            count--;
        }
    }
    return count;
}

vector<int> PlusOne(vector<int> &digits) // LintCode 407. º”“ª
{
    if (digits.empty())
    {
        return digits;
    }
    
    for (int i = digits.size() - 1; i >= 0; i--)
    {
        if (digits[i] == 9)
        {
            digits[i] = 0;
        }
        else
        {
            digits[i] ++;
            return digits;
        }
    }
    digits.push_back(1);
    for (int i = digits.size() - 1; i > 0; i--)
    {
        digits[i] = digits[i - 1];
    }
    digits[0] = 1;
    return digits;
}

void RotateString(string &str, int offset) //LintCode 8. –˝◊™◊÷∑˚¥Æ
{
    if (str.empty())
    {
        return;
    }
    int length = str.length();
    offset = offset % length;
    str = str + str;
    str.erase(str.begin(), str.begin() + length - offset);
    str.erase(str.begin() + length, str.end());
}

int SearchInsert(vector<int> &A, int target) // LintCode 60. À—À˜≤Â»ÎŒª÷√
{
    if (A.empty())
    {
        return 0;
    }
    int start_dex = 0;
    int end_dex = A.size() - 1;
    while (start_dex <= end_dex)
    {
        int middle_dex = (end_dex - start_dex) / 2 + start_dex;
        if (A[middle_dex] == target)
        {
            return middle_dex;
        }
        else if (A[middle_dex] < target)
        {
            start_dex = middle_dex + 1;
        }
        else
        {
            end_dex = middle_dex - 1;
        }
    }
    return end_dex + 1;
}

ListNode<int> * Reverse(ListNode<int> * head)  // LintCode 35. ∑≠◊™¡¥±Ì
{
    if (head == NULL || head->next == NULL)
    {
        return head;
    }
    ListNode<int> *new_head = head;
    ListNode<int> *last = new_head;
    ListNode<int> *next = head->next;
    last->next = NULL;
    while (next != NULL)
    {
        last = new_head;
        new_head = next;
        next = next->next;
        new_head->next = last;
    }
    return new_head;
}

void SortIntegers2(vector<int> &A) // LintCode 464. ’˚ ˝≈≈–Ú II
{
    sort(A.begin(), A.end());
}

long long TrailingZeros(long long n) // LintCode 2. Œ≤≤øµƒ¡„
{
    long long count = 0;
    while (n>0){
        count += (n / 5);
        n /= 5;
    }
    return count;
}

int SingleNumber(vector<int> &A) // LintCode 82. ¬‰µ•µƒ ˝
{
    int Sum = 0;
    for (uint i = 0; i < A.size(); i++)
    {
        Sum ^= A[i];
    }
    return Sum;
}

void SplitStringSub(string& s, uint pos, vector<string> cur_str, vector<vector<string>> &result)
{
    if (pos >= s.length())
    {
        result.push_back(cur_str);
    }
    
    for (uint i = pos; i < pos + 2 && i < s.size(); ++i)
    {
        vector<string> tmp = cur_str;
        cur_str.push_back(s.substr(pos, i - pos + 1));
        SplitStringSub(s, i + 1, cur_str, result);
        cur_str = tmp;
    }
}

vector<vector<string>> SplitString(string& s) // LintCode 680. ∑÷∏Ó◊÷∑˚¥Æ
{
    vector<vector<string>> result;
    vector<string> cur_str;
    if (s.empty())
    {
        return result;
    }
    
    SplitStringSub(s, 0, cur_str, result);
    return result;
}

int SubSum(int n) // LintCode 730. À˘”–◊”ºØµƒ∫Õ
{
    if (n == 0)
    {
        return 0;
    }
    int result = 0;
    for (int i = n; i > 0; i--)
    {
        result += i;
    }
    return result * (int)pow(2, n - 1);
}

bool IsIsomorphic(string s, string t) // LintCode 638. ◊÷∑˚Õ¨ππ
{
    if (s.length() != t.length())
    {
        return false;
    }
    int len = s.length();
    char substitu_arr[CHARNUM] = { 0 };
    for (int i = 0; i < len; i++)
    {
        if (substitu_arr[s[i]] == 0)
        {
            substitu_arr[s[i]] = t[i];
        }
        else if (substitu_arr[s[i]] != t[i])
        {
            return false;
        }
    }
    bool is_same_arr[CHARNUM] = { false };
    for (int i = 0; i < CHARNUM; i++)
    {
        if (substitu_arr[i] != 0)
        {
            if (is_same_arr[substitu_arr[i]])
            {
                return false;
            }
            is_same_arr[substitu_arr[i]] = true;
        }
    }
    return true;
}

int ReverseInteger(int n) // LintCode 413. ∑¥◊™’˚ ˝
{
    string n_tmp = to_string(n);
    int len = n_tmp.length();
    int j = len - 1;
    int i = 0;
    bool is_del = false;
    if (n < 0)
    {
        i = 1;
        len++;
    }
    for (int i_tmp = i; i_tmp < len / 2; i_tmp++)
    {
        char tmp = n_tmp[i_tmp];
        n_tmp[i_tmp] = n_tmp[j];
        n_tmp[j--] = tmp;
    }
    int result = 0;
    int max_num = (int)MAXINTNUM;
    for (uint i_tmp = i; i_tmp < n_tmp.length(); i_tmp++)
    {
        int cur = (max_num / ((result == 0) ? 1 : result));
        if (cur < 10 && cur > -10)
        {
            return 0;
        }
        result *= 10;
        result += n_tmp[i_tmp] - '0';
    }
    if (n < 0)
    {
        result = -result;
    }
    return result;
}

void UniquePathsWithObstaclesSub(vector<vector<int>> obstacleGrid, uint cur_x, uint cur_y, int &n)
{
    if (cur_x == obstacleGrid.size() - 1 && cur_y == obstacleGrid[0].size() - 1)
    {
        n++;
        return;
    }
    
    if (cur_x < obstacleGrid.size() - 1 && obstacleGrid[cur_x + 1][cur_y] != 1)
    {
        UniquePathsWithObstaclesSub(obstacleGrid, cur_x + 1, cur_y, n);
    }
    if (cur_y < obstacleGrid[0].size() - 1 && obstacleGrid[cur_x][cur_y + 1] != 1)
    {
        UniquePathsWithObstaclesSub(obstacleGrid, cur_x, cur_y + 1, n);
    }
    return;
}

int UniquePathsWithObstaclesError(vector<vector<int>> &obstacleGrid) // LintCode 115. ≤ªÕ¨µƒ¬∑æ∂ II µ›πÈ≥¨ ±
{
    int n = 0;
    if (obstacleGrid.empty() || obstacleGrid[0].empty() || obstacleGrid[0][0] == 1)
    {
        return 0;
    }
    UniquePathsWithObstaclesSub(obstacleGrid, 0, 0, n);
    return n;
}

int UniquePathsWithObstacles(vector<vector<int>> &obstacleGrid) // LintCode 115. ≤ªÕ¨µƒ¬∑æ∂ II ∂ØÃ¨πÊªÆ
{
    int a[100][100] = { 0 };
    int x_len = obstacleGrid.size();
    int y_len = obstacleGrid[0].size();
    
    for (int i = 0; i < x_len; i++)
    {
        if (obstacleGrid[i][0] != 1)
        {
            a[i][0] = 1;
        }
        else
        {
            break;
        }
    }
    for (int i = 0; i < y_len; i++)
    {
        if (obstacleGrid[0][i] != 1)
        {
            a[0][i] = 1;
        }
        else
        {
            break;
        }
    }
    
    for (int i = 1; i < x_len; i++)
    {
        for (int j = 1; j < y_len; j++)
        {
            if (obstacleGrid[i][j] == 1)
            {
                a[i][j] = 0;
            }
            else
            {
                a[i][j] = a[i][j - 1] + a[i - 1][j];
            }
        }
    }
    return a[x_len - 1][y_len - 1];
}

bool CheckSumOfSquareNumbers(int num) // LintCode 697. ≈–∂œ «∑ÒŒ™∆Ω∑Ω ˝÷Æ∫Õ
{
    if (num < 0)
    {
        return false;
    }
    
    int sqr_num = (int)sqrt(num);
    map<int, bool> pow1;
    for (int i = 0; i <= sqr_num; i++)
    {
        pow1.insert(pair<int, bool>(i * i, true));
    }
    for (int i = 0; i <= sqr_num; i++)
    {
        if (pow1.count(num - i * i))
        {
            return true;
        }
    }
    return false;
}

void ConvertBSTSub(TreeNode<int> * root, int &cur_sum)
{
    if (root->right)
    {
        ConvertBSTSub(root->right, cur_sum);
    }
    
    int tmp = cur_sum;
    cur_sum += root->val;
    root->val += tmp;
    
    if (root->left)
    {
        ConvertBSTSub(root->left, cur_sum);
    }
}

TreeNode<int> * ConvertBST(TreeNode<int> * root) // LintCode 661. ∞—∂˛≤ÊÀ—À˜ ˜◊™ªØ≥…∏¸¥Ûµƒ ˜
{
    if (root == NULL)
    {
        return root;
    }
    int cur_sum = 0;
    ConvertBSTSub(root, cur_sum);
    return root;
}

ListNode<int> * Partition(ListNode<int> * head, int x) // LintCode 96. ¡¥±ÌªÆ∑÷
{
    ListNode<int> * small_nodes = NULL;
    ListNode<int> * last_node = NULL;
    ListNode<int> * result = head;
    while (head != NULL)
    {
        if (head->val < x)
        {
            if (small_nodes == NULL)
            {
                small_nodes = head;
                if (small_nodes != result)
                {
                    last_node->next = head->next;
                    small_nodes->next = result;
                    result = small_nodes;
                }
            }
            else
            {
                last_node->next = head->next;
                head->next = small_nodes->next;
                small_nodes->next = head;
                small_nodes = head;
            }
        }
        last_node = head;
        head = head->next;
    }
    return result;
}

bool IsPalindrome(string &s) //LintCode 415. ”––ßªÿŒƒ¥Æ
{
    int j = s.length() - 1;
    int i = 0;
    while (i < j)
    {
        char is_c = s[i];
        char js_c = s[j];
        while ((is_c < 'a' || is_c > 'z') && (is_c < 'A' || is_c > 'Z') && (is_c > '9' || is_c < '0'))
        {
            is_c = s[++i];
        }
        while ((js_c< 'a' || js_c > 'z') && (js_c < 'A' || js_c > 'Z') && (js_c > '9' || js_c < '0'))
        {
            js_c = s[--j];
        }
        if (i > j)
        {
            break;
        }
        if (is_c != js_c)
        {
            if (is_c >= 'A' && is_c <= 'Z')
            {
                is_c += 'a' - 'A';
            }
            if (js_c >= 'A' && js_c <= 'Z')
            {
                js_c += 'a' - 'A';
            }
            if (js_c != is_c)
            {
                return false;
            }
        }
        i++;
        j--;
    }
    return true;
}

int guess(int num, int ans)
{
    if (ans == num)
    {
        return 0;
    }
    if (ans > num)
    {
        return 1;
    }
    else
    {
        return -1;
    }
}

int GuessNumber(int n, int ans) // LintCode 662. ≤¬ ˝”Œœ∑
{
    int min_num = 1;
    int max_num = n;
    while (min_num < max_num)
    {
        int middle = (max_num - min_num) / 2 + min_num;
        int cur_result = guess(middle, ans);
        if (cur_result == 0)
        {
            return middle;
        }
        else if (cur_result == -1)
        {
            min_num = middle + 1;
        }
        else
        {
            max_num = middle - 1;
        }
    }
    return -1;
}

string AddStrings(string &num1, string &num2) // LintCode 655. ¥Û’˚ ˝º”∑®
{
    string s;
    int len1, len2;
    len1 = num1.size() - 1; len2 = num2.size() - 1;
    int i = 0, flag = 0;
    while (len1>-1 && len2>-1){
        int sum = flag + (num1[len1--] - '0') + (num2[len2--] - '0');
        s += char((sum) % 10 + '0');
        flag = sum / 10;
    }
    while (len1>-1){
        int sum = flag + (num1[len1--] - '0');
        s += char((sum) % 10 + '0');
        flag = sum / 10;
    }
    while (len2>-1){
        int sum = flag + (num2[len2--] - '0');
        s += char((sum) % 10 + '0');
        flag = sum / 10;
    }
    if (flag) s += char('0' + flag);
    for (uint i = 0; i<s.size() / 2; i++){
        char c = s[i];
        s[i] = s[s.size() - i - 1];
        s[s.size() - i - 1] = c;
    }
    return s;
}

int ReplaceBlankError(char string[], int length)
{
    int tran_num = 0;
    int char_num = 0;
    bool last_is_space = false;
    for (int i = 0; i < length; i++)
    {
        if (string[i] != ' ')
        {
            char_num++;
            last_is_space = false;
        }
        else if (!last_is_space)
        {
            last_is_space = true;
            tran_num++;
        }
    }
    int new_lenth = char_num + 3 * tran_num;
    int cur_lenth = length;
    string[new_lenth] = 0;
    for (int i = length - 1; i >= 0;)
    {
        if (string[i] == ' ')
        {
            int i_tmp = i - 1;
            int cur_space_num = 1;
            while (i_tmp >= 0 && string[i_tmp] == ' ')
            {
                cur_space_num++;
                i_tmp--;
            }
            if (cur_space_num < 3)
            {
                cur_lenth += 3 - cur_space_num;
                i += 3 - cur_space_num;
                for (int j = cur_lenth - 1; j >= i; j--)
                {
                    string[j] = string[j - (3 - cur_space_num)];
                }
                string[i--] = '0';
                string[i--] = '2';
                string[i--] = '%';
            }
            else
            {
                i_tmp++;
                string[i_tmp++] = '%';
                string[i_tmp++] = '2';
                string[i_tmp++] = '0';
                for (; i_tmp < cur_lenth; i_tmp++)
                {
                    string[i_tmp] = string[i_tmp + (3 - cur_space_num)];
                }
                cur_lenth -= 3 - cur_space_num;
            }
            i = i_tmp - 1;
        }
        else
        {
            i--;
        }
    }
    return new_lenth;
}

int ReplaceBlank(char string[], int length) // LintCode 212. ø’∏ÒÃÊªª
{
    if (string == NULL)
    {
        return 0;
    }
    int tran_num = 0;
    int char_num = 0;
    for (int i = 0; i < length; i++)
    {
        if (string[i] != ' ')
        {
            char_num++;
        }
        else
        {
            tran_num++;
        }
    }
    int new_lenth = char_num + 3 * tran_num;
    string[new_lenth] = 0;
    length--;
    for (int i = new_lenth - 1; i >= 0;)
    {
        if (string[length] == ' ')
        {
            string[i--] = '0';
            string[i--] = '2';
            string[i--] = '%';
        }
        else
        {
            string[i--] = string[length];
        }
        length--;
    }
    return new_lenth;
}

long long PermutationIndex(vector<int> &A) //LintCode 197. ≈≈¡––Ú∫≈
{
    if (A.empty())
    {
        return 0;
    }
    long long result = 1;
    
    for (uint i = 0; i < A.size() - 1; i++)
    {
        int cur_num = A[i];
        int del_num = 0;
        long long cur_result = 1;
        for (uint j = A.size() - 1; j > i; j--)
        {
            if ((A[j] < A[i]) && j > i)
            {
                del_num++;
            }
            if ((j - i) > 0)
            {
                cur_result *= (j - i);
            }
        }
        result += cur_result * del_num;
    }
    return result;
}

bool IsUnique(string &str) // LintCode 157. ≈–∂œ◊÷∑˚¥Æ «∑Ò√ª”–÷ÿ∏¥◊÷∑˚
{
    for (uint i = 0; i < str.length(); i++)
    {
        for (uint j = i + 1; j < str.length(); j++)
        {
            if (str[i] == str[j])
            {
                return false;
            }
        }
    }
    return true;
}

void PartitionArray(vector<int> &nums) // LintCode 373. ∆Ê≈º∑÷∏Ó ˝◊È
{
    int start_dex = 0;
    int end_dex = nums.size() - 1;
    
    while (start_dex < end_dex)
    {
        if ((nums[start_dex] % 2 == 0) && (nums[end_dex] % 2 == 1))
        {
            int tmp = nums[start_dex];
            nums[start_dex++] = nums[end_dex];
            nums[end_dex--] = tmp;
        }
        else if (nums[start_dex] % 2 == 1)
        {
            start_dex++;
        }
        else if (nums[end_dex] % 2 == 0)
        {
            end_dex--;
        }
        else
        {
            start_dex++;
            end_dex--;
        }
    }
}

ListNode<int> * SwapPairs(ListNode<int> * head) // LintCode 451. ¡Ω¡ΩΩªªª¡¥±Ì÷–µƒΩ⁄µ„
{
    if (head == NULL)
    {
        return head;
    }
    
    ListNode<int> * behind_node = NULL;
    ListNode<int> * first_node = head;
    ListNode<int> * second_node = head->next;
    
    while (first_node != NULL && second_node != NULL)
    {
        first_node->next = second_node->next;
        second_node->next = first_node;
        if (behind_node != NULL)
        {
            behind_node->next = second_node;
        }
        else
        {
            head = second_node;
        }
        behind_node = first_node;
        first_node = first_node->next;
        if (first_node != NULL)
        {
            second_node = first_node->next;
        }
    }
    return head;
}

bool IsUgly(int num) // LintCode 517. ≥Û ˝
{
    if (num == 0)
    {
        return false;
    }
    if (num == 1)
    {
        return true;
    }
    
    int t[] = { 2, 3, 5 };
    for (int i = 0; i < 3;)
    {
        int res = num % t[i];
        int cor = num / t[i];
        if (cor == 1 && res == 0)
        {
            return true;
        }
        else if (res == 0)
        {
            num = num / t[i];
        }
        else
        {
            i++;
        }
    }
    
    return false;
}

vector<long long> ProductExcludeItself(vector<int> &nums) // LintCode 50.  ˝◊ÈÃﬁ≥˝‘™Àÿ∫Ûµƒ≥Àª˝
{
    vector<long long> Result;
    for (uint i = 0; i < nums.size(); i++)
    {
        long long cur_sum = 1;
        for (uint j = 0; j < nums.size(); j++)
        {
            if (i != j)
            {
                cur_sum *= nums[j];
            }
        }
        Result.push_back(cur_sum);
    }
    return Result;
}

vector<string> MissingString(string str1, string str2) // LintCode 684. »±…Ÿµƒ◊÷∑˚¥Æ
{
    vector<string> Result;
    int pos_1 = 0, pos_2 = 0;
    int str1_can_end = false, str2_can_end = false;
    
    int find_1 = str1.find(" ", pos_1);
    int find_2 = str2.find(" ", pos_2);
    vector<pair<string, bool>> str1_map;
    while (find_1 != string::npos)
    {
        if (pos_1 != find_1)
        {
            string str1_sub = str1.substr(pos_1, find_1 - pos_1);
            str1_map.push_back(pair<string, bool>(str1_sub, true));
        }
        pos_1 = find_1 + 1;
        find_1 = str1.find(" ", pos_1);
        if (find_1 == string::npos && !str1_can_end)
        {
            str1_can_end = true;
            find_1 = str1.length();
        }
    }
    while (find_2 != string::npos)
    {
        if (pos_2 != find_2)
        {
            string str2_sub = str2.substr(pos_2, find_2 - pos_2);
            for (uint i = 0; i < str1_map.size(); i++)
            {
                if (str1_map[i].first == str2_sub)
                {
                    str1_map[i].second = false;
                }
            }
        }
        pos_2 = find_2 + 1;
        find_2 = str2.find(" ", pos_2);
        if (find_2 == string::npos && !str2_can_end)
        {
            str2_can_end = true;
            find_2 = str2.length();
        }
    }
    for (uint i = 0; i < str1_map.size(); i++)
    {
        if (str1_map[i].second)
        {
            Result.push_back(str1_map[i].first);
        }
    }
    return Result;
}

int FirstUniqChar(string &s) // LintCode 646. µ⁄“ª∏ˆ∂¿Ãÿ◊÷∑˚Œª÷√
{
    int len = s.length();
    map<char, int> mp;
    for (int i = 0; i < len; i++)
    {
        mp[s[i]] ++;
    }
    for (int i = 0; i < len; i++)
    {
        if (mp[s[i]] == 1)
        {
            return i;
        }
    }
    return -1;
}

bool Permutation(string &A, string &B) // LintCode 211. ◊÷∑˚¥Æ÷√ªª /// hahahaha  ∫‹¥Ïµ´ «Œ“œ≤ª∂£¨◊Ó∫Û“ªµ¿ºÚµ•Ã‚£¨œ¬¿¥æÕ «ƒ—µ√¡À£¨À˚¬Ë–¥◊≈ºÚµ•ªπƒ«√¥ƒ—
{
    sort(A.begin(), A.end());
    sort(B.begin(), B.end());
    return A == B;
}

vector<ListNode<int>*> Rehashing(vector<ListNode<int>*> hashTable) //LintCode 129. ÷ÿπ˛œ£
{
    vector<ListNode<int>*> hash_last_node;
    int old_size = hashTable.size();
    int new_size = 2 * old_size;
    for (int i = 0; i < old_size; i++)
    {
        hashTable.push_back(NULL);
    }
    for (uint i = 0; i < hashTable.size(); i++)
    {
        hash_last_node.push_back(hashTable[i]);
    }
    for (int i = 0; i < old_size; i++)
    {
        ListNode<int>* cur_list = hashTable[i];
        while (cur_list != NULL)
        {
            int cur_val = cur_list->val;
            int hash_dex = (cur_val % new_size + new_size) % new_size;
            if (hash_dex != i)
            {
                if (hash_last_node[i] == cur_list)
                {
                    hashTable[i] = cur_list->next;
                    hash_last_node[i] = hashTable[i];
                }
                else
                {
                    hash_last_node[i]->next = cur_list->next;
                }
                ListNode<int>* tmp_cur_list = cur_list;
                cur_list = cur_list->next;
                tmp_cur_list->next = NULL;
                if (hash_last_node[hash_dex] == NULL)
                {
                    hashTable[hash_dex] = tmp_cur_list;
                    hash_last_node[hash_dex] = tmp_cur_list;
                }
                else
                {
                    hash_last_node[hash_dex]->next = tmp_cur_list;
                    hash_last_node[hash_dex] = tmp_cur_list;
                }
            }
            else
            {
                if (hash_last_node[i] != cur_list)
                {
                    hash_last_node[i] = hash_last_node[i]->next;
                }
                cur_list = cur_list->next;
            }
        }
    }
    return hashTable;
}

bool isBadVersion(int n, int fal)
{
    return n >= fal;
}

int FindFirstBadVersion(int n, int fal) // LintCode 74. µ⁄“ª∏ˆ¥ÌŒÛµƒ¥˙¬Î∞Ê±æ
{
    int start_ver = 1;
    int end_ver = n;
    while (start_ver < end_ver)
    {
        int middle = (end_ver - start_ver) / 2 + start_ver;
        if (!isBadVersion(middle, fal))
        {
            start_ver = middle + 1;
        }
        else
        {
            if (middle >= 0 && !isBadVersion(middle - 1, fal))
            {
                return middle;
            }
            else
            {
                end_ver = middle - 1;
            }
        }
    }
    return start_ver;
}

string Serialize(TreeNode<int> * root) // LintCode 7. ∂˛≤Ê ˜µƒ–Ú¡–ªØ∫Õ∑¥–Ú¡–ªØ
{
    static string result;
    if (!root)
    {
        result += string("#,");
    }
    else
    {
        result += to_string(root->val) + string(",");
        Serialize(root->left);
        Serialize(root->right);
    }
    return result;
}

void DeserializeSub(string &data, uint &token, TreeNode<int> * root)
{
    if (data.size() <= token)
    {
        return;
    }
    
    int find_dex = data.find(',', token);
    string cur_val;
    int int_val = 0;
    if (find_dex != string::npos)
    {
        cur_val = data.substr(token, find_dex - token);
        token = find_dex + 1;
        if (cur_val != "#")
        {
            int_val = atoi(cur_val.c_str());
            root->left = new TreeNode<int>(int_val);
            DeserializeSub(data, token, root->left);
        }
    }
    else
    {
        return;
    }
    
    find_dex = data.find(',', token);
    if (find_dex != string::npos)
    {
        cur_val = data.substr(token, find_dex - token);
        token = find_dex + 1;
        if (cur_val != "#")
        {
            int_val = atoi(cur_val.c_str());
            root->right = new TreeNode<int>(int_val);
            DeserializeSub(data, token, root->right);
        }
    }
    else
    {
        return;
    }
}

TreeNode<int> * Deserialize(string &data) // LintCode 7. ∂˛≤Ê ˜µƒ–Ú¡–ªØ∫Õ∑¥–Ú¡–ªØ
{
    if (data.empty())
    {
        return NULL;
    }
    uint token = 0;
    int find_dex = data.find(',', token);
    string cur_val;
    int int_val = 0;
    if (find_dex != string::npos)
    {
        cur_val = data.substr(token, find_dex - token);
        token = find_dex + 1;
        if (cur_val != "#")
        {
            int_val = atoi(cur_val.c_str());
        }
        else
        {
            return NULL;
        }
    }
    else
    {
        return NULL;
    }
    TreeNode<int> * result_root = new TreeNode<int>(int_val);
    DeserializeSub(data, token, result_root);
    return result_root;
}

int FindPeak(vector<int>& A) // LintCode 75. —∞’“∑Â÷µ
{
    int start_dex = 1;
    int end_dex = A.size() - 1;
    while (start_dex < end_dex)
    {
        int middle = (end_dex - start_dex) / 2 + start_dex;
        if (A[middle] > A[middle - 1] && A[middle] > A[middle + 1])
        {
            return middle;
        }
        if (A[middle - 1] > A[middle])
        {
            end_dex = middle - 1;
        }
        else
        {
            start_dex = middle + 1;
        }
    }
    return start_dex;
}

int LongestIncreasingSubsequence(vector<int> &nums) // LintCode 76. ◊Ó≥§…œ…˝◊”–Ú¡–
{
    int size = nums.size();
    int *map = new int[size];
    int result = 0;
    for (int i = 0; i < size; i++)
    {
        map[i] = 1;
    }
    for (int i = 1; i < size; i++)
    {
        for (int j = 0; j < i; j++)
        {
            if (nums[i] > nums[j])
            {
                map[i] = max(map[i], map[j] + 1);
            }
            result = max(map[i], result);
        }
    }
    delete[] map;
    return result;
}

int TreeHigh(TreeNode<int> * root)
{
    if (root->left == NULL && root->right == NULL)
    {
        return 1;
    }
    int left_high = 0;
    int right_high = 0;
    if (root->left != NULL)
    {
        left_high = TreeHigh(root->left);
    }
    if (root->right != NULL)
    {
        right_high = TreeHigh(root->right);
    }
    return max(left_high, right_high) + 1;
}

void RootToString(TreeNode<int> * root, string * root_str, int cur_dex, int node_num)
{
    if (root == NULL || cur_dex > node_num)
    {
        return;
    }
    root_str[cur_dex - 1] = to_string(root->val);
    RootToString(root->left, root_str, 2 * cur_dex, node_num);
    RootToString(root->right, root_str, 2 * cur_dex + 1, node_num);
}

vector<vector<int>> ZigzagLevelOrder(TreeNode<int> * root) // LIntCode 71. ∂˛≤Ê ˜µƒæ‚≥›–Œ≤„¥Œ±È¿˙
{
    vector<vector<int>> result;
    if (root == NULL)
    {
        return result;
    }
    int high = TreeHigh(root);
    int node_num = (int)pow(2, high) - 1;
    string * root_str = new string[node_num];
    for (int i = 0; i < node_num; i++)
    {
        root_str[i] = string("#");
    }
    RootToString(root, root_str, 1, node_num);
    bool IsLeftToRight = true;
    for (int i = 1; i <= high; i++)
    {
        vector<int> cur_hight;
        if (IsLeftToRight)
        {
            for (int j = (int)pow(2, i - 1) - 1; j < (int)pow(2, i) - 1; j++)
            {
                if (root_str[j] != string("#"))
                {
                    cur_hight.push_back(stoi(root_str[j]));
                }
            }
        }
        else
        {
            for (int j = (int)pow(2, i) - 2; j >= (int)pow(2, i - 1) - 1; j--)
            {
                if (root_str[j] != string("#"))
                {
                    cur_hight.push_back(stoi(root_str[j]));
                }
            }
        }
        result.push_back(cur_hight);
        IsLeftToRight = !IsLeftToRight;
    }
    return result;
}

bool LadderLengthSub(string a1, string a2)
{
    int different_num = 0;
    for (uint i = 0; i < a1.length(); i++)
    {
        if (a1[i] != a2[i])
        {
            different_num++;
            if (different_num >= 2)
            {
                return false;
            }
        }
    }
    return different_num == 1;
}

int LadderLengthError(string &start, string &end, unordered_set<string> &dict) // LintCode 120. µ•¥ Ω”¡˙ ≥¨ ±¡À £¨≤‹ƒ·¬Í£°£°£°
{
    int size = dict.size();
    int cur_size = 1;
    map<string, int> last_vec;
    last_vec.insert(pair<string, int>(start, 1));
    map<string, int> visited_map;
    for (int i = 0; i < size; i++)
    {
        map<string, int> cur_vec;
        for (map<string, int>::iterator last_it = last_vec.begin(); last_it != last_vec.end(); last_it++)
        {
            if (LadderLengthSub((*last_it).first, end))
            {
                return cur_size + 1;
            }
            for (unordered_set<string>::iterator it_dic = dict.begin(); it_dic != dict.end(); it_dic++)
            {
                if (LadderLengthSub((*last_it).first, *it_dic)) //
                {
                    if (!cur_vec.count(*it_dic) && !visited_map.count(*it_dic))
                    {
                        visited_map.insert(pair<string, int>(*it_dic, 1));
                        cur_vec.insert(pair<string, int>(*it_dic, 1));
                    }
                }
            }
        }
        cur_size++;
        last_vec = cur_vec;
    }
    return 0;
}

int LadderLength(string &start, string &end, unordered_set<string> &dict) // LintCode 120. µ•¥ Ω”¡˙ Œ“ µ‘⁄≤ª∂Æ’‚÷÷∂Ò–ƒµƒ–¥∑®æ”»ª±»Œ“µƒøÏ
{
    vector<vector<string>> result;
    if (start.size() == 0 || end.size() == 0 || dict.size() == 0) {
        return 0;
    }
    if (start == end)
    {
        return 1;
    }
    if (start.size() == end.size() && start.size() == 1) {
        return 1;
    }
    
    map<string, int> count;
    queue<string> qu;
    qu.push(start);
    dict.erase(start);
    count[start] = 1;
    
    int minLen = MIN_LEN;
    vector<string> curList;
    
    while (!qu.empty() && dict.size() >= 0) {
        string curString = qu.front();
        qu.pop();
        int curLen = count[curString];
        for (uint i = 0; i < curString.size(); ++i) {
            string tmp = curString;
            for (char j = 'a'; j <= 'z'; ++j) {
                if (tmp[i] == j) {
                    continue;
                }
                else {
                    tmp[i] = j;
                    
                    if (dict.find(tmp) != dict.end()) {
                        qu.push(tmp);
                        count[tmp] = curLen + 1;
                        dict.erase(tmp);
                        if (tmp == end) {
                            return count[tmp];
                        }
                    }
                    else if (tmp == end) {
                        count[tmp] = count[curString] + 1;
                        return count[tmp];
                    }
                }
            }
        }
    }
    
    return 0;
}

class SegmentTreeNode {
public:
    int start, end, max, count;
    SegmentTreeNode *left, *right;
    SegmentTreeNode(int start, int end, int count = 0, int max = 0) {
        this->start = start;
        this->end = end;
        this->max = max;
        this->count = count;
        this->left = this->right = NULL;
    }
};

int Query(SegmentTreeNode * root, int start, int end) // LintCode 202. œﬂ∂Œ ˜µƒ≤È—Ø
{
    if (root == NULL)
    {
        return (int)MININTNUM;
    }
    
    if (start == root->start && end == root->end)
    {
        return root->max;
    }
    
    int middle = (root->end - root->start) / 2 + root->start;
    if (start > middle)
    {
        return Query(root->right, start, end);
    }
    else if (end <= middle)
    {
        return Query(root->left, start, end);
    }
    else
    {
        return max(Query(root->left, start, middle), Query(root->right, middle + 1, end));
    }
}

int Query2(SegmentTreeNode * root, int start, int end) // LintCode 247. œﬂ∂Œ ˜≤È—Ø II
{
    if (root == NULL)
    {
        return (int)MININTNUM;
    }
    
    if (start == root->start && end == root->end)
    {
        return root->count;
    }
    if (start < root->start)
    {
        start = root->start;
    }
    if (end > root->end)
    {
        end = root->end;
    }
    
    int middle = (root->end - root->start) / 2 + root->start;
    if (start > middle)
    {
        return Query2(root->right, start, end);
    }
    else if (end <= middle)
    {
        return Query2(root->left, start, end);
    }
    else
    {
        return Query2(root->left, start, middle) + Query2(root->right, middle + 1, end);
    }
}

void Modify(SegmentTreeNode * root, int index, int value) // LintCode 203. œﬂ∂Œ ˜µƒ–ﬁ∏ƒ
{
    if (root == NULL)
    {
        return;
    }
    
    if (index == root->start && index == root->end)
    {
        root->max = value;
        return;
    }
    
    int middle = (root->end - root->start) / 2 + root->start;
    if (index > middle)
    {
        Modify(root->right, index, value);
    }
    else
    {
        Modify(root->left, index, value);
    }
    int cur_max = (int)MININTNUM;
    if (root->left && root->left->max > cur_max)
    {
        cur_max = root->left->max;
    }
    if (root->right && root->right->max > cur_max)
    {
        cur_max = root->right->max;
    }
    root->max = cur_max;
}

SegmentTreeNode * Build(int start, int end) // LintCode 201. œﬂ∂Œ ˜µƒππ‘Ï
{
    SegmentTreeNode * root = new SegmentTreeNode(start, end);
    if (start > end)
    {
        return NULL;
    }
    
    if (start == end)
    {
        return root;
    }
    
    int middle = (end - start) / 2 + start;
    root->left = Build(start, middle);
    root->right = Build(middle + 1, end);
    return root;
}

SegmentTreeNode * Build2Sub(vector<int> &A, int start, int end)
{
    SegmentTreeNode * cur_node = new SegmentTreeNode(start, end, (int)MININTNUM);
    if (start == end)
    {
        cur_node->max = A[start];
        return cur_node;
    }
    int middle = (end - start) / 2 + start;
    cur_node->left = Build2Sub(A, start, middle);
    cur_node->right = Build2Sub(A, middle + 1, end);
    cur_node->max = max(cur_node->left->max, cur_node->right->max);
    return cur_node;
}

SegmentTreeNode * Build2(vector<int> &A) //LintCode 439. œﬂ∂Œ ˜µƒππ‘Ï II
{
    if (A.empty())
    {
        return NULL;
    }
    
    int start = 0;
    int end = A.size() - 1;
    return Build2Sub(A, start, A.size() - 1);
}

struct UndirectedGraphNode {
    int label;
    vector<UndirectedGraphNode *> neighbors;
    UndirectedGraphNode(int x) : label(x) {};
};

UndirectedGraphNode* CloneGraph(UndirectedGraphNode* node) // LintCode 137. øÀ¬°Õº
{
    /*
     if (!node) return NULL;
     
     map<UndirectedGraphNode *, UndirectedGraphNode *> map;
     queue<UndirectedGraphNode *> q;
     q.push(node);
     
     UndirectedGraphNode *graphCopy = new UndirectedGraphNode(node->label);
     map[node] = graphCopy;
     
     while (!q.empty()) {
     UndirectedGraphNode *node1 = q.front();
     q.pop();
     int n = node1->neighbors.size();
     for (int i = 0; i < n; i++) {
     UndirectedGraphNode *neighbor = node1->neighbors[i];
     if (map.find(neighbor) == map.end()) {
     UndirectedGraphNode *p = new UndirectedGraphNode(neighbor->label);
     map[node1]->neighbors.push_back(p);
     map[neighbor] = p;
     q.push(neighbor);
     }
     else {
     map[node1]->neighbors.push_back(map[neighbor]);
     }
     }
     }
     return graphCopy;
     ø…“‘”√µƒ∫Ø ˝£¨≤ª÷™µ¿Œ™√´£¨¿œ±®”Ô∑®¥ÌŒÛ£¨‘Ù”∞œÏ¿œ÷Ω–ƒ«È£¨∆¡±ŒµÙ*/
    return NULL;
}

int PalindromicRanges(int L, int R) // LintCode 745. Palindromic Ranges
{
    int *is_PalArr = new int[R - L + 1];
    int Result = 0;
    for (int i = L; i <= R; i++)
    {
        string tmp = to_string(i);
        if (IsPalindrome(tmp))
        {
            is_PalArr[i - L] = 1;
        }
    }
    for (int start = 0; start <= R - L; start++)
    {
        int cur_sum = is_PalArr[start];
        for (int end = start; end <= R - L; end++)
        {
            if (start != end)
            {
                cur_sum += is_PalArr[end];
            }
            if (cur_sum % 2 == 0)
            {
                Result++;
            }
            
        }
    }
    return Result;
}

void PermuteSub(vector<int> nums, int * is_insert, vector<int> cur_vec, vector<vector<int>> &Result)
{
    int len = nums.size();
    int is_over = true;
    for (int i = 0; i < len; i++)
    {
        if (is_insert[i] == 0)
        {
            is_over = false;
            cur_vec.push_back(nums[i]);
            is_insert[i] = 1;
            PermuteSub(nums, is_insert, cur_vec, Result);
            is_insert[i] = 0;
            cur_vec.pop_back();
        }
    }
    if (is_over)
    {
        Result.push_back(cur_vec);
    }
}

vector<vector<int>> Permute(vector<int> &nums) // LintCode 15. »´≈≈¡–
{
    vector<vector<int>> Result;
    int len = nums.size();
    if (len <= 0)
    {
        Result.push_back(vector<int>{});
        return Result;
    }
    int *is_insert = new int[len]();
    for (int i = 0; i < len; i++)
    {
        vector<int> cur_vec;
        cur_vec.push_back(nums[i]);
        is_insert[i] = 1;
        PermuteSub(nums, is_insert, cur_vec, Result);
        is_insert[i] = 0;
    }
    return Result;
}

int PartitionArray(vector<int> &nums, int k) //LintCode 31.  ˝◊ÈªÆ∑÷
{
    if (nums.empty())
    {
        return 0;
    }
    int start_dex = 0;
    int end_dex = nums.size() - 1;
    while (start_dex < end_dex)
    {
        if (nums[start_dex] < k)
        {
            start_dex++;
        }
        else if (nums[end_dex] >= k)
        {
            end_dex--;
        }
        else if (nums[start_dex] >= k && nums[end_dex] < k)
        {
            int tmp = nums[start_dex];
            nums[start_dex++] = nums[end_dex];
            nums[end_dex--] = tmp;
        }
    }
    
    for (uint i = 0; i < nums.size(); i++)
    {
        if (nums[i] >= k)
        {
            return i;
        }
    }
    return nums.size();
}

vector<int> FindOrder(int numCourses, vector<pair<int, int>> &prerequisites) // LintCode
{
    vector<int> Result;
    if (numCourses == 0)
    {
        return Result;
    }
    
    queue<int> class_que;
    int *order_num = new int[numCourses]();
    vector<int> *order_class = new vector<int>[numCourses];
    for (uint i = 0; i < prerequisites.size(); i++)
    {
        order_class[prerequisites[i].second].push_back(prerequisites[i].first);
        order_num[prerequisites[i].first]++;
    }
    for (int i = 0; i < numCourses; i++)
    {
        if (order_num[i] == 0)
        {
            class_que.push(i);
        }
    }
    while (!class_que.empty())
    {
        int cur_class = class_que.front();
        class_que.pop();
        Result.push_back(cur_class);
        vector<int> cur_order = order_class[cur_class];
        for (uint j = 0; j < cur_order.size(); j++)
        {
            if (order_num[cur_order[j]] == 1)
            {
                order_num[cur_order[j]] = 0;
                class_que.push(cur_order[j]);
            }
            else
            {
                order_num[cur_order[j]]--;
            }
        }
    }
    return Result.size() == numCourses ? Result : vector<int>({});
}

class Tweet { // LintCode 501. √‘ƒ„Õ∆Ãÿ
public:
    int id;
    int user_id;
    string text;
    static Tweet create(int user_id, string tweet_text);
};

Tweet Tweet::create(int user_id, string tweet_text) // LintCode 501. √‘ƒ„Õ∆Ãÿ
{
    Tweet * t = new Tweet;
    t->user_id = user_id;
    t->text = tweet_text;
    return *t;
}

class MiniTwitter { // LintCode 501. √‘ƒ„Õ∆Ãÿ
public:
    vector<Tweet> Tweet_map;
    map<int, map<int, bool>> friend_groups;
    MiniTwitter() {
        // do intialization if necessary
    }
    
    Tweet postTweet(int user_id, string &tweet_text) {
        Tweet new_post = Tweet::create(user_id, tweet_text);
        Tweet_map.push_back(new_post);
        return new_post;
    }
    
    vector<Tweet> getNewsFeed(int user_id) {
        vector<Tweet> Need_Tweet;
        int sum = 0;
        for (int i = Tweet_map.size() - 1; i >= 0; i--)
        {
            if (sum >= 10)
            {
                return Need_Tweet;
            }
            Tweet cur_Tweet = Tweet_map[i];
            if ((friend_groups[user_id].count(cur_Tweet.user_id) && friend_groups[user_id][cur_Tweet.user_id]) || cur_Tweet.user_id == user_id)
            {
                sum++;
                Need_Tweet.push_back(cur_Tweet);
            }
        }
        return Need_Tweet;
    }
    
    vector<Tweet> getTimeline(int user_id) {
        vector<Tweet> Need_Tweet;
        int sum = 0;
        for (int i = Tweet_map.size() - 1; i >= 0; i--)
        {
            if (sum >= 10)
            {
                return Need_Tweet;
            }
            Tweet cur_Tweet = Tweet_map[i];
            if (user_id == cur_Tweet.user_id)
            {
                sum++;
                Need_Tweet.push_back(cur_Tweet);
            }
        }
        return Need_Tweet;
    }
    
    void follow(int from_user_id, int to_user_id) {
        map<int, bool> friend_group = friend_groups[from_user_id];
        if (!friend_group.count(to_user_id) || !friend_group[to_user_id])
        {
            friend_groups[from_user_id].insert(pair<int, bool>(to_user_id, true));
        }
        friend_groups[from_user_id][to_user_id] = true;
    }
    
    void unfollow(int from_user_id, int to_user_id) {
        if (friend_groups[from_user_id].count(to_user_id))
        {
            friend_groups[from_user_id][to_user_id] = false;
        }
    }
};

static bool MinNumberSort(string a, string b)
{
    return a + b < b + a;
}

string MinNumber(vector<int> &nums) //LintCode 379. Ω´ ˝◊È÷ÿ–¬≈≈–Ú“‘ππ‘Ï◊Ó–°÷µ
{
    string result;
    vector<string> str_nums;
    for (uint i = 0; i < nums.size(); i++)
    {
        str_nums.push_back(to_string(nums[i]));
    }
    
    sort(str_nums.begin(), str_nums.end(), MinNumberSort);
    for (uint i = 0; i < str_nums.size(); i++)
    {
        result += str_nums[i];
        nums[i] = stoi(str_nums[i]);
    }
    
    int zero_count = 0;
    for (uint i = 0; i<result.length(); i++)
    {
        if (result[i] == '0'){
            zero_count++;
        }
        else{
            break;
        }
    }
    if (zero_count == result.length()) return "0";
    return result.substr(zero_count, result.length() - zero_count);
}

int CanCompleteCircuit(vector<int> &gas, vector<int> &cost) // LintCode 187. º””Õ’æ
{
    int result = -1;
    int max_abs = 0;
    int cur_abs = 0;
    bool find_pos = false;
    for (uint i = 0; i < gas.size(); i++)
    {
        max_abs += gas[i] - cost[i];
        if (find_pos)
        {
            cur_abs += gas[i] - cost[i];
            if (cur_abs < 0)
            {
                result = -1;
                find_pos = false;
            }
        }
        if (!find_pos && gas[i] > cost[i])
        {
            result = i;
            find_pos = true;
            cur_abs = gas[i] - cost[i];
        }
    }
    return max_abs >= 0 ? result : -1;
}

vector<vector<int>> LevelOrderBottom(TreeNode<int> * root) // LintCode 70. ∂˛≤Ê ˜µƒ≤„¥Œ±È¿˙ II
{
    vector<vector<int>> Result;
    if (root == NULL)
    {
        return Result;
    }
    queue<TreeNode<int> *> root_que;
    stack<vector<int>> root_stack;
    TreeNode<int> * cut_line;
    vector<int> cur_line;
    root_que.push(root);
    root_que.push(cut_line);
    while (!((root_que.size() == 1) && (root_que.front() == cut_line)))
    {
        TreeNode<int> * cur_node = root_que.front();
        root_que.pop();
        if (cur_node == cut_line)
        {
            root_que.push(cut_line);
            root_stack.push(cur_line);
            cur_line = vector<int>();
        }
        else
        {
            if (cur_node->left)
            {
                root_que.push(cur_node->left);
            }
            if (cur_node->right)
            {
                root_que.push(cur_node->right);
            }
            cur_line.push_back(cur_node->val);
        }
    }
    root_stack.push(cur_line);
    while (!root_stack.empty())
    {
        Result.push_back(root_stack.top());
        root_stack.pop();
    }
    return Result;
}

bool LowestCommonAncestorFindNode(TreeNode<int> * root, TreeNode<int> * node, vector<TreeNode<int> *> &prient)
{
    if (root == node)
    {
        prient.push_back(root);
        return true;
    }
    prient.push_back(root);
    if (root->left)
    {
        if (LowestCommonAncestorFindNode(root->left, node, prient))
        {
            return true;
        }
    }
    if (root->right)
    {
        if (LowestCommonAncestorFindNode(root->right, node, prient))
        {
            return true;
        }
    }
    prient.pop_back();
    return false;
}

TreeNode<int> * LowestCommonAncestor(TreeNode<int> * root, TreeNode<int> * A, TreeNode<int> * B) // LintCode 88. ◊ÓΩ¸π´π≤◊Êœ»
{
    if (A == B)
    {
        return A;
    }
    vector<TreeNode<int> *> A_prient;
    vector<TreeNode<int> *> B_prient;
    LowestCommonAncestorFindNode(root, A, A_prient);
    LowestCommonAncestorFindNode(root, B, B_prient);
    int A_len = A_prient.size();
    int B_len = B_prient.size();
    if (A_len < B_len)
    {
        for (int i = 0; i < B_len - A_len; i++)
        {
            B_prient.pop_back();
        }
    }
    else if (B_len < A_len)
    {
        for (int i = 0; i < A_len - B_len; i++)
        {
            A_prient.pop_back();
        }
    }
    while (true)
    {
        if (B_prient.back() == A_prient.back())
        {
            return A_prient.back();
        }
        B_prient.pop_back();
        A_prient.pop_back();
    }
}

class RandomizedSet { //LintCode 657. Insert Delete GetRandom O(1)
public:
    unordered_map<int, int> set_map;
    vector<int> set_vec;
    int size;
    RandomizedSet() {
        size = 0;
    }
    
    bool insert(int val) {
        if (set_map.count(val))
        {
            return false;
        }
        set_map[val] = size++;
        set_vec.push_back(val);
        return true;
    }
    
    bool remove(int val) {
        if (!set_map.count(val))
        {
            return false;
        }
        int vec_dex = set_map[set_map[val]];
        int tmp = set_vec[vec_dex];
        set_vec[vec_dex] = set_vec[size - 1];
        set_vec[size - 1] = tmp;
        set_vec.pop_back();
        set_map.erase(val);
        size--;
        return true;
    }
    
    int getRandom() {
        return set_vec[rand() % size];
    }
};

string IsBuild(int x) // LintCode 749. John's backyard garden
{
    string result = "NO";
    for (int i = 0; i * 3 <= x; i++)
    {
        if ((x - 3 * i) % 7 == 0)
        {
            return "YES";
        }
    }
    return result;
}

bool CanJump(vector<int> &A) // LintCode 116. Ã¯‘æ”Œœ∑
{
    int lenth = A.size() - 1;
    if (lenth == 0 || (lenth == 1 && A[0] > 0))
    {
        return true;
    }
    map<int, int> zero_map;
    for (; lenth >= 0; lenth--)
    {
        if (zero_map.size() != 0)
        {
            map<int, int>::iterator it_map = zero_map.begin();
            if ((*it_map).second <= A[lenth])
            {
                zero_map.erase(it_map);
            }
            else
            {
                ((*it_map).second)++;
            }
        }
        else if (A[lenth] == 0)
        {
            zero_map[lenth] = 2;
        }
    }
    
    return zero_map.size() == 0;
}

int MinCost(vector<vector<int>> &costs) // LintCode 515. ∑øŒ›»æ…´
{
    if (costs.size() == 0 || costs[0].size() != 3)
    {
        return 0;
    }
    
    // 0 ∫Ï…´ 1 ¿∂…´ 2 ¬Ã…´
    int tmp[3] = { 0 };
    int dp[3] = { 0 };
    for (uint i = 0; i < costs.size(); i++)
    {
        tmp[0] = min(dp[1], dp[2]) + costs[i][0];
        tmp[1] = min(dp[0], dp[2]) + costs[i][1];
        tmp[2] = min(dp[0], dp[1]) + costs[i][2];
        dp[0] = tmp[0];
        dp[1] = tmp[1];
        dp[2] = tmp[2];
    }
    return min(min(tmp[0], tmp[1]), tmp[2]);
}

int Portal(vector<vector<char>> &Maze) // LintCode 750. Portal
{
    struct pos
    {
        int x = 0;
        int y = 0;
        int num = 0;
    };
    int i_len = Maze.size();
    int j_len = Maze[0].size();
    pos S_pos;
    for (int i = 0; i < i_len; i++)
    {
        for (int j = 0; j < j_len; j++)
        {
            if (Maze[i][j] == 'S')
            {
                S_pos.x = i;
                S_pos.y = j;
                Maze[i][j] = '#';
                break;
            }
        }
    }
    
    queue<pos> pos_que;
    pos_que.push(S_pos);
    while (!pos_que.empty())
    {
        pos cur_point = pos_que.front();
        if (Maze[cur_point.x][cur_point.y] == 'E')
        {
            return cur_point.num;
        }
        if (cur_point.x + 1 < i_len && Maze[cur_point.x + 1][cur_point.y] != '#')
        {
            pos tmp;
            tmp.x = cur_point.x + 1;
            tmp.y = cur_point.y;
            tmp.num = cur_point.num + 1;
            pos_que.push(tmp);
            if (Maze[cur_point.x + 1][cur_point.y] == '*')
            {
                Maze[cur_point.x + 1][cur_point.y] = '#';
            }
        }
        if (cur_point.y + 1 < j_len && Maze[cur_point.x][cur_point.y + 1] != '#')
        {
            pos tmp;
            tmp.x = cur_point.x;
            tmp.y = cur_point.y + 1;
            tmp.num = cur_point.num + 1;
            pos_que.push(tmp);
            if (Maze[cur_point.x][cur_point.y + 1] == '*')
            {
                Maze[cur_point.x][cur_point.y + 1] = '#';
            }
        }
        if (cur_point.y - 1 >= 0 && Maze[cur_point.x][cur_point.y - 1] != '#')
        {
            pos tmp;
            tmp.x = cur_point.x;
            tmp.y = cur_point.y - 1;
            tmp.num = cur_point.num + 1;
            pos_que.push(tmp);
            if (Maze[cur_point.x][cur_point.y - 1] == '*')
            {
                Maze[cur_point.x][cur_point.y - 1] = '#';
            }
        }
        if (cur_point.x - 1 >= 0 && Maze[cur_point.x - 1][cur_point.y] != '#')
        {
            pos tmp;
            tmp.x = cur_point.x - 1;
            tmp.y = cur_point.y;
            tmp.num = cur_point.num + 1;
            pos_que.push(tmp);
            if (Maze[cur_point.x - 1][cur_point.y] == '*')
            {
                Maze[cur_point.x - 1][cur_point.y] = '#';
            }
        }
        pos_que.pop();
    }
    return -1;
}

struct DirectedGraphNode { // ”–œÚÕºΩ⁄µ„
    int label;
    vector<DirectedGraphNode *> neighbors;
    DirectedGraphNode(int x) : label(x) {};
};

vector<DirectedGraphNode*> TopSort(vector<DirectedGraphNode*>& graph) // LintCode 127. Õÿ∆À≈≈–Ú
{
    vector<DirectedGraphNode*> Result;
    map<DirectedGraphNode*, int> Gramap;
    queue<DirectedGraphNode*> Graque;
    for (uint i = 0; i < graph.size(); i++)
    {
        for (uint j = 0; j < graph[i]->neighbors.size(); j++)
        {
            if (Gramap.count(graph[i]->neighbors[j]))
            {
                Gramap[graph[i]->neighbors[j]] ++;
            }
            else
            {
                Gramap[graph[i]->neighbors[j]] = 1;
            }
        }
    }
    for (uint i = 0; i < graph.size(); i++)
    {
        if (!Gramap.count(graph[i]))
        {
            Graque.push(graph[i]);
        }
    }
    while (!Graque.empty())
    {
        DirectedGraphNode* CurNode = Graque.front();
        Graque.pop();
        Result.push_back(CurNode);
        for (uint i = 0; i < CurNode->neighbors.size(); i++)
        {
            if (!Gramap.count(CurNode->neighbors[i]) || --Gramap[CurNode->neighbors[i]] == 0)
            {
                Graque.push(CurNode->neighbors[i]);
            }
        }
    }
    return Result;
}

string DeleteDigits(string &A, int l) // LintCode 182. …æ≥˝ ˝◊÷
{
    uint dex = 0;
    int i = 0;
    
    for (; dex < A.length() - 1;)
    {
        if (A[dex] > A[dex + 1])
        {
            A = A.substr(0, dex) + A.substr(dex + 1, A.length() - dex - 1);
            if (dex - 1 >= 0)
            {
                dex--;
            }
            if (++i >= l)
            {
                break;
            }
        }
        else
        {
            dex++;
        }
    }
    for (int j = 0; j < l - i; j++)
    {
        int cur_max_dex = 0;
        int cur_max_num = A[0];
        for (uint h = 1; h < A.length(); h++)
        {
            if (cur_max_num < A[h])
            {
                cur_max_num = A[h];
                cur_max_dex = h;
            }
        }
        A = A.substr(0, cur_max_dex) + A.substr(cur_max_dex + 1, A.length() - cur_max_dex - 1);
    }
    while (A[0] == '0')
    {
        A = A.substr(1, A.length() - 1);
    }
    return A;
}

long long GetNumberOfWays(int n, int m, int limit, vector<int> &cost) // LintCode 752. Rogue Knight Sven
{
    vector<map<int, long long>> goldmap;
    map<int, long long> zero_map;
    zero_map[cost[0]] = 1;
    goldmap.push_back(zero_map);
    
    for (int i = 1; i <= n; i++)
    {
        int dex = 1;
        long long new_num = 0;
        map<int, long long> new_map;
        for (int j = i - 1; j >= 0 && dex <= limit; j--, dex++)
        {
            map<int, long long> cur_map = goldmap[j];
            for (map<int, long long>::iterator it_map = cur_map.begin(); it_map != cur_map.end(); it_map++)
            {
                int cur_gold = (*it_map).first;
                if (cur_gold + cost[i] <= m)
                {
                    if (!new_map.count(cur_gold + cost[i]))
                    {
                        new_map[cur_gold + cost[i]] = (*it_map).second;
                    }
                    else
                    {
                        new_map[cur_gold + cost[i]] += (*it_map).second;
                    }
                }
            }
        }
        goldmap.push_back(new_map);
    }
    long long result = 0;
    for (map<int, long long>::iterator it_map = goldmap[n].begin(); it_map != goldmap[n].end(); it_map++)
    {
        result += (*it_map).second;
    }
    return result;
}

bool IsPerfectSquare(int num) // LintCode 777. Valid Perfect Square
{
    double num_sq = sqrt(num);
    if (num_sq - ((int)num_sq) == 0)
    {
        return true;
    }
    return false;
}

vector<int> SearchRange(TreeNode<int> * root, int k1, int k2)
{
    vector<int> result;
    if (root == NULL)
    {
        return result;
    }
    struct stack_node
    {
        TreeNode<int> * cur_node;
        bool left_node_ok = false;
        bool right_node_ok = false;
    };
    stack<stack_node> node_stack;
    stack_node new_node;
    new_node.cur_node = root;
    node_stack.push(new_node);
    while (!node_stack.empty())
    {
        stack_node node = node_stack.top();
        if (node.cur_node->left == NULL || node.left_node_ok)
        {
            node_stack.pop();
            if (node.cur_node->right != NULL && !node.right_node_ok)
            {
                stack_node right_node;
                right_node.cur_node = node.cur_node->right;
                node_stack.push(right_node);
            }
            if (node.cur_node->val >= k1 && node.cur_node->val <= k2)
            {
                result.push_back(node.cur_node->val);
            }
            else if (node.cur_node->val > k2)
            {
                return result;
            }
        }
        else if (!node.left_node_ok)
        {
            node_stack.pop();
            node.left_node_ok = true;
            node_stack.push(node);
            stack_node left_node;
            left_node.cur_node = node.cur_node->left;
            node_stack.push(left_node);
        }
    }
    return result;
}

vector<int> Business(vector<int> &A, uint k) // LintCode 751. ‘º∫≤µƒ…˙“‚
{
    map<int, int> busmap;
    vector<int> result;
    for (uint i = 0; i < k; i++)
    {
        if (i < A.size())
        {
            if (!busmap.count(A[i]))
            {
                busmap[A[i]] = 1;
            }
            else
            {
                busmap[A[i]] ++;
            }
        }
    }
    for (uint i = 0; i < A.size(); i++)
    {
        if (i + k < A.size())
        {
            if (!busmap.count(A[i + k]))
            {
                busmap[A[i + k]] = 1;
            }
            else
            {
                busmap[A[i + k]] ++;
            }
        }
        if (i - k > 0)
        {
            if (--busmap[A[i - k - 1]] == 0)
            {
                busmap.erase(A[i - k - 1]);
            }
        }
        int tmp = (*busmap.begin()).first;
        result.push_back(A[i] - (*busmap.begin()).first);
    }
    return result;
}

vector<vector<int>> ThreeSum(vector<int> &numbers) // LintCode 57. »˝ ˝÷Æ∫Õ
{
    vector<vector<int>> result;
    sort(numbers.begin(), numbers.end());
    for (uint i = 0; i < numbers.size() - 2; i++)
    {
        if (i > 0 && numbers[i] == numbers[i - 1])
        {
            continue;
        }
        for (uint j = i + 1; j < numbers.size() - 1; j++)
        {
            if (j > i + 1 && numbers[j] == numbers[j - 1])
            {
                continue;
            }
            for (uint z = j + 1; z < numbers.size(); z++)
            {
                if (z > j + 1 && numbers[z] == numbers[z - 1])
                {
                    continue;
                }
                if (numbers[i] + numbers[j] + numbers[z] == 0)
                {
                    vector<int> cur_vec({ numbers[i], numbers[j], numbers[z] });
                    result.push_back(cur_vec);
                }
            }
        }
    }
    return result;
}

int MonotoneDigits(int num) // LintCode 743. µ•µ˜µ›‘ˆµƒ ˝◊÷
{
    string num_str = to_string(num);
    for (uint i = 0; i < num_str.length() - 1; i++)
    {
        if (num_str[i] > num_str[i + 1])
        {
            while (i - 1 >= 0 && num_str[i] - 1 < num_str[i - 1])
            {
                i--;
            }
            num_str[i] = num_str[i] - 1;
            for (uint j = i + 1; j < num_str.length(); j++)
            {
                num_str[j] = '9';
            }
            return stoi(num_str);
        }
    }
    return num;
}

void SortColors(vector<int> &nums) // LintCode 148. —’…´∑÷¿‡
{
    int zero_dex = 0;
    int two_dex = nums.size() - 1;
    while (nums[zero_dex] == 0 && zero_dex <= two_dex){ zero_dex++; }
    while (nums[two_dex] == 2 && two_dex >= zero_dex){ two_dex--; }
    int one_dex = zero_dex;
    while (one_dex <= two_dex)
    {
        if (nums[one_dex] == 0)
        {
            if (nums[zero_dex] == 2)
            {
                nums[two_dex--] = 2;
            }
            nums[one_dex] = 1;
            nums[zero_dex++] = 0;
        }
        else if (nums[one_dex] == 2)
        {
            if (nums[two_dex] == 0)
            {
                nums[zero_dex++] = 0;
            }
            if (one_dex >= zero_dex)
            {
                nums[one_dex] = 1;
            }
            nums[two_dex--] = 2;
        }
        one_dex++;
        while (nums[two_dex] == 2 && two_dex >= one_dex){ two_dex--; }
        for (uint i = 0; i < nums.size(); i++)
        {
            cout << nums[i] << " ";
        }
        cout << endl;
    }
}

/*  »•ƒ„¬ËµƒSBÕÊ“‚£¨Ã‚ƒøø¥≤ª∂Æ£¨Ã‚“‚≤ª√˜»∑£¨≤ª∏¯÷¥––√¸¡Ó£¨Õ¯…œªπ√ª”–œ‡”¶Ω‚ Õ£¨ÕÊƒ„√√∞°£¨–“∫√¿œ◊””–“„¡¶£¨—–æø¡À“ªœ¬map reduce‘¥¬Î
 template<class T> class Input { // LintCode 549. ◊Ó≥£ π”√µƒk∏ˆµ•¥ (Map Reduce)
 public:
 bool done(); // Returns true if the iteration has elements or false.
 void next(); // Move to the next element in the iteration ,Runtime error if the iteration has no more elements
 T value(); // Get the current element, Runtime error if the iteration has no more elements
 };
 
 class Document { // LintCode 549. ◊Ó≥£ π”√µƒk∏ˆµ•¥ (Map Reduce)
 public:
 int id; // document id
 string content; // document content
 };
 
 class TopKFrequentWordsMapper { // LintCode 549. ◊Ó≥£ π”√µƒk∏ˆµ•¥ (Map Reduce)
 public:
 void Map(Input<Document>* input) {
 while (!input->done())
 {
 string str = input->value().content;
 int j = 0;
 for (int i = 0; i <= str.length(); i++)
 {
 if (str[i] == ' ' || i == str.length())
 {
 string tmp = str.substr(j, i - j);
 if (tmp != "")
 {
 output(tmp, 1);
 }
 j = i + 1;
 }
 }
 input->next();
 }
 }
 };
 
 class TopKFrequentWordsReducer{ // LintCode 549. ◊Ó≥£ π”√µƒk∏ˆµ•¥ (Map Reduce)
 public:
 map<string, int> str_map;
 int num = 0;
 struct map_cmp
 {
 bool operator()(const pair<string, int>& lhs, const pair<string, int>& rhs) {
 if (lhs.second == rhs.second)
 {
 return lhs.first < rhs.first;
 }
 return lhs.second > rhs.second;
 }
 };
 
 void setUp(int k) {
 // initialize your data structure here
 num = k;
 }
 
 void Reduce(string &key, Input<int>* input) {
 int sum = 0;
 while (!input->done())
 {
 sum += input->value();
 input->next();
 }
 str_map[key] = sum;
 }
 
 void cleanUp() {
 vector<pair<string, int>> map_vec(str_map.begin(), str_map.end());
 sort(map_vec.begin(), map_vec.end(), map_cmp());
 for (int i = 0; i < num && i < map_vec.size(); i++)
 {
 output(map_vec[i].first, map_vec[i].second);
 }
 }
 };
 */

long long GetSum(int n, vector<int> &nums) //LintCode 782. ”ÎªÚ∫Õ
{
    sort(nums.begin(), nums.end());
    long long min_or = nums[0];
    long long max_and = nums[nums.size() - 1];
    long long max_or = min_or;
    long long min_and = min_or;
    for (uint i = 1; i < nums.size(); i++)
    {
        max_or |= nums[i];
        min_and &= nums[i];
    }
    return min_or + max_and + max_or + min_and;
}

int HouseRobber2(vector<int> nums) // LintCode 534. ¥ÚΩŸ∑øŒ› II
{
    if (nums.empty())
    {
        return 0;
    }
    if (nums.size() == 1)
    {
        return nums[0];
    }
    if (nums.size() == 2)
    {
        return max(nums[0], nums[1]);
    }
    
    int *a = new int[nums.size()];
    int have_once = 0, no_once = 0;
    a[0] = nums[0];
    a[1] = nums[0];
    for (uint i = 2; i < nums.size(); i++)
    {
        if (i == nums.size() - 1)
        {
            break;
        }
        a[i] = max(a[i - 2] + nums[i], a[i - 1]);
    }
    have_once = a[nums.size() - 2];
    a[0] = 0;
    a[1] = nums[1];
    for (uint i = 2; i < nums.size(); i++)
    {
        a[i] = max(a[i - 2] + nums[i], a[i - 1]);
    }
    no_once = a[nums.size() - 1];
    return max(have_once, no_once);
}

int ValidTreeSub(vector<int> &existed, int e)
{
    while (existed[e] != -1)
    {
        e = existed[e];
    }
    return e;
}

bool ValidTree(int n, vector<vector<int>> &edges) // LintCode 178. Õº «∑Ò « ˜
{
    if (n - edges.size() != 1)
    {
        return false;
    }
    
    vector<int> existed(n, -1);
    for (uint i = 0; i < edges.size(); ++i)
    {
        int root1 = ValidTreeSub(existed, edges[i][0]);
        int root2 = ValidTreeSub(existed, edges[i][1]);
        if (root1 == root2)
        {
            return false;
        }
        
        existed[root2] = root1;
    }
    return true;
}

int TheLongestCommonPrefix(vector<string> &dic, string &target) // LintCode 784. ◊Ó≥§π´π≤«∞◊∫ II
{
    int num = 0;
    for (uint i = 0; i < dic.size(); i++)
    {
        uint s = 0;
        for (; s < target.length(); s++)
        {
            if (target[s] != dic[i][s])
            {
                break;
            }
        }
        num = max(num, (int)s);
    }
    return num;
}

int KthPrime(int n) // LintCode 792. µ⁄K∏ˆ÷  ˝
{
    int num = 1;
    for (int i = 2; i < n; i++)
    {
        int sqrt_i = (int)sqrt(i);
        bool is_true = true;
        for (int j = 2; j <= sqrt_i; j++)
        {
            double dec = (double)i / (double)j;
            if (dec - (int)dec == 0)
            {
                is_true = false;
                break;
            }
        }
        if (is_true)
        {
            num++;
        }
    }
    return num;
}

int MaxTwoSubArrays(vector<int> &nums) // LintCode 42. ◊Ó¥Û◊” ˝◊È II
{
    if (nums.empty())
    {
        return 0;
    }
    int *left_max_ar = new int[nums.size()];
    int *right_max_ar = new int[nums.size()];
    int left_cur = 0;
    int right_cur = 0;
    int left_max = (int)MININTNUM;
    int right_max = (int)MININTNUM;
    int result = (int)MININTNUM;
    
    for (uint i = 0; i < nums.size(); i++)
    {
        if (left_cur <= 0)
        {
            left_cur = nums[i];
        }
        else
        {
            left_cur += nums[i];
        }
        left_max = max(left_max, left_cur);
        left_max_ar[i] = left_max;
        
        if (right_cur <= 0)
        {
            right_cur = nums[nums.size() - 1 - i];
        }
        else
        {
            right_cur += nums[nums.size() - 1 - i];
        }
        right_max = max(right_max, right_cur);
        right_max_ar[nums.size() - 1 - i] = right_max;
    }
    
    for (uint i = 0; i < nums.size() - 1; i++)
    {
        result = max(result, left_max_ar[i] + right_max_ar[i + 1]);
    }
    return result;
}

int MergeNumber(vector<int> &numbers) // LintCode 791. ∫œ≤¢ ˝◊÷
{
    int Mernum = 0;
    map<int, int> num_map;
    for (uint i = 0; i < numbers.size(); i++)
    {
        num_map[numbers[i]] += 1;
    }
    
    while (num_map.size() >= 1)
    {
        map<int, int>::iterator it_map = num_map.begin();
        int cur_merge;
        if (num_map.size() == 1 && (*it_map).second == 1)
        {
            break;
        }
        
        if ((*it_map).second == 1)
        {
            int first_num = (*it_map).first;
            int second_num = (*++it_map).first;
            cur_merge = first_num + second_num;
            num_map.erase(first_num);
            num_map[second_num] --;
            if (num_map[second_num] == 0)
            {
                num_map.erase(second_num);
            }
        }
        else
        {
            int first_num = (*it_map).first;
            cur_merge = first_num * 2;
            num_map[first_num] -= 2;
            if (num_map[first_num] == 0)
            {
                num_map.erase(first_num);
            }
        }
        Mernum += cur_merge;
        num_map[cur_merge] ++;
    }
    return Mernum;
}

int IntersectionOfArrays(vector<vector<int>> &arrs) // LintCode 793.  ˝◊ÈµƒΩªºØ
{
    if (arrs.empty())
    {
        return 0;
    }
    
    map<int, bool> Ar_map;
    map<int, bool> Temp_map;
    for (uint i = 0; i < arrs[0].size(); i++)
    {
        Ar_map[arrs[0][i]] = true;
        Temp_map[arrs[0][i]] = true;
    }
    
    for (uint i = 0; i < arrs.size(); i++)
    {
        Ar_map.clear();
        for (uint j = 0; j < arrs[i].size(); j++)
        {
            if (Temp_map.count(arrs[i][j]))
            {
                Ar_map[arrs[i][j]] = true;
            }
        }
        Temp_map = Ar_map;
    }
    return Ar_map.size();
}

bool CanBeGeneratedSub(vector<string> char_arrs[], string str1, string str2, int dex) // LintCode 790.∑˚∫≈¥Æ…˙≥…∆˜
{
    for (uint i = dex; i < str2.size() && i < str1.size();)
    {
        if (str1[i] == str2[i])
        {
            i++;
        }
        else if ('a' <= str1[i] && str1[i] <= 'z')
        {
            return false;
        }
        else
        {
            vector<string> cur_vec = char_arrs[str1[i] - 'A'];
            for (uint j = 0; j < cur_vec.size(); j++)
            {
                if (CanBeGeneratedSub(char_arrs, str1.substr(0, i) + cur_vec[j] + str1.substr(i + 1, str1.length() - i - 1), str2, i))
                {
                    return true;
                }
            }
            return false;
        }
    }
    return str1.length() == str2.length();
}

bool CanBeGenerated(vector<string> &generator, char startSymbol, string symbolString) // LintCode 790.∑˚∫≈¥Æ…˙≥…∆˜
{
    vector<string> char_arrs[BIG_ENGLISH_CHAR_NUM];
    for (uint i = 0; i < generator.size(); i++)
    {
        string cur_str = generator[i];
        char_arrs[cur_str[0] - 'A'].push_back(cur_str.substr(5));
    }
    
    string init_str;
    init_str.push_back(startSymbol);
    return CanBeGeneratedSub(char_arrs, init_str, symbolString, 0);
}

bool ExistSub(vector<vector<char>> &board, vector<vector<bool>> &sign_load, string &word, int str_dex, uint i, uint j) //LintCode 123. µ•¥ À—À˜
{
    if (board[i][j] == word[str_dex])
    {
        if (str_dex == word.length() - 1)
        {
            return true;
        }
        
        sign_load[i][j] = false;
        if (i - 1 >= 0 && sign_load[i - 1][j] && ExistSub(board, sign_load, word, str_dex + 1, i - 1, j))
        {
            return true;
        }
        if (i + 1 < sign_load.size() && sign_load[i + 1][j] && ExistSub(board, sign_load, word, str_dex + 1, i + 1, j))
        {
            return true;
        }
        if (j - 1 >= 0 && sign_load[i][j - 1] && ExistSub(board, sign_load, word, str_dex + 1, i, j - 1))
        {
            return true;
        }
        if (j + 1 < sign_load[i].size() && sign_load[i][j + 1] && ExistSub(board, sign_load, word, str_dex + 1, i, j + 1))
        {
            return true;
        }
        sign_load[i][j] = true;
        return false;
    }
    else
    {
        return false;
    }
}

bool Exist(vector<vector<char>> &board, string &word) //LintCode 123. µ•¥ À—À˜
{
    if (word.length() == 0)
    {
        return true;
    }
    if (board.empty())
    {
        return false;
    }
    
    
    vector<vector<bool>> sign_load;
    for (uint i = 0; i < board.size(); i++)
    {
        vector<bool> cur_load;
        for (uint j = 0; j < board[i].size(); j++)
        {
            cur_load.push_back(true);
        }
        sign_load.push_back(cur_load);
    }
    
    for (uint i = 0; i < board.size(); i++)
    {
        for (uint j = 0; j < board[i].size(); j++)
        {
            if (board[i][j] == word[0] && ExistSub(board, sign_load, word, 0, i, j))
            {
                return true;
            }
        }
    }
    return false;
}

/* ”÷“ª∏ˆÃ‚“‚Œ“œÎ≤‹ƒ·¬Íœµ¡–
 class Comparator // LintCode 399. Nuts ∫Õ Bolts µƒŒ Ã‚
 {
 public:
 int cmp(string a, string b)
 {
 return a > b ? 1 : -1;
 }
 };
 
 void SortNutsAndBolts(vector<string> &nuts, vector<string> &bolts, Comparator compare) // LintCode 399. Nuts ∫Õ Bolts µƒŒ Ã‚
 {
 for (int i = 0; i<nuts.size(); i++){
 for (int j = i; j<bolts.size(); j++){
 int record = compare.cmp(nuts[i], bolts[j]);
 if (record == 0){
 string tmp = bolts[i];
 bolts[i] = bolts[j];
 bolts[j] = tmp;
 break;
 }
 }
 }
 }*/

vector<vector<string>> GroupAnagrams(vector<string> &strs) //LintCode 772. ¥ÌŒª¥ ∑÷◊È -- ”÷ «“ªµ¿ø”µ˘Ã‚£¨–¬Ã‚∂º «ø”µ˘µƒ
{
    struct str_node
    {
        int char_arr[BIG_ENGLISH_CHAR_NUM];
        int size = 0;
    };
    vector<vector<string>> Result;
    vector<str_node> node_arr;
    for (uint i = 0; i < strs.size(); i++)
    {
        string cur_str = strs[i];
        int cur_char_arr[BIG_ENGLISH_CHAR_NUM] = { 0 };
        int cur_size = cur_str.length();
        bool is_find = false;
        for (uint j = 0; j < cur_str.length(); j++)
        {
            cur_char_arr[cur_str[j] - 'a'] ++;
        }
        for (uint j = 0; j < node_arr.size(); j++)
        {
            if (node_arr[j].size == cur_str.length())
            {
                int z = 0;
                for (; z < BIG_ENGLISH_CHAR_NUM; z++)
                {
                    if (node_arr[j].char_arr[z] != cur_char_arr[z])
                    {
                        break;
                    }
                }
                if (z == BIG_ENGLISH_CHAR_NUM)
                {
                    Result[j].push_back(cur_str);
                    is_find = true;
                    break;
                }
            }
        }
        if (!is_find)
        {
            str_node new_node;
            vector<string> new_vec;
            for (int j = 0; j < BIG_ENGLISH_CHAR_NUM; j++)
            {
                new_node.char_arr[j] = cur_char_arr[j];
            }
            new_node.size = cur_size;
            node_arr.push_back(new_node);
            new_vec.push_back(strs[i]);
            Result.push_back(new_vec);
        }
    }
    return Result;
}

bool PacificAtlanticSub(vector<vector<int>> &matrix, vector<vector<int>> &sign_can_move, bool** sign_moved, uint i, uint j) // LintCode 778. Pacific Atlantic Water Flow
{
    if (sign_can_move[i][j] == 3)
    {
        return true;
    }
    else if (sign_can_move[i][j] == -1)
    {
        return false;
    }
    else if (sign_can_move[i][j] != 0)
    {
        return false;
    }
    
    bool need_find = true;
    sign_moved[i][j] = true;
    if (i == 0 || j == 0)
    {
        sign_can_move[i][j] |= 1;
    }
    if (i == matrix.size() - 1 || j == matrix[i].size() - 1)
    {
        sign_can_move[i][j] |= 2;
    }
    if (i + 1 < matrix.size() && need_find && !sign_moved[i + 1][j] && matrix[i + 1][j] <= matrix[i][j])
    {
        PacificAtlanticSub(matrix, sign_can_move, sign_moved, i + 1, j);
        if (sign_can_move[i + 1][j] != -1)
        {
            sign_can_move[i][j] |= sign_can_move[i + 1][j];
        }
    }
    if (i - 1 >= 0 && need_find && !sign_moved[i - 1][j] && matrix[i - 1][j] <= matrix[i][j])
    {
        PacificAtlanticSub(matrix, sign_can_move, sign_moved, i - 1, j);
        if (sign_can_move[i - 1][j] != -1)
        {
            sign_can_move[i][j] |= sign_can_move[i - 1][j];
        }
    }
    if (j + 1 < matrix[0].size() && need_find && !sign_moved[i][j + 1] && matrix[i][j + 1] <= matrix[i][j])
    {
        PacificAtlanticSub(matrix, sign_can_move, sign_moved, i, j + 1);
        if (sign_can_move[i][j + 1] != -1)
        {
            sign_can_move[i][j] |= sign_can_move[i][j + 1];
        }
    }
    if (j - 1 >= 0 && need_find && !sign_moved[i][j - 1] && matrix[i][j - 1] <= matrix[i][j])
    {
        PacificAtlanticSub(matrix, sign_can_move, sign_moved, i, j - 1);
        if (sign_can_move[i][j - 1] != -1)
        {
            sign_can_move[i][j] |= sign_can_move[i][j - 1];
        }
    }
    
    sign_moved[i][j] = false;
    if (sign_can_move[i][j] == 3)
    {
        return true;
    }
    if (sign_can_move[i][j] == 0)
    {
        sign_can_move[i][j] = -1;
    }
    return false;
}

vector<vector<int>> PacificAtlantic(vector<vector<int>> &matrix) // LintCode 778. Pacific Atlantic Water Flow –°«≈¡˜ÀÆ»Àº“-°∑¬Ì…œ∑≈ºŸ£¨æÕø™–ƒ
{
    vector<vector<int>> sign_can_move;
    bool** sign_moved;
    vector<vector<int>> result;
    
    sign_moved = new bool*[matrix.size()];
    
    for (uint i = 0; i < matrix.size(); i++)
    {
        sign_moved[i] = new bool[matrix[i].size()];
        vector<int> cur_sign_vec;
        for (uint j = 0; j < matrix[i].size(); j++)
        {
            sign_moved[i][j] = false;
            cur_sign_vec.push_back(0);
        }
        sign_can_move.push_back(cur_sign_vec);
    }
    
    for (uint i = 0; i < matrix.size(); i++)
    {
        for (uint j = 0; j < matrix[i].size(); j++)
        {
            if (PacificAtlanticSub(matrix, sign_can_move, sign_moved, i, j))
            {
                result.push_back(vector<int>({ (int)i, (int)j }));
            }
        }
    }
    return result;
}

bool IsPalindrome2(int n) //LintCode 807. ªÿŒƒ ˝ II
{
    int dex_lenth = 0;
    for (int i = 31; i >= 0; i--)
    {
        if ((n & (1 << i)) > 0)
        {
            dex_lenth = i + 1;
            break;
        }
    }
    for (int i = 0; i < dex_lenth / 2; i++)
    {
        if ((n >> i & 1) != (n >> (dex_lenth - i - 1) & 1))
        {
            return false;
        }
    }
    return true;
}

vector<int> WinSum(vector<int> &nums, uint k) // LintCode 604. ª¨∂Ø¥∞ø⁄ƒ⁄ ˝µƒ∫Õ
{
    vector<int> Result;
    if (k == 0)
    {
        return Result;
    }
    int cur_num = 0;
    uint i = 0;
    for (; i < k && i < nums.size(); i++)
    {
        cur_num += nums[i];
    }
    Result.push_back(cur_num);
    for (; i < nums.size(); i++)
    {
        cur_num -= nums[i - k];
        cur_num += nums[i];
        Result.push_back(cur_num);
    }
    return Result;
}

void TopKMovieSub(bool * relation_arr, vector<vector<int>> &G, int i, int S)
{
    relation_arr[i] = true;
    vector<int> cur_vec = G[i];
    for (uint j = 0; j < cur_vec.size(); j++)
    {
        if (cur_vec[j] != S && !relation_arr[cur_vec[j]])
        {
            TopKMovieSub(relation_arr, G, cur_vec[j], S);
        }
    }
}

vector<int> TopKMovie(vector<int> &rating, vector<vector<int>> &G, int S, uint K) // LintCode 808. ”∞º Õ¯¬Á
{
    struct CmpByValue {
        bool operator()(const pair<int, int>& k1, const pair<int, int>& k2) {
            return k1.second > k2.second;
        }
    };
    
    vector<int> result;
    uint n = G.size();
    bool *relation_arr = new bool[n];
    vector<pair<int, int>> relation_vec;
    
    for (uint i = 0; i < n; i++)
    {
        relation_arr[i] = false;
    }
    for (uint i = 0; i < n; i++)
    {
        if (!relation_arr[i])
        {
            vector<int> cur_vec = G[i];
            bool is_relation = false;
            for (uint j = 0; j < cur_vec.size(); j++)
            {
                if (cur_vec[j] == S)
                {
                    is_relation = true;
                    break;
                }
            }
            if (is_relation)
            {
                TopKMovieSub(relation_arr, G, i, S);
            }
        }
    }
    for (uint i = 0; i < n; i++)
    {
        if (relation_arr[i])
        {
            relation_vec.push_back(pair<int, int>(i, rating[i]));
        }
    }
    sort(relation_vec.begin(), relation_vec.end(), CmpByValue());
    for (uint i = 0; i < K && i < relation_vec.size(); i++)
    {
        result.push_back(relation_vec[i].first);
    }
    return result;
}

int BuyFruits(vector<vector<string>> &codeList, vector<string> &shoppingCart) // LintCode 806. ¬ÚÀÆπ˚
{
    if (codeList.size() == 0)
    {
        return 1;
    }
    
    int code_num = 0;
    int shopping_num = shoppingCart.size();
    for (uint i = 0; i < codeList.size(); i++)
    {
        code_num += codeList[i].size();
    }
    
    for (int z = 0; z <= shopping_num - code_num; z++)
    {
        int i = 0;
        int j = 0;
        for (int cur_z = z;; cur_z++)
        {
            if (codeList[i][j] != "anything" && codeList[i][j] != shoppingCart[cur_z])
            {
                break;
            }
            j++;
            if (j == codeList[i].size())
            {
                i++;
                if (i == codeList.size())
                {
                    return 1;
                }
                j = 0;
            }
        }
    }
    return 0;
}

void MaximumAssociationSetSub(vector<string> &cur_relation, map<string, vector<string>> &relation_map, map<string, bool> &relation_visit, string str)
{
    if (!relation_visit[str])
    {
        relation_visit[str] = true;
        cur_relation.push_back(str);
    }
    for (uint i = 0; i < relation_map[str].size(); i++)
    {
        if (!relation_visit[relation_map[str][i]])
        {
            MaximumAssociationSetSub(cur_relation, relation_map, relation_visit, relation_map[str][i]);
        }
    }
}

vector<string> MaximumAssociationSet(vector<string> &ListA, vector<string> &ListB) //LintCode 805. ◊Ó¥Ûπÿ¡™ºØ∫œ
{
    map<string, vector<string>> relation_map;
    map<string, bool> relation_visit;
    vector<string> result;
    uint max_size = 0;
    int n = relation_map.size();
    
    for (uint i = 0; i < ListA.size(); i++)
    {
        relation_map[ListA[i]].push_back(ListB[i]);
        relation_map[ListB[i]].push_back(ListA[i]);
        relation_visit[ListA[i]] = false;
        relation_visit[ListB[i]] = false;
    }
    map<string, vector<string>>::iterator it_map = relation_map.begin();
    for (; it_map != relation_map.end(); it_map++)
    {
        vector<string> cur_relation;
        MaximumAssociationSetSub(cur_relation, relation_map, relation_visit, (*it_map).first);
        if (max_size < cur_relation.size())
        {
            max_size = cur_relation.size();
            result = cur_relation;
        }
    }
    return result;
}

string InputStream(string &inputA, string &inputB) //LintCode 823.  ‰»Î¡˜  ’ºŸµ⁄“ªÃ‚£¨…Ò∑≥£¨ºŸ∆⁄∂º√ª”–º˚µΩœÎº˚µƒ»À£¨√˜√˜‘º∫√¡À
{
    stack<char> A_stack;
    stack<char> B_stack;
    
    for (uint i = 0; i < inputA.length(); i++)
    {
        if (inputA[i] == '<')
        {
            if (!A_stack.empty())
            {
                A_stack.pop();
            }
        }
        else
        {
            A_stack.push(inputA[i]);
        }
    }
    for (uint i = 0; i < inputB.length(); i++)
    {
        if (inputB[i] == '<')
        {
            if (!B_stack.empty())
            {
                B_stack.pop();
            }
        }
        else
        {
            B_stack.push(inputB[i]);
        }
    }
    
    if (A_stack.size() != B_stack.size())
    {
        return "NO";
    }
    while (!A_stack.empty())
    {
        if (A_stack.top() != B_stack.top())
        {
            return "NO";
        }
        A_stack.pop();
        B_stack.pop();
    }
    return "YES";
}

int MaxDiff(vector<vector<int>> &arrs) //LintCode 698.  ˝◊È÷–◊Ó¥Ûµƒ≤Ó÷µ
{
    if (arrs.empty())
    {
        return 0;
    }
    
    map<int, int> arrs_map;
    map<int, int> relation_map;
    
    for (uint i = 0; i < arrs.size(); i++)
    {
        int cur_size = arrs[i].size();
        arrs_map[arrs[i][0]]++;
        arrs_map[arrs[i][cur_size - 1]]++;
        relation_map[arrs[i][0]] = arrs[i][cur_size - 1];
        relation_map[arrs[i][cur_size - 1]] = arrs[i][0];
    }
    
    map<int, int>::iterator map_it_begin = arrs_map.begin();
    map<int, int>::iterator map_it_end = arrs_map.end();
    map_it_end--;
    if ((*map_it_begin).second == 1 && (*map_it_end).second == 1)
    {
        if ((*map_it_begin).second == 1 && relation_map[(*map_it_end).first] == (*map_it_begin).first)
        {
            if ((*map_it_end).second == 1)
            {
                int tmp_a = (*map_it_begin).first;
                int tmp_b = (*map_it_end).first;
                return max(abs((*--map_it_end).first - tmp_a), abs((*++map_it_begin).first - tmp_b));
            }
            map_it_begin++;
        }
        else if ((*map_it_end).second == 1 && relation_map[(*map_it_begin).first] == (*map_it_end).first)
        {
            map_it_end--;
        }
    }
    return abs((*map_it_end).first - (*map_it_begin).first);
}

int GetSingleNumber(vector<int> &nums) // LintCode 824. ¬‰µ•µƒ ˝ IV
{
    int min = 0;
    int max = nums.size() - 1;
    if (max == 1)
    {
        return nums[0];
    }
    while (true)
    {
        int middle = (max - min) / 2 + min;
        if (middle % 2 != 0)
        {
            middle--;
        }
        if (nums[middle] == nums[middle + 1])
        {
            min = middle + 2;
            continue;
        }
        else if (nums[middle] == nums[middle - 1])
        {
            max = middle - 2;
            continue;
        }
        else
        {
            return nums[middle];
        }
    }
}

void BinaryTimeSub(vector<string> &result, int *arr, int num, int dex) // LintCode 706. ∂˛Ω¯÷∆ ±º‰ ”√∫Í∫Õπ´ ΩÃ´¬È∑≥¡À£¨÷±Ω”–¥≥… ˝◊÷£¨≤ª“™‘⁄“‚ƒ«√¥∂‡
{
    if (num == 0)
    {
        // ≈–∂œ–° ± «∑Ò≥¨π˝œﬁ÷∆
        int hour = 8 * arr[0] + 4 * arr[1] + 2 * arr[2] + arr[3];
        if (hour >= 12)
        {
            return;
        }
        int min = 32 * arr[4] + 16 * arr[5] + 8 * arr[6] + 4 * arr[7] + 2 * arr[8] + arr[9];
        if (min >= 60)
        {
            return;
        }
        string cur_time = to_string(hour);
        cur_time += ":";
        cur_time += min < 10 ? "0" + to_string(min) : to_string(min);
        result.push_back(cur_time);
        return;
    }
    for (int i = dex; i <= 10 - num; i++)
    {
        arr[i] = 1;
        BinaryTimeSub(result, arr, num - 1, i + 1);
        arr[i] = 0;
    }
}

vector<string> BinaryTime(int num) // LintCode 706. ∂˛Ω¯÷∆ ±º‰ ”√∫Í∫Õπ´ ΩÃ´¬È∑≥¡À£¨÷±Ω”–¥≥… ˝◊÷£¨≤ª“™‘⁄“‚ƒ«√¥∂‡
{
    int arr[10] = { 0 };
    vector<string> result;
    BinaryTimeSub(result, arr, num, 0);
    return result;
}

string HexConversion(int n, int k) // LintCode 763. Hex Conversion
{
    string result;
    while (n > 0)
    {
        int cur_num = n % k;
        if (cur_num >= 10)
        {
            char s = cur_num - 10 + 'A';
            result = s + result;
        }
        else
        {
            result = to_string(cur_num) + result;
        }
        n -= cur_num;
        n /= k;
    }
    if (result.length() == 0)
    {
        result = "0";
    }
    return result;
}

char LowercaseToUppercase(char character) // LintCode 145. ¥Û–°–¥◊™ªª
{
    return character <= 'Z' ? character - 'A' + 'a' : character - 'a' + 'A';
}

int reverseInteger(int number) // LintCode 37. ∑¥◊™“ª∏ˆ3Œª’˚ ˝
{
    int hundred = number / 100;
    int teen = (number - hundred * 100) / 10;
    int bit = number - hundred * 100 - teen * 10;
    return hundred + teen * 10 + bit * 100;
}

vector<vector<int>> SpiralArray(int n) // LintCode 769. ¬›–˝æÿ’Û
{
    vector<vector<int>> result;
    for (int i = 0; i < n; i++)
    {
        vector<int> cur_vec;
        for (int j = 0; j < n; j++)
        {
            cur_vec.push_back(0);
        }
        result.push_back(cur_vec);
    }
    
    int start_i = 0;
    int end_i = n - 1;
    int start_j = 0;
    int end_j = n - 1;
    int num = 1;
    while (start_i <= end_i || start_j <= end_j)
    {
        int i = start_i;
        int j = start_j;
        for (; i <= end_i; i++)
        {
            result[j][i] = num++;
        }
        i = end_i;
        j++;
        for (; j <= end_j; j++)
        {
            result[j][i] = num++;
        }
        j = end_j;
        i--;
        for (; i >= start_i; i--)
        {
            result[j][i] = num++;
        }
        i = start_i;
        j--;
        for (; j > start_j; j--)
        {
            result[j][i] = num++;
        }
        start_i++;
        start_j++;
        end_i--;
        end_j--;
    }
    return result;
}

bool IsLeapYear(int n) // LintCode 766. »ÚƒÍ
{
    return ((n % 4 == 0) && (n % 100 != 0)) || (n % 400 == 0);
}

bool IsValidTriangle(int a, int b, int c) //LintCode 765. Valid Triangle
{
    int min, middle, max;
    if (a < b)
    {
        if (c < a)
        {
            min = c;
            middle = a;
            max = b;
        }
        else if (c < b)
        {
            min = a;
            middle = c;
            max = b;
        }
        else
        {
            min = a;
            middle = b;
            max = c;
        }
    }
    else
    {
        if (c < b)
        {
            min = c;
            middle = b;
            max = a;
        }
        else if (c < a)
        {
            min = b;
            middle = c;
            max = a;
        }
        else
        {
            min = b;
            middle = a;
            max = c;
        }
    }
    return min + middle > max;
}

vector<double> Calculate(int r) // LintCode 764. Calculate Circumference And Area
{
    vector<double> result;
    result.push_back(2 * r * CIRCLE_PI);
    result.push_back(CIRCLE_PI * pow(r, 2));
    return result;
}

vector<vector<int>> CalcYangHuisTriangle(int n) // LintCode 768. —Óª‘»˝Ω«
{
    vector<vector<int>> result;
    if (n == 0)
    {
        return result;
    }
    result.push_back(vector<int>({ 1 }));
    for (int i = 2; i <= n; i++)
    {
        vector<int> cur_vec({ 1 });
        for (int j = 2; j < i; j++)
        {
            cur_vec.push_back(result[i - 2][j - 2] + result[i - 2][j - 1]);
        }
        cur_vec.push_back(1);
        result.push_back(cur_vec);
    }
    return result;
}

void ReverseArray(vector<int> &nums) // LintCode 767. ∑≠◊™ ˝◊È
{
    int n = nums.size();
    for (int i = 0; i < n / 2; i++)
    {
        int tmp = nums[i];
        nums[i] = nums[n - 1 - i];
        nums[n - 1 - i] = tmp;
    }
}

vector<int> MaxAndMin(vector<vector<int>> &matrix) // LintCode 770. Maximum and Minimum
{
    vector<int> result;
    if (matrix.empty())
    {
        return result;
    }
    int max = matrix[0][0];
    int min = matrix[0][0];
    for (uint i = 0; i < matrix.size(); i++)
    {
        for (uint j = 0; j < matrix[i].size(); j++)
        {
            if (matrix[i][j] < min)
            {
                min = matrix[i][j];
            }
            if (matrix[i][j] > max)
            {
                max = matrix[i][j];
            }
        }
    }
    result.push_back(max);
    result.push_back(min);
    return result;
}

bool WordSortCmp(string a, string b)
{
    return a < b;
}

vector<string> WordSort(string &alphabet, vector<string> &words) // LintCode 819. µ•¥ ≈≈–Ú
{
    vector<string> result = words;
    char char_hash[SMALL_ENGLISH_CHAR_NUM];
    for (int i = 0; i < SMALL_ENGLISH_CHAR_NUM; i++)
    {
        char_hash[alphabet[i] - 'a'] = i + 'a';
    }
    for (uint i = 0; i < result.size(); i++)
    {
        for (uint j = 0; j < result[i].length(); j++)
        {
            result[i][j] = char_hash[result[i][j] - 'a'];
        }
    }
    sort(result.begin(), result.end(), WordSortCmp);
    for (uint i = 0; i < result.size(); i++)
    {
        for (uint j = 0; j < result[i].length(); j++)
        {
            result[i][j] = alphabet[result[i][j] - 'a'];
        }
    }
    return result;
}

class Point
{
public:
    int x, y;
    Point(int x, int y)
    {
        this->x = x;
        this->y = y;
    }
};

string Rectangle(vector<Point> &pointSet) // LintCode 820. æÿ–Œ
{
    unordered_set<string> PointSet;
    for (uint i = 0; i < pointSet.size(); i++)
    {
        string cur_str = to_string(pointSet[i].x) + "," + to_string(pointSet[i].y);
        PointSet.insert(cur_str);
    }
    
    for (uint i = 0; i < pointSet.size(); i++)
    {
        for (uint j = i + 1; j < pointSet.size(); j++)
        {
            if ((pointSet[i].x == pointSet[j].x) || (pointSet[i].y == pointSet[j].y))
            {
                continue;
            }
            string str1 = to_string(pointSet[i].x) + "," + to_string(pointSet[j].y);
            string str2 = to_string(pointSet[j].x) + "," + to_string(pointSet[i].y);
            if (PointSet.count(str1) && PointSet.count(str2))
            {
                return "YES";
            }
        }
    }
    return "NO";
}

int FriendRequest(vector<int> &ages) // LintCode 895. ∫√”—«Î«Û
{
    int result = 0;
    for (uint i = 0; i < ages.size(); i++)
    {
        for (uint j = i + 1; j < ages.size(); j++)
        {
            int age1 = ages[i];
            int age2 = ages[j];
            if ((age1 >(age2 / 2 + 7)) && (age1 <= age2) && (age1 >= 100 || age2 <= 100))
            {
                result++;
            }
            if ((age2 >(age1 / 2 + 7)) && (age2 <= age1) && (age2 >= 100 || age1 <= 100))
            {
                result++;
            }
        }
    }
    return result;
}

int PairNumbers(vector<Point> &p) // LintCode 844.  ˝∂‘Õ≥º∆
{
    int result = 0;
    for (uint i = 0; i < p.size() - 1; i++)
    {
        for (uint j = i + 1; j < p.size(); j++)
        {
            if (((p[i].x + p[j].x) % 2 == 0) && ((p[i].y + p[j].y) % 2 == 0))
            {
                result++;
            }
        }
    }
    return result;
}

bool ClosestValueSub(TreeNode<int> * root, double target, int &value) // LintCode 900. Closest Binary Search Tree Value
{
    if (root == NULL)
    {
        return false;
    }
    value = root->val;
    int left_value, right_value;
    bool left_effective, right_effective;
    left_effective = ClosestValueSub(root->left, target, left_value);
    right_effective = ClosestValueSub(root->right, target, right_value);
    if (left_effective && right_effective && abs(left_value - target) < abs(right_value - target))
    {
        value = left_value;
    }
    else if (left_effective && right_effective)
    {
        value = right_value;
    }
    else
    {
        value = left_effective ? left_value : right_value;
    }
    
    if (abs(root->val - target) < abs(value - target))
    {
        value = root->val;
    }
    return true;
}

int ClosestValue(TreeNode<int> * root, double target) // LintCode 900. Closest Binary Search Tree Value
{
    if (root == NULL)
    {
        return 0;
    }
    int value;
    ClosestValueSub(root, target, value);
    return value;
}

bool ValidWordSquare(vector<string> &words) // LintCode 888. Valid Word Square
{
    for (uint i = 0; i < words.size(); i++)
    {
        for (uint j = 0; j < words[i].length(); j++)
        {
            if (words[i][j] != words[j][i])
            {
                return false;
            }
        }
    }
    return true;
}

string SameNumber(vector<int> &nums, int k) // LintCode 1368. œ‡Õ¨ ˝◊÷
{
    map<int, int> num_map;
    for (uint i = 0; i < nums.size(); i++)
    {
        int cur_num = nums[i];
        if (num_map.count(cur_num))
        {
            if ((i - num_map[cur_num]) < (uint)k)
            {
                return "YES";
            }
        }
        num_map[cur_num] = i;
    }
    return "NO";
}

int DigitConvert(int n) // LintCode 952.  ˝◊÷Œ Ã‚
{
    int num = 0;
    while (n != 1)
    {
        if (n % 2 == 0)
        {
            n /= 2;
        }
        else
        {
            n = 3 * n + 1;
        }
        num++;
    }
    return num;
}

int ReachNumber(int target)
{
    int num = 0;
    int once_len = 1;
    int cur_dex = 0;
    if (target > 0)
    {
        while (cur_dex + once_len <= target)
        {
            cur_dex += once_len;
            once_len++;
            num++;
        }
        if ((target - cur_dex) > 0)
        {
            int sur = once_len - (target - cur_dex);
            if (sur % 2 == 0)
            {
                num++;
            }
            else
            {
                if (once_len % 2 == 0)
                {
                    num += 2;
                }
                else
                {
                    num += 3;
                }
            }
        }
    }
    else if (target < 0)
    {
        while (cur_dex - once_len >= target)
        {
            cur_dex -= once_len;
            once_len++;
            num++;
        }
        if ((cur_dex - target) > 0)
        {
            int sur = once_len - (cur_dex - target);
            if (sur % 2 == 0)
            {
                num++;
            }
            else
            {
                if (once_len % 2 == 0)
                {
                    num += 2;
                }
                else
                {
                    num += 3;
                }
            }
        }
    }
    return num;
}

vector<int> AnagramMappings(vector<int> &A, vector<int> &B) // LintCode 813. ’“µΩ”≥…‰–Ú¡–
{
    map<int, int> B_map;
    vector<int> result;
    for (uint i = 0; i < B.size(); i++)
    {
        B_map[B[i]] = i;
    }
    for (uint i = 0; i < A.size(); i++)
    {
        result.push_back(B_map[A[i]]);
    }
    return result;
}

vector<string> logSort(vector<string> &logs) // LintCode 1380. »’÷æ≈≈–Ú -- ’‚À˚¬Ë «ºÚµ•?
{
    vector<string> Result;
    vector<string> Num_Vec;
    for (uint i = 0; i < logs.size(); i++)
    {
        string cur_str = logs[i];
        int space_dex = cur_str.find(' ');
        string key = cur_str.substr(0, space_dex);
        string value = cur_str.substr(space_dex + 1, cur_str.length());
        if (value[0] <= '9' && value[0] >= '0') // ˝◊÷£¨≤ª”√≈≈–Ú
        {
            Num_Vec.push_back(cur_str);
        }
        else
        {
            if (Result.empty())
            {
                Result.push_back(cur_str);
            }
            else
            {
                for (uint j = 0; j < Result.size(); j++)
                {
                    string Res_str = Result[j];
                    int cur_space_index = Res_str.find(' ');
                    string cur_key = Res_str.substr(0, cur_space_index);
                    string cur_value = Res_str.substr(cur_space_index + 1, Res_str.length());
                    if (value < cur_value)
                    {
                        Result.insert(Result.begin() + j, cur_str);
                        break;
                    }
                    if (value == cur_value)
                    {
                        if (key < cur_key)
                        {
                            Result.insert(Result.begin() + j, cur_str);
                            break;
                        }
                    }
                    if (j == Result.size() - 1)
                    {
                        Result.push_back(cur_str);
                        break;
                    }
                }
            }
        }
    }
    Result.insert(Result.end(), Num_Vec.begin(), Num_Vec.end());
    return Result;
}

bool NumIslandCitiesSub(vector<vector<int>> &grid, uint i, uint j, vector<vector<bool>> &IsVisitGrid)
{
    bool Resut = false;
    bool SubResult = false;
    uint x = grid.size();
    uint y = grid[0].size();
    
    IsVisitGrid[i][j] = true;
    if (grid[i][j] == 2)
    {
        Resut = true;
    }
    else  if (grid[i][j] == 0)
    {
        return false;
    }
    if (i + 1 < x && !IsVisitGrid[i + 1][j])
    {
        SubResult = SubResult | NumIslandCitiesSub(grid, i + 1, j, IsVisitGrid);
    }
    if ((int)i - 1 >= 0 && !IsVisitGrid[i - 1][j])
    {
        SubResult = SubResult | NumIslandCitiesSub(grid, i - 1, j, IsVisitGrid);
    }
    if (j + 1 < y && !IsVisitGrid[i][j + 1])
    {
        SubResult = SubResult | NumIslandCitiesSub(grid, i, j + 1, IsVisitGrid);
    }
    if ((int)j - 1 >= 0 && !IsVisitGrid[i][j - 1])
    {
        SubResult = SubResult | NumIslandCitiesSub(grid, i, j - 1, IsVisitGrid);
    }
    return Resut | SubResult;
}

int NumIslandCities(vector<vector<int>> &grid) // LintCode 897. ∫£µ∫≥« –
{
    int Result = 0;
    uint x = grid.size();
    uint y = grid[0].size();
    vector<vector<bool>> IsVisitGrid;
    for (uint i = 0; i < x; i++)
    {
        vector<bool> CurVec;
        for (uint j = 0; j < y; j++)
        {
            CurVec.push_back(false);
        }
        IsVisitGrid.push_back(CurVec);
    }
    
    for (uint i = 0; i < x; i++)
    {
        for (uint j = 0; j < y; j++)
        {
            if (!IsVisitGrid[i][j] && NumIslandCitiesSub(grid, i, j, IsVisitGrid))
            {
                Result++;
            }
        }
    }
    return Result;
}

int gcd(int a, int b) // LintCode 845. Greatest Common Divisor
{
    if (a == b)
    {
        return a;
    }
    int big_num;
    int small_num;
    
    if (a > b)
    {
        big_num = a;
        small_num = b;
    }
    else
    {
        big_num = b;
        small_num = a;
    }
    while (big_num % small_num != 0)
    {
        int tmp = big_num;
        big_num = small_num;
        small_num = tmp % small_num;
    }
    return small_num;
}

int SubarraySumEqualsK(vector<int> &nums, int k) // LintCode 838. Subarray Sum Equals K
{
    vector<int> sum_vec;
    unordered_map<int, int> hash;
    int cur_sum = 0;
    int res = 0;
    hash[0] = 1; // ”√¿¥¥¶¿Ì“ª∏ˆsumæÕµ»”⁄kµƒ«Èøˆ
    
    for (uint i = 0; i < nums.size(); i++)
    {
        cur_sum += nums[i];
        sum_vec.push_back(cur_sum);
    }
    
    for (uint i = 0; i < nums.size(); i++)
    {
        if (hash[sum_vec[i] - k])
        {
            res += hash[sum_vec[i] - k];
        }
        hash[sum_vec[i]]++;
    }
    return res;
}

vector<vector<int>> TwitchWords(string &str) // LintCode 1401. ≥È¥§¥
{
    vector<vector<int>> res;
    uint start_index = 0;
    int last_word = str[0];
    
    for (uint i = 1; i < str.length(); i++)
    {
        if (last_word != str[i])
        {
            if (i - start_index > 2)
            {
                res.push_back(vector<int>{(int)start_index, (int)i - 1});
            }
            start_index = i;
            last_word = str[i];
        }
    }
    if (str.length() - start_index > 2)
    {
        res.push_back(vector<int>{(int)start_index, (int)str.length() - 1});
    }
    return res;
}

class MovingAverage // LintCode 642. Moving Average from Data Stream
{
    int size;
    vector<int> val_vec;
    long long sum = 0;
    uint start_index = 0;
public:
    
    MovingAverage(int size)
    {
        this->size = size;
    }
    
    double next(int val)
    {
        val_vec.push_back(val);
        sum += val;
        if (val_vec.size() - start_index > (uint)size)
        {
            sum -= val_vec[start_index++];
            return (double)(sum / size);
        }
        else
        {
            return (double)(sum / val_vec.size());
        }
    }
};

bool IsSentenceSimilarity(vector<string> &words1, vector<string> &words2, vector<vector<string>> &pairs) // LintCode 856. Sentence Similarity
{
    unordered_map<string, vector<string>> word_map;
    for (uint i = 0; i < pairs.size(); i++)
    {
        vector<string> cur_vec = pairs[i];
        if (word_map.count(cur_vec[0]))
        {
            word_map[cur_vec[0]].push_back(cur_vec[1]);
        }
        else
        {
            word_map[cur_vec[0]] = vector<string>{cur_vec[1]};
        }
        if (word_map.count(cur_vec[1]))
        {
            word_map[cur_vec[1]].push_back(cur_vec[0]);
        }
        else
        {
            word_map[cur_vec[1]] = vector<string>{cur_vec[0]};
        }
    }
    for (uint i = 0; i < words1.size(); i++)
    {
        if (!word_map.count(words1[i]))
        {
            return false;
        }
        
        bool is_find = false;
        for (uint j = 0; j < word_map[words1[i]].size(); j++)
        {
            if (word_map[words1[i]][j] == words2[i])
            {
                is_find = true;
            }
        }
        if (!is_find)
        {
            return false;
        }
    }
    return true;
}

bool RadarDetection2Sub(int a[2], int b[2])
{
    return a[1] <= b[1];
}

string RadarDetection2(vector<Point> &coordinates, vector<int> &radius) //LintCode ≤ª∫√“‚Àº£¨Œ“œÎ∂‡¡À
{
    /*vector<int[2]> radius_vec;
     for (uint i = 0; i < coordinates.size(); i++)
     {
     Point cur_point = coordinates[i];
     if (cur_point.y <= 0 && cur_point.y + radius[i] > 0)
     {
     int tmp[2] = { 0, 0 };
     tmp[1] = cur_point.y + radius[i];
     radius_vec.push_back(tmp);
     }
     else if (cur_point.y >= 1 && cur_point.y - radius[i] < 0)
     {
     int tmp[2] = { 1, 1 };
     tmp[0] = cur_point.y - radius[i];
     radius_vec.push_back(tmp);
     }
     }
     sort(radius_vec.begin(), radius_vec.end(), RadarDetection2Sub);
     if (radius_vec[0][0] != 0)
     {
     return "NO";
     }
     
     for (uint i = 0; radius_vec.size(); i++)
     {
     //◊ˆ“ª∏ˆ≈≈¡–
     }*/
    return "YES";
}

string RadarDetection(vector<Point> &coordinates, vector<int> &radius) //LintCode ¿◊¥ÔºÏ≤‚
{
    for (uint i = 0; i < coordinates.size(); i++)
    {
        Point cur_p = coordinates[i];
        if (cur_p.y < 0 && radius[i] >= (1 - cur_p.y))
        {
            return "YES";
        }
        if (cur_p.y > 1 && radius[i] >= (cur_p.y - 0))
        {
            return "YES";
        }
    }
    return "NO";
}

bool WordPattern(string &pattern, string &str) // LintCode 828. ◊÷ƒ£ Ω
{
    unordered_map<string, char> value_to_key;
    bool key[SMALL_ENGLISH_CHAR_NUM] = { false };
    string value[SMALL_ENGLISH_CHAR_NUM];
    string cur_str = str;
    
    for (uint i = 0; i < pattern.size(); i++)
    {
        int index = pattern[i] - 'a';
        int find_space = cur_str.find(" ");
        string sub_str = cur_str.substr(0, find_space);
        cur_str = cur_str.substr(find_space + 1, cur_str.size());
        if (!key[index])
        {
            key[index] = true;
            value[index] = sub_str;
        }
        else
        {
            if (value[index] != sub_str)
            {
                return false;
            }
        }
        if (value_to_key.count(sub_str))
        {
            if (pattern[i] != value_to_key[sub_str])
            {
                return false;
            }
        }
        else
        {
            value_to_key[sub_str] = pattern[i];
        }
    }
    return true;
}

int HammingDistance(int x, int y) // LintCode 835. Hamming Distance
{
    int res = 0;
    long tmp = x ^ y;
    while (tmp)
    {
        /*if ((tmp & 0x01) != 0)
         {
         res++;
         }
         tmp = tmp >> 1; ∏–æıÕÍ»´Œª‘ÀÀ„∑¥∂¯∏¸øÏ */
        res++;
        tmp &= tmp - 1;
    }
    return res;
}

class NumArray { // LintCode 943. Range Sum Query - Immutable
    vector<int> sum_array;
public:
    NumArray(vector<int> nums) {
        int sum = 0;
        for (uint i = 0; i < nums.size(); i++)
        {
            sum += nums[i];
            sum_array.push_back(sum);
        }
    }
    
    int sumRange(int i, int j) {
        if (i == 0)
        {
            return sum_array[j];
        }
        return sum_array[j] - sum_array[i - 1];
    }
};

struct BigBusinessKeyValue
{
    int key;
    int value;
};

bool BigBusinessSub(BigBusinessKeyValue a, BigBusinessKeyValue b)
{
    return a.key <= b.key;
}

int BigBusiness(vector<int> &a, vector<int> &b, int k) // LintCode 970. ¥Û…˙“‚
{
    vector<BigBusinessKeyValue> vec_data;
    
    for (uint i = 0; i < a.size(); i++)
    {
        BigBusinessKeyValue tmp;
        tmp.key = a[i];
        tmp.value = b[i];
        vec_data.push_back(tmp);
    }
    sort(vec_data.begin(), vec_data.end(), BigBusinessSub);
    
    for (uint i = 0; i < vec_data.size(); i++)
    {
        if (k < vec_data[i].key)
        {
            return k;
        }
        
        int dec = vec_data[i].value - vec_data[i].key;
        if (dec <= 0)
        {
            continue;
        }
        k += dec;
    }
    return k;
}

int CalPoints(vector<string> &ops) // LintCode 983. Baseball Game
{
    vector<int> data_vec;
    int sum = 0;
    
    for (uint i = 0; i < ops.size(); i++)
    {
        int add_num = 0;
        string cur_str = ops[i];
        if (cur_str[0] == 'D')
        {
            int last_num = data_vec.back() * 2;
            data_vec.push_back(last_num);
            add_num += last_num;
        }
        else if (cur_str[0] == 'C')
        {
            int last_num = data_vec.back();
            data_vec.pop_back();
            add_num += -last_num;
        }
        else if (cur_str[0] == '+')
        {
            int last_num = data_vec.back();
            int next_num = data_vec[data_vec.size() - 2];
            data_vec.push_back(last_num + next_num);
            add_num += last_num + next_num;
        }
        else
        {
            int cur_num = stoi(cur_str);
            add_num += cur_num;
            data_vec.push_back(cur_num);
        }
        sum += add_num;
    }
    return sum;
}

bool HasAlternatingBits(int n) // LintCode 987. Binary Number with Alternating Bits
{
    bool is_zero = (n % 2) ? false : true;
    n /= 2;
    while (n)
    {
        if (is_zero && (n % 2))
        {
            is_zero = false;
        }
        else if (!is_zero && !(n % 2))
        {
            is_zero = true;
        }
        else
        {
            return false;
        }
        n /= 2;
    }
    return true;
}

int FindTiltSub(TreeNode<int> * root, int &result)
{
    int left_val = 0;
    int right_val = 0;
    if (root->left != NULL)
    {
        left_val = FindTiltSub(root->left, result);
    }
    if (root->right != NULL)
    {
        right_val = FindTiltSub(root->right, result);
    }
    result += abs(left_val - right_val);
    return left_val + right_val + root->val;
}

int FindTilt(TreeNode<int> * root) // LintCode 1172. Binary Tree Tilt
{
    if (root == NULL)
    {
        return 0;
    }
    
    int result = 0;
    FindTiltSub(root, result);
    return result;
}

bool CanPlaceFlowers(vector<int> &flowerbed, int n) // LintCode 1138. Can Place Flowers
{
    int last_not_empty = -1;
    for (uint i = 0; i < flowerbed.size(); i++)
    {
        if (flowerbed[i])
        {
            int dictance = i - 1 - last_not_empty;
            last_not_empty = i;
            int num = (dictance - 1) / 2;
            if (n <= num)
            {
                return true;
            }
            n -= num;
        }
    }
    return false;
}

int MinDiffInBSTSub(TreeNode<int> * root, int &min_num, int &max_num)
{
    int cur_root_min = MAXINTNUM;
    if (root->left)
    {
        int tmp_max = MININTNUM;
        int tmp = MinDiffInBSTSub(root->left, min_num, tmp_max);
        cur_root_min = min(cur_root_min, tmp);
        cur_root_min = min(cur_root_min, root->val - tmp_max);
    }
    if (root->right)
    {
        int tmp_min = MAXINTNUM;
        int tmp = MinDiffInBSTSub(root->right, tmp_min, max_num);
        cur_root_min = min(cur_root_min, tmp);
        cur_root_min = min(cur_root_min, tmp_min - root->val);
    }
    
    min_num = min(min_num, root->val);
    max_num = max(max_num, root->val);
    return cur_root_min;
}

int MinDiffInBST(TreeNode<int> * root) // LintCode 1033. Minimum Distance Between BST Nodes   ;;;   1188. Minimum Absolute Difference in BST  LintCode’Êµƒ‘Ω¿¥‘Ω¿√µƒ£¨Õ¨∏ˆÀ„∑®¡Ω∏ˆÃ‚
{
    int min_num = MAXINTNUM;
    int max_num = MININTNUM;
    if (root == NULL)
    {
        return -1;
    }
    
    return MinDiffInBSTSub(root, min_num, max_num);
}

string Compress(string &str) // LintCode 213. ◊÷∑˚¥Æ—πÀı
{
    string result;
    char last_char = str[0];
    char last_size = 1;
    
    for (uint i = 1; i < str.length(); i++)
    {
        if (last_char != str[i])
        {
            result += last_char + to_string(last_size);
            last_size = 1;
            last_char = str[i];
        }
        else
        {
            last_size++;
        }
    }
    result += last_char + to_string(last_size);
    
    return result.length() < str.length() ? result : str;
}

vector<int> ReverseStore(ListNode<int> * head) // LintCode 822. œ‡∑¥µƒÀ≥–Ú¥Ê¥¢
{
    vector<int> result;
    while (head != NULL)
    {
        result.push_back(head->val);
        head = head->next;
    }
    
    uint vec_size = result.size();
    
    for (uint i = 0; i < vec_size / 2; i++)
    {
        int tmp = result[i];
        result[i] = result[vec_size - i - 1];
        result[vec_size - i - 1] = tmp;
    }
    
    return result;
}

int FindContentChildren(vector<int> &g, vector<int> &s) // LintCode 1230. Assign Cookies
{
    sort(g.begin(), g.end());
    sort(s.begin(), s.end());
    
    uint g_i = 0, s_i = 0;
    int result = 0;
    for (; g_i < g.size() && s_i < s.size(); s_i++)
    {
        if (g[g_i] <= s[s_i])
        {
            g_i++;
            result++;
            continue;
        }
    }
    return result;
}

int MaxCount(int m, int n, vector<vector<int>> &ops) // 1144. Range Addition II
{
    int x = m;
    int y = n;
    for (auto it : ops)
    {
        if (it[0] <= x && it[0] > 0)
        {
            x = it[0];
        }
        if (it[1] <= y && it[1] > 0)
        {
            y = it[1];
        }
    }
    return x * y;
}

bool IsToeplitzMatrix(vector<vector<int>> &matrix) // 1042. Toeplitz Matrix
{
    for (uint i = matrix.size() - 1, j = 0; i != 0 || j != matrix[0].size();)
    {
        int cur_num = matrix[i][j];
        for (uint tmp_i = i + 1, tmp_j = j + 1; tmp_i < matrix.size() && tmp_j < matrix[0].size(); tmp_i++, tmp_j++)
        {
            if (cur_num != matrix[tmp_i][tmp_j])
            {
                return false;
            }
        }
        if (i > 0)
        {
            i--;
        }
        else
        {
            j++;
        }
    }
    return true;
}

string LastFourDigitsOfFn2(int n) // 949. Fibonacci II // «Û«∞º∏Œªµƒ£¨µıµıµı Fibonacciπ´ Ω Fn = £®°Ã5 / 5£© * pow(£®£®1 + °Ã5£©/ 2£© , n) - £®°Ã5 / 5£© * pow(£®£®1 - °Ã5£©/ 2£© , n)
{
    if (n == 0)
    {
        return "0";
    }
    
    int last_num = 0;
    int cur_num = 1;
    if (n <= 20)
    {
        for (int i = 2; i <= n; i++)
        {
            int tmp = (last_num + cur_num) % 10000;
            last_num = cur_num;
            cur_num = tmp;
        }
    }
    else
    {
        double a = (1.0 + sqrt(5.0)) / 2.0;
        double ans = n*log10(a) - 0.5*log10(5.0);
        ans = ans - floor(ans);
        ans = pow(10, ans);
        int res = (int)(ans * 1000);
        printf("%d\n", res);
    }
    return to_string(cur_num);
}

bool LastFourDigitsOfFnSub(vector<int>& matrix, int & start, int n) // µ›πÈ∫√◊ˆ“ª–©
{
    int i = 2;
    vector<int> base_matrix = { 1, 1, 0 };
    for (; start + i <= n; i *= 2)
    {
        int base_fn_max = base_matrix[0] * base_matrix[0] + base_matrix[1] * base_matrix[1];
        int base_fn = base_matrix[0] * base_matrix[1] + base_matrix[1] * base_matrix[2];
        int base_fn_min = base_matrix[1] * base_matrix[1] + base_matrix[2] * base_matrix[2];
        base_matrix[0] = base_fn_max % 10000;
        base_matrix[1] = base_fn % 10000;
        base_matrix[2] = base_fn_min % 10000;
    }
    
    int Fn_max = matrix[0] * base_matrix[0] + matrix[1] * base_matrix[1];
    int Fn = matrix[0] * base_matrix[1] + matrix[1] * base_matrix[2];
    int Fn_min = matrix[2] * base_matrix[1] + matrix[3] * base_matrix[2];
    matrix[0] = Fn_max % 10000;
    matrix[1] = Fn % 10000;
    matrix[2] = matrix[1];
    matrix[3] = Fn_min % 10000;
    
    if (start + i / 2 == n)
    {
        return true;
    }
    start += i / 2;
    return false;
}

string LastFourDigitsOfFn(int n) // 949. Fibonacci II // «Û∫Ûº∏Œªµƒ£¨æÿ’Ûœ‡≥À∑®£¨ø…“‘Ω´ ±º‰∏¥‘”∂»ŒﬁœﬁΩ”Ω¸”Îlog(n)£¨–÷µ‹£¨≈£±∆∞°£¨“ª∏ˆºÚµ•Ã‚∏„’‚√¥ƒ—£¨≈™µƒ¿œ◊”≤Óµ„Õ∂Àﬂ
{
    if (n == 0)
    {
        return "0";
    }
    
    vector<int> matrix = { 1, 0, 0, 1 };
    int start = 0;
    while (!LastFourDigitsOfFnSub(matrix, start, n))
    {
        
    }
    return to_string(matrix[1] % 10000);
}

int UniqueMorseRepresentations(vector<string> &words) // LintCode 1013. Unique Morse Code Words
{
    vector<string> hash = { ".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.." };
    unordered_map<string, bool> trans_string;
    for (auto it : words)
    {
        string cur_key = "";
        for (auto str_chr : it)
        {
            cur_key += hash[str_chr - 'a'];
        }
        if (!trans_string.count(cur_key))
        {
            trans_string[cur_key] = true;
        }
    }
    return trans_string.size();
}

int Fibonacci(int n) // LintCode 366. Ï≥≤®ƒ…∆ı ˝¡– ƒ„‘⁄∂∫Œ“£ø
{
    if (n == 1)
    {
        return 0;
    }
    int result = int(sqrt(5) / 5 * (pow(((1 + sqrt(5)) / 2), (n - 1)) - pow(((1 - sqrt(5)) / 2), (n - 1))));
    return result;
}

double FindMaxAverage(vector<int> &nums, int k) // LintCode 868. ◊” ˝◊Èµƒ◊Ó¥Û∆Ωæ˘÷µ
{
    double cur_sum = 0;
    double move_sum = 0;
    uint i = 0;
    for (; i < nums.size() && i < (uint)k; i++)
    {
        cur_sum += nums[i];
    }
    
    move_sum = cur_sum;
    
    for (; i < nums.size(); i++)
    {
        move_sum += nums[i] - nums[i - k];
        cur_sum = max(cur_sum, move_sum);
    }
    return cur_sum / min((uint)k, (uint)nums.size());
}

int CountPalindromicSubstrings(string &str) // LintCode 837. Palindromic Substrings
{
    int result = 0;
    for (uint i = 0; i < str.length(); i++)
    {
        result++;
        // ÷––ƒµ„
        for (int j = i, distance = 1; j - distance >= 0 && distance + (uint)j < str.length(); distance++)
        {
            if (str[j - distance] == str[distance + j])
            {
                result++;
            }
            else
            {
                break;
            }
        }
        
        // ÷––ƒ◊Ûµ„
        for (int j = i, distance = 1; j + 1 - distance >= 0 && distance + (uint)j < str.length(); distance++)
        {
            if (str[j + 1 - distance] == str[distance + j])
            {
                result++;
            }
            else
            {
                break;
            }
        }
    }
    return result;
}

/*------------------------------------------------------------ ª™¿ˆ¿ˆXXX£¨œÓƒø–Ë“™£¨∫Û–¯∑Á∏Ò±‰¡À --------------------------------------------------------------*/
TreeNode<int> * mergeTrees(TreeNode<int> * t1, TreeNode<int> * t2) // LintCode 1126. Merge Two Binary Trees
{
    if (!t1)
    {
        return t2;
    }
    if (!t2)
    {
        return t1;
    }
    
    TreeNode<int> * resultNode = NULL;
    int val = t1->val + t2->val;
    resultNode = new TreeNode<int>(val);
    
    resultNode->left = mergeTrees(t1->left, t2->left);
    resultNode->right = mergeTrees(t1->right, t2->right);
    return resultNode;
}

bool isSymmetricSub(TreeNode<int> * leftNode, TreeNode<int> * rightNode)
{
    if (leftNode && rightNode)
    {
        if (leftNode->val == rightNode->val)
        {
            return isSymmetricSub(leftNode->left, rightNode->right) && isSymmetricSub(leftNode->right, rightNode->left);
        }
    }
    if (!leftNode && !rightNode)
    {
        return true;
    }
    return false;
}

bool isSymmetric(TreeNode<int> * root) // LintCode 1360. Symmetric Tree
{
    if (!root)
    {
        return true;
    }
    
    return isSymmetricSub(root->left, root->right);
}

int findPairs(vector<int> &nums, int k) // LintCode 1187. K-diff Pairs in an Array
{
    unordered_map<int, bool> needMap;
    int result = 0;
    for (auto num : nums)
    {
        int num1 = num - k;
        int num2 = num + k;
        
        if (needMap.count(num) && needMap[num])
        {
            result++;
            needMap[num] = false;
        }
        
        if (!needMap.count(num1))
        {
            needMap[num1] = true;
        }
        
        if (!needMap.count(num2))
        {
            needMap[num2] = true;
        }
    }
    return result;
}

vector<int> rotateArray(vector<int> &nums, int k) // LintCode 1334. Rotate Array
{
    vector<int> tmpNums;
    uint len = nums.size();
    int i;
    
    k = k % len;
    for (i = 0; i < k; i++)
    {
        tmpNums.push_back(nums[len - k + i]);
    }
    for (i = 0; i < (int)len - k; i++)
    {
        tmpNums.push_back(nums[i]);
    }
    
    return tmpNums;
}

vector<vector<int>> floodFill(vector<vector<int>> &image, int sr, int sc, int newColor) // LintCode 1062. Flood Fill
{
    coordinate<int> seed;
    int seedVal = image[sr][sc];
    seed.i = sr;
    seed.j = sc;
    
    vector<coordinate<int>> coorVec = { seed };
    while (!coorVec.empty())
    {
        auto curCoor = coorVec.back();
        coorVec.pop_back();
        image[curCoor.i][curCoor.j] = newColor;
        if (curCoor.i - 1 >= 0 && image[curCoor.i - 1][curCoor.j] == seedVal)
        {
            coordinate<int> coor;
            coor.i = curCoor.i - 1;
            coor.j = curCoor.j;
            coorVec.push_back(coor);
        }
        
        if (curCoor.j - 1 >= 0 && image[curCoor.i][curCoor.j - 1] == seedVal)
        {
            coordinate<int> coor;
            coor.i = curCoor.i;
            coor.j = curCoor.j - 1;
            coorVec.push_back(coor);
        }
        
        if (curCoor.i + 1 < (int)image.size() && image[curCoor.i + 1][curCoor.j] == seedVal)
        {
            coordinate<int> coor;
            coor.i = curCoor.i + 1;
            coor.j = curCoor.j;
            coorVec.push_back(coor);
        }
        
        if (curCoor.j + 1 < (int)image[0].size() && image[curCoor.i][curCoor.j + 1] == seedVal)
        {
            coordinate<int> coor;
            coor.i = curCoor.i;
            coor.j = curCoor.j + 1;
            coorVec.push_back(coor);
        }
    }
    return image;
}

int findRadius(vector<int> &houses, vector<int> &heaters) // LintCode 1219. Heaters Fuck Fuck ‘⁄’‚µ¿Ã‚…œ¿œ◊”∂º–¥‘Œ¡À
{
    sort(houses.begin(), houses.end());
    sort(heaters.begin(), heaters.end());
    int result = 0;
    int heaterIndex = 0;
    
    if (heaters.size() == 1)
    {
        return max(abs(heaters[0] - houses[0]), abs(heaters[0] - houses[houses.size() - 1]));
    }
    for (uint i = 0; i < houses.size(); i++)
    {
        while (houses[i] > heaters[heaterIndex + 1] && heaterIndex < (int)heaters.size() - 2)
        {
            heaterIndex++;
        }
        if (heaterIndex == 0 && heaters[0] > houses[i])
        {
            result = max(result, abs(houses[i] - heaters[heaterIndex]));
        }
        else
        {
            result = max(result, min(abs(houses[i] - heaters[heaterIndex]), abs(heaters[heaterIndex + 1] - houses[i])));
        }
    }
    return result;
}

string longestWord(vector<string> &words) // LintCode 1071. Longest Word in Dictionary
{
    vector<map<string, vector<char>>> wordvec;
    vector<char> oneWord;
    unordered_set<string> result;
    uint resultLen = 0;
    
    for (uint i = 0; i < words.size(); i++)
    {
        string curStr = words[i];
        if (curStr.length() == 1)
        {
            oneWord.push_back(curStr[0]);
            continue;
        }
        
        string subStr = curStr.substr(0, curStr.length() - 1);
        for (uint j = wordvec.size(); j < curStr.length() - 1; j++)
        {
            map<string, vector<char>> newMap;
            wordvec.push_back(newMap);
        }
        
        if (wordvec[curStr.length() - 2].count(subStr))
        {
            wordvec[curStr.length() - 2][subStr].push_back(curStr[curStr.length() - 1]);
        }
        else
        {
            vector<char> newVec = { curStr[curStr.length() - 1] };
            wordvec[curStr.length() - 2][subStr] = newVec;
        }
    }
    
    for (uint i = 0; i < oneWord.size(); i++)
    {
        unordered_set<string> curChar = { { oneWord[i] } };
        unordered_set<string> tmpChar;
        for (uint j = 0; j < wordvec.size(); j++)
        {
            for (auto oneChar : curChar)
            {
                if (wordvec[j].count(oneChar))
                {
                    bool needInsert = false;
                    if (oneChar.length() + 1 > resultLen)
                    {
                        resultLen = (uint)oneChar.length() + 1;
                        result.clear();
                        needInsert = true;
                    }
                    else if (oneChar.length() + 1 == resultLen)
                    {
                        needInsert = true;
                    }
                    
                    for (auto nextWord : wordvec[j][oneChar])
                    {
                        tmpChar.insert(oneChar + nextWord);
                        if (needInsert)
                        {
                            result.insert(oneChar + nextWord);
                        }
                    }
                }
            }
            if (tmpChar.size() == 0)
            {
                break;
            }
            curChar = tmpChar;
        }
    }
    
    if (result.empty() && oneWord.empty())
    {
        return "";
    }
    
    if (result.empty())
    {
        char minChar = oneWord[0];
        for (auto oneChar : oneWord)
        {
            if (minChar > oneChar)
            {
                minChar = oneChar;
            }
        }
        return { minChar };
    }
    string minString = *(result.begin());
    for (auto resultWord : result)
    {
        if (minString > resultWord)
        {
            minString = resultWord;
        }
    }
    return minString;
}

int countBinarySubstrings(string &s) // LintCode 1079. Count Binary Substrings
{
    int count = 0;
    for (int i = 0; i < (int)s.size() - 1; i++)
    {
        if (s[i] == s[i + 1])
        {
            continue;
        }
        
        count++;
        int dexDis = 0;
        while (i + dexDis + 2 < (int)s.size() && i - dexDis - 1 >= 0)
        {
            if (s[i - dexDis] == s[i - dexDis - 1] && s[i + dexDis  + 1] == s[i + dexDis + 2])
            {
                count++;
            }
            else
            {
                break;
            }
            dexDis++;
        }
    }
    
    return count;
}

int similarRGBSub(char str)
{
    if (str <= '9')
    {
        return str - '0';
    }
    else if (str <= 'f')
    {
        return str - 'a' + 10;
    }
    else
    {
        return str - 'F' + 10;
    }
}

string similarRGB(string &color) // LintCode 1017. Similar RGB Color
{
    string result = "#";
    
    for (uint i = 1; i < color.length(); i += 2)
    {
        if (color[i] == '0')
        {
            if (similarRGBSub(color[i + 1]) - 0 <= 8)
            {
                result += "00";
            }
            else
            {
                result += "11";
            }
        }
        else if (color[i] == 'F' || color[i] == 'f')
        {
            if (15 - similarRGBSub(color[i + 1]) <= 8)
            {
                result += "ff";
            }
            else
            {
                result += "ee";
            }
        }
        else
        {
            if (color[i + 1] > color[i])
            {
                if (similarRGBSub(color[i + 1]) - similarRGBSub(color[i]) <= 8)
                {
                    result += color[i];
                    result += color[i];
                }
                else
                {
                    result += color[i] + 1;
                    result += color[i] + 1;
                }
            }
            else if (color[i + 1] < color[i])
            {
                if (similarRGBSub(color[i]) - similarRGBSub(color[i + 1]) <= 8)
                {
                    result += color[i];
                    result += color[i];
                }
                else
                {
                    result += color[i] - 1;
                    result += color[i] - 1;
                }
            }
            else
            {
                result += color[i];
                result += color[i];
            }
        }
    }
    return result;
}

vector<string> subdomainVisits(vector<string> &cpdomains) // LintCode 1006. Subdomain Visit Count
{
    vector<string> result;
    unordered_map<string, int> strMap;
    
    for (auto cur_str : cpdomains)
    {
        int spaceIndex = cur_str.find(' ');
        int nextIndex = cur_str.length() - 1;
        int num = atoi(cur_str.substr(0, spaceIndex).c_str());
        
        while (nextIndex = cur_str.find_last_of('.', nextIndex - 1))
        {
            string subStr;
            if (string::npos == nextIndex)
            {
                subStr = cur_str.substr(spaceIndex + 1);
                strMap[subStr] += num;
                break;
            }
            else
            {
                subStr = cur_str.substr(nextIndex + 1);
                strMap[subStr] += num;
            }
        }
    }
    
    char numStr[10];
    for (auto strIt : strMap)
    {
        sprintf(numStr, "%d", strIt.second);
        result.push_back(numStr + string(" ") + strIt.first);
    }
    
    return result;
}

int findNthDigit(int n) // LintCode 1256. Nth Digit
{
    int curDex = 1;
    long long curMul = 9;
    
    while (n > curMul * curDex)
    {
        n -= (int)curMul * curDex;
        curMul *= 10;
        curDex++;
    }
    
    char curVal[15];
    sprintf(curVal, "%d", (int)(pow(10, curDex - 1) + (n - 1) / curDex));
    return curVal[(n - 1) % curDex] - '0';
}

bool repeatedSubstringPattern(string &s) // LintCode 1227. Repeated Substring Pattern
{
    if (s.length() == 1)
    {
        return true;
    }
    
    char firstCha = s[0];
    for (uint i = 1; i < s.length(); i++)
    {
        if (s[i] == firstCha)
        {
            cout << i << endl;
            if (s.length() % i != 0)
            {
                continue;
            }
            string subStr = s.substr(0, i);
            bool isOk = true;
            for (uint j = i + 1; j < s.length(); j++)
            {
                if (s[j] != subStr[j % i])
                {
                    isOk = false;
                    break;
                }
            }
            if (!isOk)
            {
                continue;
            }
            return true;
        }
    }
    return false;
}

string convertToTitle(int n) // LintCode 1350. Excel Sheet Column Title
{
    string result;
    int index = 14;
    result.append("0", 15);
    
    while (n != 0)
    {
        result[index--] = (n - 1) % BIG_ENGLISH_CHAR_NUM + 'A';
        n = (n - 1) / BIG_ENGLISH_CHAR_NUM;
    }
    result.erase(0, index + 1);
    return result;
}

vector<int> numberOfLines(vector<int> &widths, string &S) // LintCode 1011. Number of Lines To Write String
{
    if (S.empty())
    {
        return vector<int> {0, 0};         
    }

    int count = 1;
    int cur_num = 0;
    for( auto curChr : S )
    {
        if( cur_num + widths[curChr - 'a'] <= 100)
        {
            cur_num += widths[curChr - 'a'];
        }
        else
        {
            cur_num = widths[curChr - 'a'];
            count++;
        }
    }
    return vector<int> {count, cur_num};
}

vector<int> findDisappearedNumbers(vector<int> &nums) // LintCode 1236. Find All Numbers Disappeared in an Array
{
    string numStr = string(nums.size(), '0');
    vector<int> result;
    for ( auto curNum : nums )
    {
        numStr[curNum - 1] = '1';
    }

    for (size_t i = 0; i < nums.size(); i++)
    {
        if(numStr[i] != '1')
        {
            result.push_back((int)i + 1);
        }
    }

    return result;
}

vector<string> findRestaurant(vector<string> &list1, vector<string> &list2) // LintCode 1143. Minimum Index Sum of Two Lists
{
    unordered_map<string, size_t> listSet;
    vector<string> result;
    size_t curIndexSum = SIZE_T_MAX;
    

    for (size_t i = 0; i < list1.size(); i++)
    {
        listSet[list1[i]] = i;
    }
    
    for(size_t i = 0; i < list2.size(); i++)
    {
        if(listSet.count(list2[i]) > 0)
        {
            if(listSet[list2[i]] + i < curIndexSum)
            {
                curIndexSum = listSet[list2[i]] + i;
                result.clear();
                result.push_back(list2[i]);
            }
            else if(listSet[list2[i]] + i == curIndexSum)
            {
                result.push_back(list2[i]);
            }
        }
    }

    return result;
}

int countPrimes(int n) //LintCode 1324. Count Primes
{
    if(n <= 2)
    {
        return 0;
    }

    bool* numPtr = new bool[n - 2];
    int result = 0;

    for (uint i = 0; i < n - 2; i++)
    {
        numPtr[i] = true;
    }
    for (uint i = 0; i < n - 2; i++)
    {
        if(numPtr[i])
        {
            uint count = (n - 1) / (i + 2);
            for(uint mul = 2; mul <= count; mul++)
            {
                numPtr[(i + 2) * mul - 2] = false;
            }
        }
    }

    for (uint i = 0; i < n - 2; i++)
    {
        if(numPtr[i])
        {
            result++;
        }
    }
    
    delete[] numPtr;
    return result;
}

ParentTreeNode<int> * lowestCommonAncestorII(ParentTreeNode<int> * root, ParentTreeNode<int> * A, ParentTreeNode<int> * B)
{
    unordered_set<ParentTreeNode<int> *> nodeSet1;
    unordered_set<ParentTreeNode<int> *> nodeSet2;
    ParentTreeNode<int> * ATemp = A;
    ParentTreeNode<int> * BTemp = B;

    while(ATemp != NULL || BTemp != NULL)
    {
        if(ATemp)
        {
            if(nodeSet2.count(ATemp) > 0)
            {
                return ATemp;
            }
            else
            {
                nodeSet1.insert(ATemp);
                ATemp = ATemp->parent;
            }
        }
        if(BTemp)
        {
            if(nodeSet1.count(BTemp) > 0)
            {
                return BTemp;
            }
            else
            {
                nodeSet2.insert(BTemp);
                BTemp = BTemp->parent;
            }
        }
    }
    return root;
}

string licenseKeyFormatting(string &S, int K) //LintCode 1214. License Key Formatting(这有些翻译真难)
{
    string result;
    int curDashNum = 0;
    for (int i = S.length() - 1; i >= 0; i--)
    {
        if(S[i] == '-')
        {
            continue;
        }
        if(curDashNum == K)
        {
            result.push_back('-');
            curDashNum = 0;
        }
        if(S[i] <= 'z' && S[i] >= 'a')
        {
            result.push_back(S[i] - 'a' + 'A');
        }
        else
        {
            result.push_back(S[i]);
        }
        curDashNum ++;
    }
    reverse(result.begin(), result.end());
    return result;
}

int sumOfLeftLeaves(TreeNode<int> * root) // LintCode 1254. Sum of Left Leaves
{
    if (!root) return 0;
    if (root->left && !root->left->left && !root->left->right)
    {
        return root->left->val + sumOfLeftLeaves(root->right);
    }
    return sumOfLeftLeaves(root->left) + sumOfLeftLeaves(root->right);
}

string removeDuplicateLetters(string &s) // LintCode 834. Remove Duplicate Letters 
{
    vector<int> m(SMALL_ENGLISH_CHAR_NUM,0);
    vector<int> visited (SMALL_ENGLISH_CHAR_NUM,0);
    string res = "0";
    for (auto a : s)
    {
        ++m[a - 'a'];   
    }
    for (auto a : s) {
        --m[a - 'a'];
        if (visited[a - 'a']) continue;
        while (a < res.back() && m[res.back() - 'a']) {
            visited[res.back() - 'a'] = 0;
            res.pop_back();
        }
        res += a;
        visited[a - 'a'] = 1;
    }
    return res.substr(1);
}

int hammingWeight(unsigned int n) // LintCode 1332. Number of 1 Bits
{
    int count = 0;
    while(n > 0)
    {
        if(n & 0X01)
        {
            count++;
        }
        n >>= 1;
    }
    return count;
}

char findTheDifferenceError(string &s, string &t) // LintCode 1266. Find the Difference LintCode 真的是越来越烂了，题目翻译有问题，各种解释不明确
{
    int start = 0;
    int end = t.length() - 1;
    int lastEnd = 0;
    while(start != end)
    {
        int middle = (end - start) / 2 + start;
        if(t[middle] != s[middle])
        {
            end = middle;
            continue;
        }

        bool isOver = false;
        for(int i = middle - 1; i > lastEnd; i--)
        {
            if(t[i] != s[i])
            {
                end = middle;
                isOver = true;
                break;
            }
            if(t[i] != t[i+1])
            {
                start = middle + 1;
                lastEnd = start;
                isOver = true;
                break;
            }
        }
        if(!isOver)
        {
            start = middle + 1;
        }
    }
    return t[start];
}

char findTheDifference(string &s, string &t) // LintCode 1266. Find the Difference
{
    bool charArr[SMALL_ENGLISH_CHAR_NUM] = {false};
    for(auto i : s)
    {
        charArr[i - 'a'] = !charArr[i - 'a'];
    }
    for(auto i : t)
    {
        charArr[i - 'a'] = !charArr[i - 'a'];
    }
    for(int i = 0; i < SMALL_ENGLISH_CHAR_NUM; i++)
    {
        if(charArr[i])
        {
            return i + 'a';
        }
    }
    return 'a';
}

bool validPalindrome(string &s) // LintCode 891. Valid Palindrome II
{
    int start = 0;
    int end = s.length() - 1;
    bool isOver = false;

    for(; start < end; start++, end--)
    {
        if(s[start] != s[end])
        {
            isOver = true;
            break;
        }
    }

    if(!isOver)
    {
        return true;
    }

    int startTemp = start + 1;
    int endTemp = end;
    for(; startTemp < endTemp; startTemp++, endTemp--)
    {
        if(s[startTemp] != s[endTemp])
        {
            isOver = false;
            break;
        }
    }

    if(isOver)
    {
        return true;
    }

    startTemp = start;
    endTemp = end - 1;
    for(; startTemp < endTemp; startTemp++, endTemp--)
    {
        if(s[startTemp] != s[endTemp])
        {
            isOver = true;
            return false;
        }
    }

    return true;
}

#endif func.h