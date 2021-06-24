#include "cgo.h"
#include <stdio.h>

void foo() {
//    extern int add(int a, int b);
    printf("sum = %d\n", add(1, 1));
}

int main() {
    foo();

    return 0;
}