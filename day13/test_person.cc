#include "./person/person.h"

#include <stdio.h>

// go build -buildmode=c-archive -o person_capi.a person_capi.go
// g++ test_person.cc person_capi.a -lpthread -o test_person -std=c++11
int main() {
    auto p = new Person("gopher", 10);

    char buf[64];
    char* name = p->GetName(buf, sizeof(buf)-1);
    int age = p->GetAge();

    printf("%s, %d years old.\n", name, age);
    delete p;

    return 0;
}