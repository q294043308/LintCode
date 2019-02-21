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
    vector<vector<int>> lists = { { -203653, -202653 }, {-288256, -287256}, {19396, 20396}, {835984, 836984}, {-976643, -975643}, {368729, 369729}, {501747, 502747}, {847647, 848647} };
    isInterval(lists, 368722);
};
