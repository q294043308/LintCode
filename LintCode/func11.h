//  func3.h
//  LeetCode
//
//  Created by DongXuxuan on 2019/3/30.
//  Copyright 2019 DongXuxuan. All rights reserved.

#ifndef func3
#define func3

#include "define.h"

// 6. ZigZag Conversion
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

// 8. String to Integer (atoi)
int myAtoi(string str) {
    long long res = 0;
    uint index = 0;
    bool isPositive = true;

    for (; index < str.length(); index ++){
        if (str[index] != ' '){
            if (str[index] == '-'){
                isPositive = false;
                index++;
                break;
            }
            else if (str[index] == '+'){
                isPositive = true;
                index++;
                break;
            }
            else if (str[index] <= '9' && str[index] >= 0){
                break;
            }
            else{
                return int(res);
            }
        }
    }
    
    for (; index < str.length(); index++){
        if (str[index] > '9' || str[index] < '0'){
            return int(res);
        }
        if (isPositive){
            if (res * 10 + str[index] - '0' < MAXINTNUM){
                res = res * 10 + str[index] - '0';
            }
            else{
                return MAXINTNUM;
            }
        }
        else{
            if (res * 10 - str[index] + '0' > MININTNUM){
                res = res * 10 - str[index] + '0';
            }
            else{
                return MININTNUM;
            }
        }
    }
    return int(res);
}

#endif