#include "number.h"

// gcc -c -o number.o number.c
// ar rcs libnumber.a number.o
// gcc -shared -o libnumber.so number.c 动态库
// int number_add_mod(int a, int b, int mod) {
//     return (a+b)%mod;
// }