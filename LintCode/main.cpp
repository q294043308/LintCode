//  main.cpp
//  LintCode
//
//  Created by DongXuxuan on 2018/7/25.
//  Copyright © 2018 DongXuxuan. All rights reserved.

#include "define.h"
#include "func1.h"
#include "func2.h"
#include "func11.h"

class Base1
{
public:
    virtual void base1_fun1() {
        cout << "Base1:base1_fun1" << endl;
    }
    virtual void base1_fun2() {
        cout << "Base1:base1_fun2" << endl;
    }

    Base1(){
        cout << "new" << endl;
    }

    Base1(const Base1 &b){
        cout << "copy new" << endl;
    }
};

class Derive1 : public Base1
{
public:
    virtual void base1_fun1() {
        cout << "Derive1:base1_fun1" << endl;
    }
    virtual void base1_fun2() {
        cout << "Derive1:base1_fun2" << endl;
    }
};

int main(int argc, char* argv[])
{
    Derive1* a = new(Derive1);
    a->base1_fun1();
    Base1* b = (Base1*)a;
    b->base1_fun1();
};
