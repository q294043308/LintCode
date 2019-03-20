//  func3.h
//  LeetCode
//
//  Created by DongXuxuan on 2019/3/30.
//  Copyright 2019 DongXuxuan. All rights reserved.

#ifndef func3
#define func3

#include "define.h"

string convert(string s, int numRows) {
    if (numRows == 1){
        return s;
    }

    string res;
    for (int i = 0; i < numRows; i++){
        res += s[i];

        for (int j = i;;){
            if (i != 0 && i != numRows - 1){
                if (j + 2 * numRows - i * 2 - 2 < int(s.length())){
                    res += s[j + 2 * numRows - i * 2 - 2];
                }
                else{
                    break;
                }
            }
            if (j + 2 * numRows - 2 < int(s.length())){
                res += s[j + 2 * numRows - 2];
            }
            else{
                break;
            }
            j += 2 * numRows - 2;
        }
    }

    return res;
}

#endif