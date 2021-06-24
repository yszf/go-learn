#include "libcgo.h"
#include <stdio.h>

// go tool cgo cgo.go
// go build -buildmode=c-archive -o libcgo.a cgo.go
// gcc test_cgo.c -L. -lcgo -lpthread -o test_cgo 
// gcc test_cgo.c libcgo.a -lpthread -o test_cgo 

void foo() {
//    extern int add(int a, int b);
    printf("sum = %d\n", add(1, 1));
}

int main() {
    foo();

    return 0;
}