//  define.h
//  LintCode
//
//  Created by DongXuxuan on 2018/8/21.
//  Copyright © 2018 DongXuxuan. All rights reserved.

#ifndef define
#define define

#include "stdio.h"
#include "iostream"
#include "string.h"
#include "math.h"
#include "stdlib.h"
#include "vector"
#include "map"
#include "set"
#include "string"
#include "stack"
#include "queue"
#include "unordered_set"
#include "unordered_map"
#include <algorithm>
#include "memory.h"
#include "time.h"

using namespace std;
typedef unsigned int                                    uint;

#define MAXWINTNUM                               (int)pow(2, 61) - 1
#define MAXINTNUM                               (int)pow(2, 31) - 1
#define MININTNUM                                - (int)(pow(2, 31)) + 1
#define CHARNUM                                     128
#define ENGLISH_CHAR_NUM                    (26 * 2)
#define BIG_ENGLISH_CHAR_NUM            26
#define SMALL_ENGLISH_CHAR_NUM      26
#define MIN_LEN                                       0x7fffffff
#define CIRCLE_PI                                       3.14
#define THIRTY                                       30

template<class val_type> class TreeNode {
public:
    val_type val;
    TreeNode *left, *right;
    TreeNode(val_type val)
    {
        this->val = val;
        this->left = this->right = NULL;
    }
};

template<class val_type> class ParentTreeNode {
public:
    val_type val;
    ParentTreeNode *parent, *left, *right;
    ParentTreeNode(val_type val)
    {
        this->val = val;
        this->parent = this->left = this->right = NULL;
    }
};

template<class val_type> class ListNode
{
public:
    val_type val;
    ListNode *next;
    ListNode(val_type val)
    {
        this->val = val;
        this->next = NULL;
    }
    ~ListNode()
    {
        cout << "delete node" << endl;
    }
};

template<class val_type> bool IsEqual(TreeNode<val_type> * T1, TreeNode<val_type> * T2)
{
    if (T1 == NULL && T2 == NULL)
    {
        return true;
    }
    if (T1 == NULL || T2 == NULL)
    {
        return false;
    }
    if (T1->val == T2->val)
    {
        bool left_equal = IsEqual(T1->left, T2->left);
        bool right_equal = IsEqual(T1->right, T2->right);
        return left_equal && right_equal;
    }
    return false;
}

class Solution { // LintCode 204
public:
    static Solution* getInstance()
    {
        static Solution* is_have = NULL;
        if (is_have == NULL)
        {
            is_have = new Solution();
        }
        return is_have;
    }
};

class Toy { // LintCode 496
public:
    virtual void talk() const = 0;
};

class Dog : public Toy {
    void talk() const
    {
        cout << "Wow" << endl;
    }
};

class Cat : public Toy {
    void talk() const
    {
        cout << "Meow" << endl;
    }
};

class ToyFactory {
public:
    Toy* getToy(string& type) {
        if (type == "Dog")
        {
            return new Dog();
        }
        if (type == "Cat")
        {
            return new Cat();
        }
        return NULL;
    }
};

template<class val_type> struct coordinate
{
    val_type i = 0;
    val_type j = 0;
};

template<class val_type> int CountNodes(ListNode<val_type> * head) //  LintCode 466
{
    int count = 0;
    ListNode<val_type>* tmp = head;
    while (tmp != NULL)
    {
        count++;
        tmp = tmp->next;
    }
    return count;
}

template<class val_type> vector<vector<val_type>> CreatDoubleVector(val_type argc[][256], int size)
{
    vector<vector<val_type>> DoubleVector;
    for (int i = 0; i < size; i++)
    {
        vector<val_type> tmp(argc[i]);
        DoubleVector.insert(DoubleVector.end(), tmp);
    }
}

template<class val_type> ListNode<val_type> * CreatListNode(val_type argc[], int size)
{
    ListNode<val_type> * lastnode = NULL;
    ListNode<val_type> * listnode = NULL;
    for (int i = 0; i < size; i++)
    {
        if (i == 0)
        {
            lastnode = listnode = new ListNode<val_type>(argc[i]);
        }
        else
        {
            lastnode->next = new ListNode<val_type>(argc[i]);
            lastnode = lastnode->next;
        }
    }
    return listnode;
}

template<class val_type> val_type GetMinValue(val_type a, val_type b)
{
    return a < b ? a : b;
}

void Recover(vector<int> &nums, int start_dex, int end_dex)
{
    while (start_dex < end_dex)
    {
        int tmp = nums[start_dex];
        nums[start_dex] = nums[end_dex];
        nums[end_dex] = tmp;
        start_dex++;
        end_dex--;
    }
}

template<class val_type> void FindNode(TreeNode<val_type> * T1, TreeNode<val_type> *T2, vector<TreeNode<val_type>* > &NodeList)
{
    if (T1 == NULL)
    {
        return;
    }
    
    if (T1->val == T2->val)
    {
        NodeList.push_back(T1);
    }
    FindNode(T1->left, T2, NodeList);
    FindNode(T1->right, T2, NodeList);
}

#endif
