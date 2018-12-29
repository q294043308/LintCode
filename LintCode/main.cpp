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
    vector<string> str = { "ab.cd+cd@jiu.zhang.com", "ab.cd+cd@jiuzhang.com", "ab+cd.cd@jiuzhang.com" };
    countGroups(str);
};
