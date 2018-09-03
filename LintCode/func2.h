//  func2.h
//  LintCode
//
//  Created by 董旭轩 on 2018/8/21.
//  Copyright © 2018年 董旭轩. All rights reserved.

#ifndef func2.h
#define func2.h

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



#endif