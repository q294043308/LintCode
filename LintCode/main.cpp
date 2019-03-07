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
    vector<vector<int>>date = { { 1, 0 }, { 1, 2 }, { 2, 1 }, { 0, 1 } };
    nearestRestaurant(date, 3);
};
