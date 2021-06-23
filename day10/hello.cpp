// hello.c
#include "hello.h"

void SayHello(const char* s) {
    printf("%s\n", s);
} 

int main() {
    SayHello("hello world!");

    return 0;
}