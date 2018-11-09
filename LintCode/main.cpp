//  main.cpp
//  LintCode
//
//  Created by DongXuxuan on 2018/7/25.
//  Copyright © 2018 DongXuxuan. All rights reserved.

#include "define.h"
#include "func1.h"
#include "func2.h"

int main(int argc, char* argv[])
{
    vector<vector<int>> s;
    s.push_back(vector<int>{1, 2});
    s.push_back(vector<int>{1, 3});
    s.push_back(vector<int>{3, 2});
    findRedundantConnection(s);
};
