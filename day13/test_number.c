#include <stdio.h>

void goPrintln(char*);
int number_add_mod(int a, int b, int mod);

// go build -buildmode=c-archive -o number.a test_number2.go
// gcc test_number.c number.a -o test_number -lpthread
int main() {
    int a = 10;
    int b = 5;
    int c = 12;

    int x = number_add_mod(a, b, c);
    printf("(%d+%d)%%%d = %d\n", a, b, c, x);

    goPrintln("done");
    return 0;
}